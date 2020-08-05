package rdbms

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// store/compose_namespaces.yaml

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/store"
	"github.com/jmoiron/sqlx"
)

// SearchComposeNamespaces returns all matching rows
//
// This function calls convertComposeNamespaceFilter with the given
// types.NamespaceFilter and expects to receive a working squirrel.SelectBuilder
func (s Store) SearchComposeNamespaces(ctx context.Context, f types.NamespaceFilter) (types.NamespaceSet, types.NamespaceFilter, error) {
	query, err := s.convertComposeNamespaceFilter(f)
	if err != nil {
		return nil, f, err
	}

	var set = types.NamespaceSet{}

	if f.Count, err = Count(ctx, s.db, query); err != nil || f.Count == 0 {
		return nil, f, err

	}

	return set, f, FetchPaged(ctx, s.db, query, f.PageFilter, &set)

}

// LookupComposeNamespaceBySlug searches for namespace by slug (case-insensitive)
func (s Store) LookupComposeNamespaceBySlug(ctx context.Context, slug string) (*types.Namespace, error) {
	return s.ComposeNamespaceLookup(ctx, squirrel.Eq{
		"cns.slug": slug,
	})
}

// CreateComposeNamespace creates one or more rows in compose_namespace table
func (s Store) CreateComposeNamespace(ctx context.Context, rr ...*types.Namespace) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Insert(s.ComposeNamespaceTable()).SetMap(s.ComposeNamespaceEnc(res)))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// UpdateComposeNamespace updates one or more existing rows in compose_namespace
func (s Store) UpdateComposeNamespace(ctx context.Context, rr ...*types.Namespace) error {
	return s.PartialUpdateComposeNamespace(ctx, nil, rr...)
}

// PartialUpdateComposeNamespace updates one or more existing rows in compose_namespace
//
// It wraps the update into transaction and can perform partial update by providing list of updatable columns
func (s Store) PartialUpdateComposeNamespace(ctx context.Context, onlyColumns []string, rr ...*types.Namespace) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = s.ExecUpdateComposeNamespaces(
				ctx,
				squirrel.Eq{s.preprocessColumn("cns.id", ""): s.preprocessValue(res.ID, "")},
				s.ComposeNamespaceEnc(res).Skip("id").Only(onlyColumns...))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveComposeNamespace removes one or more rows from compose_namespace table
func (s Store) RemoveComposeNamespace(ctx context.Context, rr ...*types.Namespace) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Delete(s.ComposeNamespaceTable("cns")).Where(squirrel.Eq{s.preprocessColumn("cns.id", ""): s.preprocessValue(res.ID, "")}))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveComposeNamespaceByID removes row from the compose_namespace table
func (s Store) RemoveComposeNamespaceByID(ctx context.Context, ID uint64) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Delete(s.ComposeNamespaceTable()).Where(squirrel.Eq{s.preprocessColumn("cns.id", ""): s.preprocessValue(ID, "")}))
}

// TruncateComposeNamespaces removes all rows from the compose_namespace table
func (s Store) TruncateComposeNamespaces(ctx context.Context) error {
	return Truncate(ctx, s.DB(), s.ComposeNamespaceTable())
}

// ExecUpdateComposeNamespaces updates all matchhed (cnd) rows in compose_namespace with given data
func (s Store) ExecUpdateComposeNamespaces(ctx context.Context, cnd squirrel.Sqlizer, set store.Payload) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Update(s.ComposeNamespaceTable("cns")).Where(cnd).SetMap(set))
}

// ComposeNamespaceLookup calls Lookup() and returns (if found) types.Namespace
func (s Store) ComposeNamespaceLookup(ctx context.Context, cnd squirrel.Sqlizer) (*types.Namespace, error) {
	var (
		u   = &types.Namespace{}
		err = s.Lookup(ctx, u, s.QueryComposeNamespaces(), cnd)
	)

	if err == nil {
		return u, nil
	}

	return nil, err
	//return s.scanComposeNamespace(s.LookupWithScan(ctx, s.QueryComposeNamespaces(), cnd))
}

// QueryComposeNamespaces returns squirrel.SelectBuilder with set table and all columns
func (s Store) QueryComposeNamespaces() squirrel.SelectBuilder {
	return s.Select(s.ComposeNamespaceTable("cns"), s.ComposeNamespaceColumns("cns")...)
}

// ComposeNamespaceTable name of the db table
func (Store) ComposeNamespaceTable(aa ...string) string {
	var alias string
	if len(aa) > 0 {
		alias = " AS " + aa[0]
	}

	return "compose_namespace" + alias
}

// ComposeNamespaceColumns returns all defined table columns
//
// With optional string arg, all columns are returned aliased
func (Store) ComposeNamespaceColumns(aa ...string) []string {
	var alias string
	if len(aa) > 0 {
		alias = aa[0] + "."
	}

	return []string{
		alias + "id",
		alias + "name",
		alias + "slug",
		alias + "enabled",
		alias + "meta",
		alias + "created_at",
		alias + "updated_at",
		alias + "deleted_at",
	}
}

// ComposeNamespaceEnc encodes fields from types.Namespace to store.Payload (map)
func (Store) ComposeNamespaceEnc(res *types.Namespace) store.Payload {
	return store.Payload{
		"id":         res.ID,
		"name":       res.Name,
		"slug":       res.Slug,
		"enabled":    res.Enabled,
		"meta":       res.Meta,
		"created_at": res.CreatedAt,
		"updated_at": res.UpdatedAt,
		"deleted_at": res.DeletedAt,
	}
}