package rdbms

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// corteza/store/users.yaml

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/corteza/store"
	"github.com/cortezaproject/corteza-server/system/types"
	"github.com/jmoiron/sqlx"
)

// SearchUsers returns all matching rows
//
// This function calls convertUserFilter with the given
// types.UserFilter and expects to receive a working squirrel.SelectBuilder
func (s Store) SearchUsers(ctx context.Context, f types.UserFilter) (types.UserSet, types.UserFilter, error) {
	query, err := s.convertUserFilter(f)
	if err != nil {
		return nil, f, err
	}

	var set = types.UserSet{}

	if f.Count, err = Count(ctx, s.db, query); err != nil || f.Count == 0 {
		return nil, f, err

	}

	return set, f, FetchPaged(ctx, s.db, query, f.PageFilter, &set)

}

// LookupUserByID searches for user by their ID
//
// It returns user even if deleted or suspended
func (s Store) LookupUserByID(ctx context.Context, id uint64) (*types.User, error) {
	return s.UserLookup(ctx, squirrel.Eq{
		"usr.id": id,
	})
}

// LookupUserByEmail searches for user by their email
//
// It returns only valid users (not deleted, not suspended)
func (s Store) LookupUserByEmail(ctx context.Context, email string) (*types.User, error) {
	return s.UserLookup(ctx, squirrel.Eq{
		"usr.email":        email,
		"usr.deleted_at":   nil,
		"usr.suspended_at": nil,
	})
}

// LookupUserByHandle searches for user by their email
//
// It returns only valid users (not deleted, not suspended)
func (s Store) LookupUserByHandle(ctx context.Context, handle string) (*types.User, error) {
	return s.UserLookup(ctx, squirrel.Eq{
		"usr.handle":       handle,
		"usr.deleted_at":   nil,
		"usr.suspended_at": nil,
	})
}

// LookupUserByUsername searches for user by their username
//
// It returns only valid users (not deleted, not suspended)
func (s Store) LookupUserByUsername(ctx context.Context, username string) (*types.User, error) {
	return s.UserLookup(ctx, squirrel.Eq{
		"usr.username":     username,
		"usr.deleted_at":   nil,
		"usr.suspended_at": nil,
	})
}

// CreateUser creates one or more rows in sys_user table
func (s Store) CreateUser(ctx context.Context, rr ...*types.User) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Insert(s.UserTable()).SetMap(s.UserEnc(res)))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// UpdateUser updates one or more existing rows in sys_user
func (s Store) UpdateUser(ctx context.Context, rr ...*types.User) error {
	return s.PartialUpdateUser(ctx, nil, rr...)
}

// PartialUpdateUser updates one or more existing rows in sys_user
//
// It wraps the update into transaction and can perform partial update by providing list of updatable columns
func (s Store) PartialUpdateUser(ctx context.Context, onlyColumns []string, rr ...*types.User) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = s.ExecUpdateUsers(
				ctx,
				squirrel.Eq{s.preprocessColumn("usr.id", ""): s.preprocessValue(res.ID, "")},
				s.UserEnc(res).Skip("id").Only(onlyColumns...))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveUser removes one or more rows from sys_user table
func (s Store) RemoveUser(ctx context.Context, rr ...*types.User) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Delete(s.UserTable("usr")).Where(squirrel.Eq{s.preprocessColumn("usr.id", ""): s.preprocessValue(res.ID, "")}))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveUserByID removes row from the sys_user table
func (s Store) RemoveUserByID(ctx context.Context, ID uint64) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Delete(s.UserTable()).Where(squirrel.Eq{s.preprocessColumn("usr.id", ""): s.preprocessValue(ID, "")}))
}

// TruncateUsers removes all rows from the sys_user table
func (s Store) TruncateUsers(ctx context.Context) error {
	return Truncate(ctx, s.DB(), s.UserTable())
}

// ExecUpdateUsers updates all matchhed (cnd) rows in sys_user with given data
func (s Store) ExecUpdateUsers(ctx context.Context, cnd squirrel.Sqlizer, set store.Payload) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Update(s.UserTable("usr")).Where(cnd).SetMap(set))
}

// UserLookup calls Lookup() and returns (if found) types.User
func (s Store) UserLookup(ctx context.Context, cnd squirrel.Sqlizer) (*types.User, error) {
	var (
		u   = &types.User{}
		err = s.Lookup(ctx, u, s.QueryUsers(), cnd)
	)

	if err == nil {
		return u, nil
	}

	return nil, err
	//return s.scanUser(s.LookupWithScan(ctx, s.QueryUsers(), cnd))
}

// QueryUsers returns squirrel.SelectBuilder with set table and all columns
func (s Store) QueryUsers() squirrel.SelectBuilder {
	return s.Select(s.UserTable("usr"), s.UserColumns("usr")...)
}

// UserTable name of the db table
func (Store) UserTable(aa ...string) string {
	var alias string
	if len(aa) > 0 {
		alias = " AS " + aa[0]
	}

	return "sys_user" + alias
}

// UserColumns returns all defined table columns
//
// With optional string arg, all columns are returned aliased
func (Store) UserColumns(aa ...string) []string {
	var alias string
	if len(aa) > 0 {
		alias = aa[0] + "."
	}

	return []string{
		alias + "id",
		alias + "email",
		alias + "email_confirmed",
		alias + "username",
		alias + "name",
		alias + "handle",
		alias + "meta",
		alias + "kind",
		alias + "created_at",
		alias + "updated_at",
		alias + "suspended_at",
		alias + "deleted_at",
	}
}

// UserEnc encodes fields from types.User to store.Payload (map)
func (Store) UserEnc(res *types.User) store.Payload {
	return store.Payload{
		"id":              res.ID,
		"email":           res.Email,
		"email_confirmed": res.EmailConfirmed,
		"username":        res.Username,
		"name":            res.Name,
		"handle":          res.Handle,
		"meta":            res.Meta,
		"kind":            res.Kind,
		"created_at":      res.CreatedAt,
		"updated_at":      res.UpdatedAt,
		"suspended_at":    res.SuspendedAt,
		"deleted_at":      res.DeletedAt,
	}
}
