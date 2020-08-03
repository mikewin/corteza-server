package rdbms

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// corteza/store/compose_pages.yaml

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/corteza/store"
	"github.com/jmoiron/sqlx"
)

// SearchComposePages returns all matching rows
//
// This function calls convertComposePageFilter with the given
// types.PageFilter and expects to receive a working squirrel.SelectBuilder
func (s Store) SearchComposePages(ctx context.Context, f types.PageFilter) (types.PageSet, types.PageFilter, error) {
	query, err := s.convertComposePageFilter(f)
	if err != nil {
		return nil, f, err
	}

	var set = types.PageSet{}

	if f.Count, err = Count(ctx, s.db, query); err != nil || f.Count == 0 {
		return nil, f, err

	}

	return set, f, FetchPaged(ctx, s.db, query, f.PageFilter, &set)

}

// LookupComposePageByHandle searches for page chart by handle (case-insensitive)
func (s Store) LookupComposePageByHandle(ctx context.Context, handle string) (*types.Page, error) {
	return s.ComposePageLookup(ctx, squirrel.Eq{
		"cpg.handle": handle,
	})
}

// CreateComposePage creates one or more rows in compose_page table
func (s Store) CreateComposePage(ctx context.Context, rr ...*types.Page) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Insert(s.ComposePageTable()).SetMap(s.ComposePageEnc(res)))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// UpdateComposePage updates one or more existing rows in compose_page
func (s Store) UpdateComposePage(ctx context.Context, rr ...*types.Page) error {
	return s.PartialUpdateComposePage(ctx, nil, rr...)
}

// PartialUpdateComposePage updates one or more existing rows in compose_page
//
// It wraps the update into transaction and can perform partial update by providing list of updatable columns
func (s Store) PartialUpdateComposePage(ctx context.Context, onlyColumns []string, rr ...*types.Page) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = s.ExecUpdateComposePages(
				ctx,
				squirrel.Eq{s.preprocessColumn("cpg.id", ""): s.preprocessValue(res.ID, "")},
				s.ComposePageEnc(res).Skip("id").Only(onlyColumns...))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveComposePage removes one or more rows from compose_page table
func (s Store) RemoveComposePage(ctx context.Context, rr ...*types.Page) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Delete(s.ComposePageTable("cpg")).Where(squirrel.Eq{s.preprocessColumn("cpg.id", ""): s.preprocessValue(res.ID, "")}))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveComposePageByID removes row from the compose_page table
func (s Store) RemoveComposePageByID(ctx context.Context, ID uint64) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Delete(s.ComposePageTable()).Where(squirrel.Eq{s.preprocessColumn("cpg.id", ""): s.preprocessValue(ID, "")}))
}

// TruncateComposePages removes all rows from the compose_page table
func (s Store) TruncateComposePages(ctx context.Context) error {
	return Truncate(ctx, s.DB(), s.ComposePageTable())
}

// ExecUpdateComposePages updates all matchhed (cnd) rows in compose_page with given data
func (s Store) ExecUpdateComposePages(ctx context.Context, cnd squirrel.Sqlizer, set store.Payload) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Update(s.ComposePageTable("cpg")).Where(cnd).SetMap(set))
}

// ComposePageLookup calls Lookup() and returns (if found) types.Page
func (s Store) ComposePageLookup(ctx context.Context, cnd squirrel.Sqlizer) (*types.Page, error) {
	var (
		u   = &types.Page{}
		err = s.Lookup(ctx, u, s.QueryComposePages(), cnd)
	)

	if err == nil {
		return u, nil
	}

	return nil, err
	//return s.scanComposePage(s.LookupWithScan(ctx, s.QueryComposePages(), cnd))
}

// QueryComposePages returns squirrel.SelectBuilder with set table and all columns
func (s Store) QueryComposePages() squirrel.SelectBuilder {
	return s.Select(s.ComposePageTable("cpg"), s.ComposePageColumns("cpg")...)
}

// ComposePageTable name of the db table
func (Store) ComposePageTable(aa ...string) string {
	var alias string
	if len(aa) > 0 {
		alias = " AS " + aa[0]
	}

	return "compose_page" + alias
}

// ComposePageColumns returns all defined table columns
//
// With optional string arg, all columns are returned aliased
func (Store) ComposePageColumns(aa ...string) []string {
	var alias string
	if len(aa) > 0 {
		alias = aa[0] + "."
	}

	return []string{
		alias + "id",
		alias + "self_id",
		alias + "rel_namespace",
		alias + "rel_module",
		alias + "handle",
		alias + "title",
		alias + "description",
		alias + "blocks",
		alias + "visible",
		alias + "weight",
		alias + "created_at",
		alias + "updated_at",
		alias + "deleted_at",
	}
}

// ComposePageEnc encodes fields from types.Page to store.Payload (map)
func (Store) ComposePageEnc(res *types.Page) store.Payload {
	return store.Payload{
		"id":            res.ID,
		"self_id":       res.SelfID,
		"rel_namespace": res.NamespaceID,
		"rel_module":    res.ModuleID,
		"handle":        res.Handle,
		"title":         res.Title,
		"description":   res.Description,
		"blocks":        res.Blocks,
		"visible":       res.Visible,
		"weight":        res.Weight,
		"created_at":    res.CreatedAt,
		"updated_at":    res.UpdatedAt,
		"deleted_at":    res.DeletedAt,
	}
}
