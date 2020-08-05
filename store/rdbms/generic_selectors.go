package rdbms

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/pkg/rh"
	"github.com/jmoiron/sqlx"
	"github.com/lann/builder"
)

// FetchOne fetches single row from table
func FetchOne(ctx context.Context, db *sqlx.DB, q squirrel.SelectBuilder, one interface{}) (err error) {
	var (
		sql  string
		args []interface{}
	)

	if sql, args, err = q.ToSql(); err != nil {
		return
	}

	if err = db.GetContext(ctx, one, sql, args...); err != nil {
		return
	}

	return
}

// Count counts all rows that match conditions from given query builder
func Count(ctx context.Context, db *sqlx.DB, q squirrel.SelectBuilder) (count uint, err error) {
	// Remove order-bys for counting
	q = builder.Delete(q, "OrderByParts").(squirrel.SelectBuilder)

	// Replace columns with count(*)
	q = builder.Delete(q, "Columns").(squirrel.SelectBuilder).Column("COUNT(*)")

	if sqlSelect, argsSelect, err := q.ToSql(); err != nil {
		return 0, err
	} else {
		if err := db.GetContext(ctx, &count, sqlSelect, argsSelect...); err != nil {
			return 0, err
		}
	}

	return count, nil
}

// FetchPaged fetches paged rows
func FetchPaged(ctx context.Context, db *sqlx.DB, q squirrel.SelectBuilder, p rh.PageFilter, set interface{}) error {
	if p.Limit+p.Offset == 0 {
		// When both, offset & limit are 0,
		// calculate both values from page/perPage params
		if p.PerPage > 0 {
			p.Limit = p.PerPage
		}

		if p.Page < 1 {
			p.Page = 1
		}

		p.Offset = (p.Page - 1) * p.PerPage
	}

	if p.Limit > 0 {
		q = q.Limit(uint64(p.Limit))
	}

	if p.Offset > 0 {
		q = q.Offset(uint64(p.Offset))
	}

	return FetchAll(ctx, db, q, set)
}

// FetchPaged fetches paged rows
func FetchAll(ctx context.Context, db *sqlx.DB, q squirrel.Sqlizer, set interface{}) error {
	if sqlSelect, argsSelect, err := q.ToSql(); err != nil {
		return err
	} else {
		return db.SelectContext(ctx, set, sqlSelect, argsSelect...)
	}
}