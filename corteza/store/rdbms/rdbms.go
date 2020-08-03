package rdbms

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/corteza/store"
	"github.com/cortezaproject/corteza-server/corteza/store/bulk"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

// persistance layer
//
// all functions go under one struct
//   why? because it will be easier to initialize and pass around
//
// each domain will be in it's own file
//
// connection logic will be built in the persistence layer (making pkg/db obsolete)
//

type (
	txRetryOnErrHandler func(int, error) bool
	columnPreprocFn     func(string, string) string
	valuePreprocFn      func(interface{}, string) interface{}

	Config struct {
		DriverName     string
		DataSourceName string
		DBName         string

		PlaceholderFormat squirrel.PlaceholderFormat

		// How many times should we retry failed transaction?
		TxMaxRetries int

		// TxRetryErrHandler should return true if transaction should be retried
		//
		// Because retry algorithm varies between concrete rdbms implementations
		TxRetryErrHandler txRetryOnErrHandler

		ColumnPreprocessors map[string]columnPreprocFn
		ValuePreprocessors  map[string]valuePreprocFn
	}

	Store struct {
		config *Config

		db *sqlx.DB
	}
)

// This is the absolute maximum retries we'll allow
const TxRetryHardLimit = 100

var (
	now = func() time.Time {
		return time.Now()
	}
)

//func Instrumentation(log *zap.Logger) {
//	logger := instrumentedsql.LoggerFunc(func(ctx context.Context, msg string, keyvals ...interface{}) {
//		//spew.Dump(msg, keyvals)
//		log.With(zap.Any("kv", keyvals)).Info(msg)
//	})
//
//	sql.Register(
//		"mysql+instrumented",
//		instrumentedsql.WrapDriver(&mysql.MySQLDriver{}, instrumentedsql.WithLogger(logger)))
//
//	sql.Register(
//		"postgres+instrumented",
//		instrumentedsql.WrapDriver(&pq.Driver{}, instrumentedsql.WithLogger(logger)))
//}

func New(ctx context.Context, cfg *Config) (*Store, error) {
	var s = &Store{
		config: cfg,
	}

	if s.config.PlaceholderFormat == nil {
		s.config.PlaceholderFormat = squirrel.Question
	}

	if s.config.TxMaxRetries == 0 {
		s.config.TxMaxRetries = TxRetryHardLimit
	}

	if s.config.TxRetryErrHandler == nil {
		// Default transaction retry handler
		s.config.TxRetryErrHandler = TxNoRetry
	}

	if err := s.Connect(ctx); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Store) Connect(ctx context.Context) (err error) {
	s.db, err = sqlx.ConnectContext(ctx, s.config.DriverName, s.config.DataSourceName)
	return err
}

// Select is a shorthand for squirrel.SelectBuilder
//
// Sets passed table & columns and configured placeholder format
func (s Store) Select(table string, cc ...string) squirrel.SelectBuilder {
	return squirrel.Select(cc...).From(table).PlaceholderFormat(s.config.PlaceholderFormat)
}

// Lookup is a generic lookup query that takes select builder with additional conditions and fetches a single row
func (s *Store) Lookup(ctx context.Context, res interface{}, q squirrel.SelectBuilder, cnd squirrel.Sqlizer) error {
	query, args, err := q.Where(cnd).ToSql()
	if err != nil {
		return fmt.Errorf("could not assemble query: %w", err)
	}

	if err = s.db.GetContext(ctx, res, query, args...); err == sql.ErrNoRows {
		return store.ErrNotFound
	}

	return err
}

// Insert is a shorthand for squirrel.InsertBuilder
//
// Sets passed table and configured placeholder format
func (s Store) Insert(table string) squirrel.InsertBuilder {
	return squirrel.Insert(table).PlaceholderFormat(s.config.PlaceholderFormat)
}

// Update is a shorthand for squirrel.UpdateBuilder
//
// Sets passed table and configured placeholder format
func (s Store) Update(table string) squirrel.UpdateBuilder {
	return squirrel.Update(table).PlaceholderFormat(s.config.PlaceholderFormat)
}

// Delete is a shorthand for squirrel.DeleteBuilder
//
// Sets passed table and configured placeholder format
func (s Store) Delete(table string) squirrel.DeleteBuilder {
	return squirrel.Delete(table).PlaceholderFormat(s.config.PlaceholderFormat)
}

// LookupWithScan returns row instead of filling in the passed stuct
//func (s *Store) LookupWithScan(ctx context.Context, q squirrel.SelectBuilder, cnd squirrel.Sqlizer) (*sql.Row, error) {
//	query, args, err := q.Where(cnd).ToSql()
//	if err != nil {
//		return nil, fmt.Errorf("could not assemble query: %w", err)
//	}
//
//	return s.db.QueryRowContext(ctx, query, args...), nil
//}

func (s Store) DB() *sqlx.DB {
	return s.db
}

func (s Store) Config() *Config {
	return s.config
}

// Bulk returns channel that accepts jobs and executes them inside a transaction
//
// Note: This is experimental function!
// Final version might not return channel directly
//
func (s Store) Bulk(ctx context.Context) chan bulk.Job {
	jc := make(chan bulk.Job)

	go Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		var job bulk.Job

		for {
			select {
			case <-ctx.Done():
				return ctx.Err()

			case job = <-jc:
				if err = job.Do(ctx, s); err != nil {
					return
				}
			}
		}
	})

	return jc
}

// column preprocessor logic to modify db value before using it in condition filter
//
// It checks registered ColumnPreprocessors from config
// and then the standard set
//
// No preprocessor ("") is intentionally checked after checking registered list of preprocessors
func (s Store) preprocessColumn(col string, p string) string {
	if fn, has := s.config.ColumnPreprocessors[p]; has {
		return fn(col, p)
	}

	switch p {
	case "":
		return col
	case "lower":
		return fmt.Sprintf("LOWER(%s)", col)
	default:
		panic(fmt.Sprintf("unknown preprocessor %q used for column %q", p, col))
	}
}

// value preprocessor logic to modify input value before using it in condition filters
//
// It checks registered ValuePreprocessors from config
// and then the standard set
//
// No preprocessor ("") is intentionally checked after checking registered list of preprocessors
func (s Store) preprocessValue(val interface{}, p string) interface{} {
	if fn, has := s.config.ValuePreprocessors[p]; has {
		return fn(val, p)
	}

	switch p {
	case "":
		return val
	case "lower":
		if str, ok := val.(string); ok {
			return strings.ToLower(str)
		}
		panic(fmt.Sprintf("preprocessor %q not compatible with type %T (value: %v)", p, val, val))

	default:
		panic(fmt.Sprintf("unknown preprocessor %q used for value %v", p, val))
	}
}

func ExecuteSqlizer(ctx context.Context, db sqlx.ExecerContext, sqlizer squirrel.Sqlizer) error {
	query, args, err := sqlizer.ToSql()

	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, query, args...)

	return err
}

func Truncate(ctx context.Context, db sqlx.ExecerContext, table string) error {
	_, err := db.ExecContext(ctx, "TRUNCATE "+table)
	return err
}

// Tx begins a new db transaction and handles it's retries when possible
//
// It
func Tx(ctx context.Context, db *sqlx.DB, cfg *Config, txOpt *sql.TxOptions, task func(*sqlx.Tx) error) error {
	var (
		lastTaskErr error
		err         error
		tx          *sqlx.Tx
		try         = 1
	)

	for {
		try++

		// Start transaction
		tx, err = db.BeginTxx(ctx, txOpt)
		if err != nil {
			return nil
		}

		if lastTaskErr = task(tx); lastTaskErr == nil {
			// Task completed successfully
			break
		}

		if cfg.TxRetryErrHandler(try, lastTaskErr) {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return fmt.Errorf("failed to rollback transaction (tries: %d) on error %v: %w", try, lastTaskErr, rollbackErr)
			}

			time.Sleep(50 * time.Duration(try*50))
		}

		if try >= cfg.TxMaxRetries || try >= TxRetryHardLimit {
			err = fmt.Errorf("failed to perform transaction (tries: %d), last error: %w", try, lastTaskErr)
			break
		}

	}

	if lastTaskErr != nil {
		if err = tx.Rollback(); err != nil {
			fmt.Errorf("failed to rollback transaction: %w", err)
		}

		return lastTaskErr
	}

	return tx.Commit()
}

// TxNoRetry - Transaction retry handler
//
// Only returns false so transactions will never retry
func TxNoRetry(int, error) bool { return false }
