package rdbms

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// corteza/store/system_attachments.yaml

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/corteza/store"
	"github.com/cortezaproject/corteza-server/system/types"
	"github.com/jmoiron/sqlx"
)

// SearchAttachments returns all matching rows
//
// This function calls convertAttachmentFilter with the given
// types.AttachmentFilter and expects to receive a working squirrel.SelectBuilder
func (s Store) SearchAttachments(ctx context.Context, f types.AttachmentFilter) (types.AttachmentSet, types.AttachmentFilter, error) {
	query, err := s.convertAttachmentFilter(f)
	if err != nil {
		return nil, f, err
	}

	var set = types.AttachmentSet{}

	if f.Count, err = Count(ctx, s.db, query); err != nil || f.Count == 0 {
		return nil, f, err

	}

	return set, f, FetchPaged(ctx, s.db, query, f.PageFilter, &set)

}

// LookupAttachmentByID searches for attachment by its ID
//
// It returns attachment even if deleted
func (s Store) LookupAttachmentByID(ctx context.Context, id uint64) (*types.Attachment, error) {
	return s.AttachmentLookup(ctx, squirrel.Eq{
		"att.id": id,
	})
}

// CreateAttachment creates one or more rows in sys_attachment table
func (s Store) CreateAttachment(ctx context.Context, rr ...*types.Attachment) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Insert(s.AttachmentTable()).SetMap(s.AttachmentEnc(res)))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// UpdateAttachment updates one or more existing rows in sys_attachment
func (s Store) UpdateAttachment(ctx context.Context, rr ...*types.Attachment) error {
	return s.PartialUpdateAttachment(ctx, nil, rr...)
}

// PartialUpdateAttachment updates one or more existing rows in sys_attachment
//
// It wraps the update into transaction and can perform partial update by providing list of updatable columns
func (s Store) PartialUpdateAttachment(ctx context.Context, onlyColumns []string, rr ...*types.Attachment) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = s.ExecUpdateAttachments(
				ctx,
				squirrel.Eq{s.preprocessColumn("att.id", ""): s.preprocessValue(res.ID, "")},
				s.AttachmentEnc(res).Skip("id").Only(onlyColumns...))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveAttachment removes one or more rows from sys_attachment table
func (s Store) RemoveAttachment(ctx context.Context, rr ...*types.Attachment) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Delete(s.AttachmentTable("att")).Where(squirrel.Eq{s.preprocessColumn("att.id", ""): s.preprocessValue(res.ID, "")}))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveAttachmentByID removes row from the sys_attachment table
func (s Store) RemoveAttachmentByID(ctx context.Context, ID uint64) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Delete(s.AttachmentTable()).Where(squirrel.Eq{s.preprocessColumn("att.id", ""): s.preprocessValue(ID, "")}))
}

// TruncateAttachments removes all rows from the sys_attachment table
func (s Store) TruncateAttachments(ctx context.Context) error {
	return Truncate(ctx, s.DB(), s.AttachmentTable())
}

// ExecUpdateAttachments updates all matchhed (cnd) rows in sys_attachment with given data
func (s Store) ExecUpdateAttachments(ctx context.Context, cnd squirrel.Sqlizer, set store.Payload) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Update(s.AttachmentTable("att")).Where(cnd).SetMap(set))
}

// AttachmentLookup calls Lookup() and returns (if found) types.Attachment
func (s Store) AttachmentLookup(ctx context.Context, cnd squirrel.Sqlizer) (*types.Attachment, error) {
	var (
		u   = &types.Attachment{}
		err = s.Lookup(ctx, u, s.QueryAttachments(), cnd)
	)

	if err == nil {
		return u, nil
	}

	return nil, err
	//return s.scanAttachment(s.LookupWithScan(ctx, s.QueryAttachments(), cnd))
}

// QueryAttachments returns squirrel.SelectBuilder with set table and all columns
func (s Store) QueryAttachments() squirrel.SelectBuilder {
	return s.Select(s.AttachmentTable("att"), s.AttachmentColumns("att")...)
}

// AttachmentTable name of the db table
func (Store) AttachmentTable(aa ...string) string {
	var alias string
	if len(aa) > 0 {
		alias = " AS " + aa[0]
	}

	return "sys_attachment" + alias
}

// AttachmentColumns returns all defined table columns
//
// With optional string arg, all columns are returned aliased
func (Store) AttachmentColumns(aa ...string) []string {
	var alias string
	if len(aa) > 0 {
		alias = aa[0] + "."
	}

	return []string{
		alias + "id",
		alias + "rel_owner",
		alias + "kind",
		alias + "url",
		alias + "preview_url",
		alias + "name",
		alias + "meta",
		alias + "created_at",
		alias + "updated_at",
		alias + "deleted_at",
	}
}

// AttachmentEnc encodes fields from types.Attachment to store.Payload (map)
func (Store) AttachmentEnc(res *types.Attachment) store.Payload {
	return store.Payload{
		"id":          res.ID,
		"rel_owner":   res.OwnerID,
		"kind":        res.Kind,
		"url":         res.Url,
		"preview_url": res.PreviewUrl,
		"name":        res.Name,
		"meta":        res.Meta,
		"created_at":  res.CreatedAt,
		"updated_at":  res.UpdatedAt,
		"deleted_at":  res.DeletedAt,
	}
}
