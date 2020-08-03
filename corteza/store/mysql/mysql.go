package mysql

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/corteza/store/provisioner"
	"github.com/cortezaproject/corteza-server/corteza/store/rdbms"
	"github.com/go-sql-driver/mysql"
	"strings"
)

type (
	Store struct {
		*rdbms.Store
	}
)

func New(ctx context.Context, dsn string) (s *Store, err error) {
	var cfg *rdbms.Config

	if cfg, err = ProcDataSourceName(dsn); err != nil {
		return nil, err
	}

	cfg.PlaceholderFormat = squirrel.Question
	cfg.TxRetryErrHandler = txRetryErrHandler

	s = new(Store)
	if s.Store, err = rdbms.New(ctx, cfg); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Store) Provision() provisioner.Executor {
	return provisioner.Do(
		s.ProvisionCore(),
		s.ProvisionCompose(),
		s.ProvisionMessaging(),
	)
}

// ProcDataSourceName validates given DSN and ensures
// params are present and correct
//
// If param is missing it sets it to default but returns
// error in case of incorrect param value
//
// See https://github.com/go-sql-driver/mysql for available dsn params
//
// @todo similar fallback (autoconfig) for collation=utf8mb4_general_ci
//       but allow different collation values
//
func ProcDataSourceName(in string) (*rdbms.Config, error) {
	const (
		schemeDel   = "://"
		validScheme = "mysql"
	)

	var (
		endOfSchema = strings.Index(in, schemeDel)
		c           = &rdbms.Config{}
	)

	if endOfSchema > 0 && (in[:endOfSchema] == validScheme || strings.HasPrefix(in[:endOfSchema], validScheme+"+")) {
		c.DriverName = in[:endOfSchema]
		c.DataSourceName = in[endOfSchema+len(schemeDel):]
	} else {
		return nil, fmt.Errorf("expecting valid schema (mysql://) at the beginning of the DSN (%s)", in)
	}

	if pdsn, err := mysql.ParseDSN(c.DataSourceName); err != nil {
		return nil, err
	} else {
		// Ensure driver parses time
		pdsn.ParseTime = true

		c.DataSourceName = pdsn.FormatDSN()
		c.DBName = pdsn.DBName
	}

	return c, nil
}

func txRetryErrHandler(try int, err error) bool {
	for errors.Unwrap(err) != nil {
		err = errors.Unwrap(err)
	}

	var mysqlErr, ok = err.(*mysql.MySQLError)
	if !ok || mysqlErr == nil {
		return false
	}

	switch mysqlErr.Number {
	case
		1205, // lock within transaction
		1206, // total number of locks exceeded the lock table size
		1208, // not allowed while thread is holding global read lock
		1213: // deadlock found
		return true

	}

	return false
}
