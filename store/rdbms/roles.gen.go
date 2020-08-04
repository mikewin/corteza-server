package rdbms

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// store/roles.yaml

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/store"
	"github.com/cortezaproject/corteza-server/system/types"
	"github.com/jmoiron/sqlx"
)

// SearchRoles returns all matching rows
//
// This function calls convertRoleFilter with the given
// types.RoleFilter and expects to receive a working squirrel.SelectBuilder
func (s Store) SearchRoles(ctx context.Context, f types.RoleFilter) (types.RoleSet, types.RoleFilter, error) {
	query, err := s.convertRoleFilter(f)
	if err != nil {
		return nil, f, err
	}

	var set = types.RoleSet{}

	if f.Count, err = Count(ctx, s.db, query); err != nil || f.Count == 0 {
		return nil, f, err

	}

	return set, f, FetchPaged(ctx, s.db, query, f.PageFilter, &set)

}

// LookupRoleByID searches for role by their ID
//
// It returns role even if deleted or suspended
func (s Store) LookupRoleByID(ctx context.Context, id uint64) (*types.Role, error) {
	return s.RoleLookup(ctx, squirrel.Eq{
		"rl.id": id,
	})
}

// LookupRoleByHandle searches for role by its handle
//
// It returns only valid roles (not deleted, not archived)
func (s Store) LookupRoleByHandle(ctx context.Context, handle string) (*types.Role, error) {
	return s.RoleLookup(ctx, squirrel.Eq{
		"rl.handle":      handle,
		"rl.archived_at": nil,
		"rl.deleted_at":  nil,
	})
}

// LookupRoleByName searches for role by its name
//
// It returns only valid roles (not deleted, not archived)
func (s Store) LookupRoleByName(ctx context.Context, name string) (*types.Role, error) {
	return s.RoleLookup(ctx, squirrel.Eq{
		"rl.name":        name,
		"rl.archived_at": nil,
		"rl.deleted_at":  nil,
	})
}

// CreateRole creates one or more rows in sys_role table
func (s Store) CreateRole(ctx context.Context, rr ...*types.Role) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Insert(s.RoleTable()).SetMap(s.RoleEnc(res)))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// UpdateRole updates one or more existing rows in sys_role
func (s Store) UpdateRole(ctx context.Context, rr ...*types.Role) error {
	return s.PartialUpdateRole(ctx, nil, rr...)
}

// PartialUpdateRole updates one or more existing rows in sys_role
//
// It wraps the update into transaction and can perform partial update by providing list of updatable columns
func (s Store) PartialUpdateRole(ctx context.Context, onlyColumns []string, rr ...*types.Role) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = s.ExecUpdateRoles(
				ctx,
				squirrel.Eq{s.preprocessColumn("rl.id", ""): s.preprocessValue(res.ID, "")},
				s.RoleEnc(res).Skip("id").Only(onlyColumns...))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveRole removes one or more rows from sys_role table
func (s Store) RemoveRole(ctx context.Context, rr ...*types.Role) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Delete(s.RoleTable("rl")).Where(squirrel.Eq{s.preprocessColumn("rl.id", ""): s.preprocessValue(res.ID, "")}))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveRoleByID removes row from the sys_role table
func (s Store) RemoveRoleByID(ctx context.Context, ID uint64) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Delete(s.RoleTable()).Where(squirrel.Eq{s.preprocessColumn("rl.id", ""): s.preprocessValue(ID, "")}))
}

// TruncateRoles removes all rows from the sys_role table
func (s Store) TruncateRoles(ctx context.Context) error {
	return Truncate(ctx, s.DB(), s.RoleTable())
}

// ExecUpdateRoles updates all matchhed (cnd) rows in sys_role with given data
func (s Store) ExecUpdateRoles(ctx context.Context, cnd squirrel.Sqlizer, set store.Payload) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Update(s.RoleTable("rl")).Where(cnd).SetMap(set))
}

// RoleLookup calls Lookup() and returns (if found) types.Role
func (s Store) RoleLookup(ctx context.Context, cnd squirrel.Sqlizer) (*types.Role, error) {
	var (
		u   = &types.Role{}
		err = s.Lookup(ctx, u, s.QueryRoles(), cnd)
	)

	if err == nil {
		return u, nil
	}

	return nil, err
	//return s.scanRole(s.LookupWithScan(ctx, s.QueryRoles(), cnd))
}

// QueryRoles returns squirrel.SelectBuilder with set table and all columns
func (s Store) QueryRoles() squirrel.SelectBuilder {
	return s.Select(s.RoleTable("rl"), s.RoleColumns("rl")...)
}

// RoleTable name of the db table
func (Store) RoleTable(aa ...string) string {
	var alias string
	if len(aa) > 0 {
		alias = " AS " + aa[0]
	}

	return "sys_role" + alias
}

// RoleColumns returns all defined table columns
//
// With optional string arg, all columns are returned aliased
func (Store) RoleColumns(aa ...string) []string {
	var alias string
	if len(aa) > 0 {
		alias = aa[0] + "."
	}

	return []string{
		alias + "id",
		alias + "name",
		alias + "handle",
		alias + "archived_at",
		alias + "created_at",
		alias + "updated_at",
		alias + "deleted_at",
	}
}

// RoleEnc encodes fields from types.Role to store.Payload (map)
func (Store) RoleEnc(res *types.Role) store.Payload {
	return store.Payload{
		"id":          res.ID,
		"name":        res.Name,
		"handle":      res.Handle,
		"archived_at": res.ArchivedAt,
		"created_at":  res.CreatedAt,
		"updated_at":  res.UpdatedAt,
		"deleted_at":  res.DeletedAt,
	}
}
