package rdbms

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// store/credentials.yaml

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/store"
	"github.com/cortezaproject/corteza-server/system/types"
	"github.com/jmoiron/sqlx"
)

// SearchCredentials returns all matching rows
//
// This function calls convertCredentialsFilter with the given
// types.CredentialsFilter and expects to receive a working squirrel.SelectBuilder
func (s Store) SearchCredentials(ctx context.Context, f types.CredentialsFilter) (types.CredentialsSet, types.CredentialsFilter, error) {
	query, err := s.convertCredentialsFilter(f)
	if err != nil {
		return nil, f, err
	}

	var set = types.CredentialsSet{}

	return set, f, FetchAll(ctx, s.db, query, &set)

}

// LookupCredentialsByID searches for credentials by their ID
//
// It returns credentials even if deleted or suspended
func (s Store) LookupCredentialsByID(ctx context.Context, id uint64) (*types.Credentials, error) {
	return s.CredentialsLookup(ctx, squirrel.Eq{
		"crd.id": id,
	})
}

// CreateCredentials creates one or more rows in sys_credentials table
func (s Store) CreateCredentials(ctx context.Context, rr ...*types.Credentials) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Insert(s.CredentialsTable()).SetMap(s.CredentialsEnc(res)))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// UpdateCredentials updates one or more existing rows in sys_credentials
func (s Store) UpdateCredentials(ctx context.Context, rr ...*types.Credentials) error {
	return s.PartialUpdateCredentials(ctx, nil, rr...)
}

// PartialUpdateCredentials updates one or more existing rows in sys_credentials
//
// It wraps the update into transaction and can perform partial update by providing list of updatable columns
func (s Store) PartialUpdateCredentials(ctx context.Context, onlyColumns []string, rr ...*types.Credentials) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = s.ExecUpdateCredentials(
				ctx,
				squirrel.Eq{s.preprocessColumn("crd.id", ""): s.preprocessValue(res.ID, "")},
				s.CredentialsEnc(res).Skip("id").Only(onlyColumns...))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveCredentials removes one or more rows from sys_credentials table
func (s Store) RemoveCredentials(ctx context.Context, rr ...*types.Credentials) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Delete(s.CredentialsTable("crd")).Where(squirrel.Eq{s.preprocessColumn("crd.id", ""): s.preprocessValue(res.ID, "")}))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveCredentialsByID removes row from the sys_credentials table
func (s Store) RemoveCredentialsByID(ctx context.Context, ID uint64) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Delete(s.CredentialsTable()).Where(squirrel.Eq{s.preprocessColumn("crd.id", ""): s.preprocessValue(ID, "")}))
}

// TruncateCredentials removes all rows from the sys_credentials table
func (s Store) TruncateCredentials(ctx context.Context) error {
	return Truncate(ctx, s.DB(), s.CredentialsTable())
}

// ExecUpdateCredentials updates all matchhed (cnd) rows in sys_credentials with given data
func (s Store) ExecUpdateCredentials(ctx context.Context, cnd squirrel.Sqlizer, set store.Payload) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Update(s.CredentialsTable("crd")).Where(cnd).SetMap(set))
}

// CredentialsLookup calls Lookup() and returns (if found) types.Credentials
func (s Store) CredentialsLookup(ctx context.Context, cnd squirrel.Sqlizer) (*types.Credentials, error) {
	var (
		u   = &types.Credentials{}
		err = s.Lookup(ctx, u, s.QueryCredentials(), cnd)
	)

	if err == nil {
		return u, nil
	}

	return nil, err
	//return s.scanCredentials(s.LookupWithScan(ctx, s.QueryCredentials(), cnd))
}

// QueryCredentials returns squirrel.SelectBuilder with set table and all columns
func (s Store) QueryCredentials() squirrel.SelectBuilder {
	return s.Select(s.CredentialsTable("crd"), s.CredentialsColumns("crd")...)
}

// CredentialsTable name of the db table
func (Store) CredentialsTable(aa ...string) string {
	var alias string
	if len(aa) > 0 {
		alias = " AS " + aa[0]
	}

	return "sys_credentials" + alias
}

// CredentialsColumns returns all defined table columns
//
// With optional string arg, all columns are returned aliased
func (Store) CredentialsColumns(aa ...string) []string {
	var alias string
	if len(aa) > 0 {
		alias = aa[0] + "."
	}

	return []string{
		alias + "id",
		alias + "rel_owner",
		alias + "kind",
		alias + "label",
		alias + "credentials",
		alias + "meta",
		alias + "last_used_at",
		alias + "expires_at",
		alias + "created_at",
		alias + "updated_at",
		alias + "deleted_at",
	}
}

// CredentialsEnc encodes fields from types.Credentials to store.Payload (map)
func (Store) CredentialsEnc(res *types.Credentials) store.Payload {
	return store.Payload{
		"id":           res.ID,
		"rel_owner":    res.OwnerID,
		"kind":         res.Kind,
		"label":        res.Label,
		"credentials":  res.Credentials,
		"meta":         res.Meta,
		"last_used_at": res.LastUsedAt,
		"expires_at":   res.ExpiresAt,
		"created_at":   res.CreatedAt,
		"updated_at":   res.UpdatedAt,
		"deleted_at":   res.DeletedAt,
	}
}
