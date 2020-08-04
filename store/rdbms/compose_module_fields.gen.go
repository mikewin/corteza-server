package rdbms

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// store/compose_module_fields.yaml

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/store"
	"github.com/jmoiron/sqlx"
)

// CreateComposeModuleField creates one or more rows in compose_module_field table
func (s Store) CreateComposeModuleField(ctx context.Context, rr ...*types.ModuleField) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Insert(s.ComposeModuleFieldTable()).SetMap(s.ComposeModuleFieldEnc(res)))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// UpdateComposeModuleField updates one or more existing rows in compose_module_field
func (s Store) UpdateComposeModuleField(ctx context.Context, rr ...*types.ModuleField) error {
	return s.PartialUpdateComposeModuleField(ctx, nil, rr...)
}

// PartialUpdateComposeModuleField updates one or more existing rows in compose_module_field
//
// It wraps the update into transaction and can perform partial update by providing list of updatable columns
func (s Store) PartialUpdateComposeModuleField(ctx context.Context, onlyColumns []string, rr ...*types.ModuleField) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = s.ExecUpdateComposeModuleFields(
				ctx,
				squirrel.Eq{s.preprocessColumn("cmf.name", ""): s.preprocessValue(res.Name, "")},
				s.ComposeModuleFieldEnc(res).Skip("name").Only(onlyColumns...))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveComposeModuleField removes one or more rows from compose_module_field table
func (s Store) RemoveComposeModuleField(ctx context.Context, rr ...*types.ModuleField) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Delete(s.ComposeModuleFieldTable("cmf")).Where(squirrel.Eq{s.preprocessColumn("cmf.name", ""): s.preprocessValue(res.Name, "")}))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveComposeModuleFieldByName removes row from the compose_module_field table
func (s Store) RemoveComposeModuleFieldByName(ctx context.Context, name string) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Delete(s.ComposeModuleFieldTable()).Where(squirrel.Eq{

		s.preprocessColumn("cmf.name", ""): s.preprocessValue(name, ""),
	}))
}

// TruncateComposeModuleFields removes all rows from the compose_module_field table
func (s Store) TruncateComposeModuleFields(ctx context.Context) error {
	return Truncate(ctx, s.DB(), s.ComposeModuleFieldTable())
}

// ExecUpdateComposeModuleFields updates all matchhed (cnd) rows in compose_module_field with given data
func (s Store) ExecUpdateComposeModuleFields(ctx context.Context, cnd squirrel.Sqlizer, set store.Payload) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Update(s.ComposeModuleFieldTable("cmf")).Where(cnd).SetMap(set))
}

// ComposeModuleFieldLookup calls Lookup() and returns (if found) types.ModuleField
func (s Store) ComposeModuleFieldLookup(ctx context.Context, cnd squirrel.Sqlizer) (*types.ModuleField, error) {
	var (
		u   = &types.ModuleField{}
		err = s.Lookup(ctx, u, s.QueryComposeModuleFields(), cnd)
	)

	if err == nil {
		return u, nil
	}

	return nil, err
	//return s.scanComposeModuleField(s.LookupWithScan(ctx, s.QueryComposeModuleFields(), cnd))
}

// QueryComposeModuleFields returns squirrel.SelectBuilder with set table and all columns
func (s Store) QueryComposeModuleFields() squirrel.SelectBuilder {
	return s.Select(s.ComposeModuleFieldTable("cmf"), s.ComposeModuleFieldColumns("cmf")...)
}

// ComposeModuleFieldTable name of the db table
func (Store) ComposeModuleFieldTable(aa ...string) string {
	var alias string
	if len(aa) > 0 {
		alias = " AS " + aa[0]
	}

	return "compose_module_field" + alias
}

// ComposeModuleFieldColumns returns all defined table columns
//
// With optional string arg, all columns are returned aliased
func (Store) ComposeModuleFieldColumns(aa ...string) []string {
	var alias string
	if len(aa) > 0 {
		alias = aa[0] + "."
	}

	return []string{
		alias + "rel_module",
		alias + "place",
		alias + "kind",
		alias + "name",
		alias + "label",
		alias + "options",
		alias + "private",
		alias + "required",
		alias + "visible",
		alias + "multi",
		alias + "default_value",
		alias + "created_at",
		alias + "updated_at",
		alias + "deleted_at",
	}
}

// ComposeModuleFieldEnc encodes fields from types.ModuleField to store.Payload (map)
func (Store) ComposeModuleFieldEnc(res *types.ModuleField) store.Payload {
	return store.Payload{
		"rel_module":    res.ModuleID,
		"place":         res.Place,
		"kind":          res.Kind,
		"name":          res.Name,
		"label":         res.Label,
		"options":       res.Options,
		"private":       res.Private,
		"required":      res.Required,
		"visible":       res.Visible,
		"multi":         res.Multi,
		"default_value": res.DefaultValue,
		"created_at":    res.CreatedAt,
		"updated_at":    res.UpdatedAt,
		"deleted_at":    res.DeletedAt,
	}
}
