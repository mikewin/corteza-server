package rdbms

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// store/compose_charts.yaml

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/store"
	"github.com/jmoiron/sqlx"
)

// SearchComposeCharts returns all matching rows
//
// This function calls convertComposeChartFilter with the given
// types.ChartFilter and expects to receive a working squirrel.SelectBuilder
func (s Store) SearchComposeCharts(ctx context.Context, f types.ChartFilter) (types.ChartSet, types.ChartFilter, error) {
	query, err := s.convertComposeChartFilter(f)
	if err != nil {
		return nil, f, err
	}

	var set = types.ChartSet{}

	if f.Count, err = Count(ctx, s.db, query); err != nil || f.Count == 0 {
		return nil, f, err

	}

	return set, f, FetchPaged(ctx, s.db, query, f.PageFilter, &set)

}

// LookupComposeChartByHandle searches for compose chart by handle (case-insensitive)
func (s Store) LookupComposeChartByHandle(ctx context.Context, handle string) (*types.Chart, error) {
	return s.ComposeChartLookup(ctx, squirrel.Eq{
		"cch.handle": handle,
	})
}

// CreateComposeChart creates one or more rows in compose_chart table
func (s Store) CreateComposeChart(ctx context.Context, rr ...*types.Chart) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Insert(s.ComposeChartTable()).SetMap(s.ComposeChartEnc(res)))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// UpdateComposeChart updates one or more existing rows in compose_chart
func (s Store) UpdateComposeChart(ctx context.Context, rr ...*types.Chart) error {
	return s.PartialUpdateComposeChart(ctx, nil, rr...)
}

// PartialUpdateComposeChart updates one or more existing rows in compose_chart
//
// It wraps the update into transaction and can perform partial update by providing list of updatable columns
func (s Store) PartialUpdateComposeChart(ctx context.Context, onlyColumns []string, rr ...*types.Chart) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = s.ExecUpdateComposeCharts(
				ctx,
				squirrel.Eq{s.preprocessColumn("cch.id", ""): s.preprocessValue(res.ID, "")},
				s.ComposeChartEnc(res).Skip("id").Only(onlyColumns...))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveComposeChart removes one or more rows from compose_chart table
func (s Store) RemoveComposeChart(ctx context.Context, rr ...*types.Chart) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Delete(s.ComposeChartTable("cch")).Where(squirrel.Eq{s.preprocessColumn("cch.id", ""): s.preprocessValue(res.ID, "")}))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveComposeChartByID removes row from the compose_chart table
func (s Store) RemoveComposeChartByID(ctx context.Context, ID uint64) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Delete(s.ComposeChartTable()).Where(squirrel.Eq{s.preprocessColumn("cch.id", ""): s.preprocessValue(ID, "")}))
}

// TruncateComposeCharts removes all rows from the compose_chart table
func (s Store) TruncateComposeCharts(ctx context.Context) error {
	return Truncate(ctx, s.DB(), s.ComposeChartTable())
}

// ExecUpdateComposeCharts updates all matchhed (cnd) rows in compose_chart with given data
func (s Store) ExecUpdateComposeCharts(ctx context.Context, cnd squirrel.Sqlizer, set store.Payload) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Update(s.ComposeChartTable("cch")).Where(cnd).SetMap(set))
}

// ComposeChartLookup calls Lookup() and returns (if found) types.Chart
func (s Store) ComposeChartLookup(ctx context.Context, cnd squirrel.Sqlizer) (*types.Chart, error) {
	var (
		u   = &types.Chart{}
		err = s.Lookup(ctx, u, s.QueryComposeCharts(), cnd)
	)

	if err == nil {
		return u, nil
	}

	return nil, err
	//return s.scanComposeChart(s.LookupWithScan(ctx, s.QueryComposeCharts(), cnd))
}

// QueryComposeCharts returns squirrel.SelectBuilder with set table and all columns
func (s Store) QueryComposeCharts() squirrel.SelectBuilder {
	return s.Select(s.ComposeChartTable("cch"), s.ComposeChartColumns("cch")...)
}

// ComposeChartTable name of the db table
func (Store) ComposeChartTable(aa ...string) string {
	var alias string
	if len(aa) > 0 {
		alias = " AS " + aa[0]
	}

	return "compose_chart" + alias
}

// ComposeChartColumns returns all defined table columns
//
// With optional string arg, all columns are returned aliased
func (Store) ComposeChartColumns(aa ...string) []string {
	var alias string
	if len(aa) > 0 {
		alias = aa[0] + "."
	}

	return []string{
		alias + "id",
		alias + "handle",
		alias + "name",
		alias + "config",
		alias + "rel_namespace",
		alias + "created_at",
		alias + "updated_at",
		alias + "deleted_at",
	}
}

// ComposeChartEnc encodes fields from types.Chart to store.Payload (map)
func (Store) ComposeChartEnc(res *types.Chart) store.Payload {
	return store.Payload{
		"id":            res.ID,
		"handle":        res.Handle,
		"name":          res.Name,
		"config":        res.Config,
		"rel_namespace": res.NamespaceID,
		"created_at":    res.CreatedAt,
		"updated_at":    res.UpdatedAt,
		"deleted_at":    res.DeletedAt,
	}
}
