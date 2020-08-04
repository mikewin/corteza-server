package rdbms

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// store/settings.yaml

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/pkg/settings"
	"github.com/cortezaproject/corteza-server/store"
	"github.com/jmoiron/sqlx"
)

// SearchSettings returns all matching rows
//
// This function calls convertSettingFilter with the given
// settings.Filter and expects to receive a working squirrel.SelectBuilder
func (s Store) SearchSettings(ctx context.Context, f settings.Filter) (settings.ValueSet, settings.Filter, error) {
	query, err := s.convertSettingFilter(f)
	if err != nil {
		return nil, f, err
	}

	var set = settings.ValueSet{}

	if f.Count, err = Count(ctx, s.db, query); err != nil || f.Count == 0 {
		return nil, f, err

	}

	return set, f, FetchPaged(ctx, s.db, query, f.PageFilter, &set)

}

// LookupSettingByNameOwnedBy searches for settings by name and owner
func (s Store) LookupSettingByNameOwnedBy(ctx context.Context, name string, owned_by uint64) (*settings.Value, error) {
	return s.SettingLookup(ctx, squirrel.Eq{
		"st.name":     name,
		"st.owned_by": owned_by,
	})
}

// CreateSetting creates one or more rows in settings table
func (s Store) CreateSetting(ctx context.Context, rr ...*settings.Value) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Insert(s.SettingTable()).SetMap(s.SettingEnc(res)))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// UpdateSetting updates one or more existing rows in settings
func (s Store) UpdateSetting(ctx context.Context, rr ...*settings.Value) error {
	return s.PartialUpdateSetting(ctx, nil, rr...)
}

// PartialUpdateSetting updates one or more existing rows in settings
//
// It wraps the update into transaction and can perform partial update by providing list of updatable columns
func (s Store) PartialUpdateSetting(ctx context.Context, onlyColumns []string, rr ...*settings.Value) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = s.ExecUpdateSettings(
				ctx,
				squirrel.Eq{s.preprocessColumn("st.name", ""): s.preprocessValue(res.Name, ""),
					s.preprocessColumn("st.owned_by", ""): s.preprocessValue(res.OwnedBy, ""),
				},
				s.SettingEnc(res).Skip("name", "owned_by").Only(onlyColumns...))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveSetting removes one or more rows from settings table
func (s Store) RemoveSetting(ctx context.Context, rr ...*settings.Value) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Delete(s.SettingTable("st")).Where(squirrel.Eq{s.preprocessColumn("st.name", ""): s.preprocessValue(res.Name, ""),
				s.preprocessColumn("st.owned_by", ""): s.preprocessValue(res.OwnedBy, ""),
			}))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveSettingByNameOwnedBy removes row from the settings table
func (s Store) RemoveSettingByNameOwnedBy(ctx context.Context, name string, ownedBy uint64) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Delete(s.SettingTable()).Where(squirrel.Eq{s.preprocessColumn("st.name", ""): s.preprocessValue(name, ""),

		s.preprocessColumn("st.owned_by", ""): s.preprocessValue(ownedBy, ""),
	}))
}

// TruncateSettings removes all rows from the settings table
func (s Store) TruncateSettings(ctx context.Context) error {
	return Truncate(ctx, s.DB(), s.SettingTable())
}

// ExecUpdateSettings updates all matchhed (cnd) rows in settings with given data
func (s Store) ExecUpdateSettings(ctx context.Context, cnd squirrel.Sqlizer, set store.Payload) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Update(s.SettingTable("st")).Where(cnd).SetMap(set))
}

// SettingLookup calls Lookup() and returns (if found) settings.Value
func (s Store) SettingLookup(ctx context.Context, cnd squirrel.Sqlizer) (*settings.Value, error) {
	var (
		u   = &settings.Value{}
		err = s.Lookup(ctx, u, s.QuerySettings(), cnd)
	)

	if err == nil {
		return u, nil
	}

	return nil, err
	//return s.scanSetting(s.LookupWithScan(ctx, s.QuerySettings(), cnd))
}

// QuerySettings returns squirrel.SelectBuilder with set table and all columns
func (s Store) QuerySettings() squirrel.SelectBuilder {
	return s.Select(s.SettingTable("st"), s.SettingColumns("st")...)
}

// SettingTable name of the db table
func (Store) SettingTable(aa ...string) string {
	var alias string
	if len(aa) > 0 {
		alias = " AS " + aa[0]
	}

	return "settings" + alias
}

// SettingColumns returns all defined table columns
//
// With optional string arg, all columns are returned aliased
func (Store) SettingColumns(aa ...string) []string {
	var alias string
	if len(aa) > 0 {
		alias = aa[0] + "."
	}

	return []string{
		alias + "name",
		alias + "value",
		alias + "owned_by",
		alias + "updated_by",
		alias + "updated_at",
	}
}

// SettingEnc encodes fields from settings.Value to store.Payload (map)
func (Store) SettingEnc(res *settings.Value) store.Payload {
	return store.Payload{
		"name":       res.Name,
		"value":      res.Value,
		"owned_by":   res.OwnedBy,
		"updated_by": res.UpdatedBy,
		"updated_at": res.UpdatedAt,
	}
}
