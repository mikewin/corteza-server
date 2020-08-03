package rdbms

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// corteza/store/applications.yaml

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/corteza/store"
	"github.com/cortezaproject/corteza-server/system/types"
	"github.com/jmoiron/sqlx"
)

// SearchApplications returns all matching rows
//
// This function calls convertApplicationFilter with the given
// types.ApplicationFilter and expects to receive a working squirrel.SelectBuilder
func (s Store) SearchApplications(ctx context.Context, f types.ApplicationFilter) (types.ApplicationSet, types.ApplicationFilter, error) {
	query, err := s.convertApplicationFilter(f)
	if err != nil {
		return nil, f, err
	}

	var set = types.ApplicationSet{}

	if f.Count, err = Count(ctx, s.db, query); err != nil || f.Count == 0 {
		return nil, f, err

	}

	return set, f, FetchPaged(ctx, s.db, query, f.PageFilter, &set)

}

// LookupApplicationByID searches for application by their ID
//
// It returns application even if deleted or suspended
func (s Store) LookupApplicationByID(ctx context.Context, id uint64) (*types.Application, error) {
	return s.ApplicationLookup(ctx, squirrel.Eq{
		"app.id": id,
	})
}

// CreateApplication creates one or more rows in sys_application table
func (s Store) CreateApplication(ctx context.Context, rr ...*types.Application) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Insert(s.ApplicationTable()).SetMap(s.ApplicationEnc(res)))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// UpdateApplication updates one or more existing rows in sys_application
func (s Store) UpdateApplication(ctx context.Context, rr ...*types.Application) error {
	return s.PartialUpdateApplication(ctx, nil, rr...)
}

// PartialUpdateApplication updates one or more existing rows in sys_application
//
// It wraps the update into transaction and can perform partial update by providing list of updatable columns
func (s Store) PartialUpdateApplication(ctx context.Context, onlyColumns []string, rr ...*types.Application) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = s.ExecUpdateApplications(
				ctx,
				squirrel.Eq{s.preprocessColumn("app.id", ""): s.preprocessValue(res.ID, "")},
				s.ApplicationEnc(res).Skip("id").Only(onlyColumns...))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveApplication removes one or more rows from sys_application table
func (s Store) RemoveApplication(ctx context.Context, rr ...*types.Application) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Delete(s.ApplicationTable("app")).Where(squirrel.Eq{s.preprocessColumn("app.id", ""): s.preprocessValue(res.ID, "")}))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveApplicationByID removes row from the sys_application table
func (s Store) RemoveApplicationByID(ctx context.Context, ID uint64) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Delete(s.ApplicationTable()).Where(squirrel.Eq{s.preprocessColumn("app.id", ""): s.preprocessValue(ID, "")}))
}

// TruncateApplications removes all rows from the sys_application table
func (s Store) TruncateApplications(ctx context.Context) error {
	return Truncate(ctx, s.DB(), s.ApplicationTable())
}

// ExecUpdateApplications updates all matchhed (cnd) rows in sys_application with given data
func (s Store) ExecUpdateApplications(ctx context.Context, cnd squirrel.Sqlizer, set store.Payload) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Update(s.ApplicationTable("app")).Where(cnd).SetMap(set))
}

// ApplicationLookup calls Lookup() and returns (if found) types.Application
func (s Store) ApplicationLookup(ctx context.Context, cnd squirrel.Sqlizer) (*types.Application, error) {
	var (
		u   = &types.Application{}
		err = s.Lookup(ctx, u, s.QueryApplications(), cnd)
	)

	if err == nil {
		return u, nil
	}

	return nil, err
	//return s.scanApplication(s.LookupWithScan(ctx, s.QueryApplications(), cnd))
}

// QueryApplications returns squirrel.SelectBuilder with set table and all columns
func (s Store) QueryApplications() squirrel.SelectBuilder {
	return s.Select(s.ApplicationTable("app"), s.ApplicationColumns("app")...)
}

// ApplicationTable name of the db table
func (Store) ApplicationTable(aa ...string) string {
	var alias string
	if len(aa) > 0 {
		alias = " AS " + aa[0]
	}

	return "sys_application" + alias
}

// ApplicationColumns returns all defined table columns
//
// With optional string arg, all columns are returned aliased
func (Store) ApplicationColumns(aa ...string) []string {
	var alias string
	if len(aa) > 0 {
		alias = aa[0] + "."
	}

	return []string{
		alias + "id",
		alias + "name",
		alias + "rel_owner",
		alias + "enabled",
		alias + "unify",
		alias + "created_at",
		alias + "updated_at",
		alias + "deleted_at",
	}
}

// ApplicationEnc encodes fields from types.Application to store.Payload (map)
func (Store) ApplicationEnc(res *types.Application) store.Payload {
	return store.Payload{
		"id":         res.ID,
		"name":       res.Name,
		"rel_owner":  res.OwnerID,
		"enabled":    res.Enabled,
		"unify":      res.Unify,
		"created_at": res.CreatedAt,
		"updated_at": res.UpdatedAt,
		"deleted_at": res.DeletedAt,
	}
}
