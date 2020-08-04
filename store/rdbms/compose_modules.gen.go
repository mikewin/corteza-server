package rdbms

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// store/compose_modules.yaml

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/store"
	"github.com/jmoiron/sqlx"
)

// SearchComposeModules returns all matching rows
//
// This function calls convertComposeModuleFilter with the given
// types.ModuleFilter and expects to receive a working squirrel.SelectBuilder
func (s Store) SearchComposeModules(ctx context.Context, f types.ModuleFilter) (types.ModuleSet, types.ModuleFilter, error) {
	query, err := s.convertComposeModuleFilter(f)
	if err != nil {
		return nil, f, err
	}

	var set = types.ModuleSet{}

	if f.Count, err = Count(ctx, s.db, query); err != nil || f.Count == 0 {
		return nil, f, err

	}

	return set, f, FetchPaged(ctx, s.db, query, f.PageFilter, &set)

}

// LookupComposeModuleByHandle searches for compose module by handle (case-insensitive)
func (s Store) LookupComposeModuleByHandle(ctx context.Context, handle string) (*types.Module, error) {
	return s.ComposeModuleLookup(ctx, squirrel.Eq{
		"cmd.handle": handle,
	})
}

// CreateComposeModule creates one or more rows in compose_module table
func (s Store) CreateComposeModule(ctx context.Context, rr ...*types.Module) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Insert(s.ComposeModuleTable()).SetMap(s.ComposeModuleEnc(res)))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// UpdateComposeModule updates one or more existing rows in compose_module
func (s Store) UpdateComposeModule(ctx context.Context, rr ...*types.Module) error {
	return s.PartialUpdateComposeModule(ctx, nil, rr...)
}

// PartialUpdateComposeModule updates one or more existing rows in compose_module
//
// It wraps the update into transaction and can perform partial update by providing list of updatable columns
func (s Store) PartialUpdateComposeModule(ctx context.Context, onlyColumns []string, rr ...*types.Module) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = s.ExecUpdateComposeModules(
				ctx,
				squirrel.Eq{s.preprocessColumn("cmd.id", ""): s.preprocessValue(res.ID, "")},
				s.ComposeModuleEnc(res).Skip("id").Only(onlyColumns...))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveComposeModule removes one or more rows from compose_module table
func (s Store) RemoveComposeModule(ctx context.Context, rr ...*types.Module) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Delete(s.ComposeModuleTable("cmd")).Where(squirrel.Eq{s.preprocessColumn("cmd.id", ""): s.preprocessValue(res.ID, "")}))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveComposeModuleByID removes row from the compose_module table
func (s Store) RemoveComposeModuleByID(ctx context.Context, ID uint64) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Delete(s.ComposeModuleTable()).Where(squirrel.Eq{s.preprocessColumn("cmd.id", ""): s.preprocessValue(ID, "")}))
}

// TruncateComposeModules removes all rows from the compose_module table
func (s Store) TruncateComposeModules(ctx context.Context) error {
	return Truncate(ctx, s.DB(), s.ComposeModuleTable())
}

// ExecUpdateComposeModules updates all matchhed (cnd) rows in compose_module with given data
func (s Store) ExecUpdateComposeModules(ctx context.Context, cnd squirrel.Sqlizer, set store.Payload) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Update(s.ComposeModuleTable("cmd")).Where(cnd).SetMap(set))
}

// ComposeModuleLookup calls Lookup() and returns (if found) types.Module
func (s Store) ComposeModuleLookup(ctx context.Context, cnd squirrel.Sqlizer) (*types.Module, error) {
	var (
		u   = &types.Module{}
		err = s.Lookup(ctx, u, s.QueryComposeModules(), cnd)
	)

	if err == nil {
		return u, nil
	}

	return nil, err
	//return s.scanComposeModule(s.LookupWithScan(ctx, s.QueryComposeModules(), cnd))
}

// QueryComposeModules returns squirrel.SelectBuilder with set table and all columns
func (s Store) QueryComposeModules() squirrel.SelectBuilder {
	return s.Select(s.ComposeModuleTable("cmd"), s.ComposeModuleColumns("cmd")...)
}

// ComposeModuleTable name of the db table
func (Store) ComposeModuleTable(aa ...string) string {
	var alias string
	if len(aa) > 0 {
		alias = " AS " + aa[0]
	}

	return "compose_module" + alias
}

// ComposeModuleColumns returns all defined table columns
//
// With optional string arg, all columns are returned aliased
func (Store) ComposeModuleColumns(aa ...string) []string {
	var alias string
	if len(aa) > 0 {
		alias = aa[0] + "."
	}

	return []string{
		alias + "id",
		alias + "handle",
		alias + "name",
		alias + "meta",
		alias + "rel_namespace",
		alias + "created_at",
		alias + "updated_at",
		alias + "deleted_at",
	}
}

// ComposeModuleEnc encodes fields from types.Module to store.Payload (map)
func (Store) ComposeModuleEnc(res *types.Module) store.Payload {
	return store.Payload{
		"id":            res.ID,
		"handle":        res.Handle,
		"name":          res.Name,
		"meta":          res.Meta,
		"rel_namespace": res.NamespaceID,
		"created_at":    res.CreatedAt,
		"updated_at":    res.UpdatedAt,
		"deleted_at":    res.DeletedAt,
	}
}
