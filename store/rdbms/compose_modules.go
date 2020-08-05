package rdbms

import (
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/pkg/rh"
	"strings"
)

func (s Store) convertComposeModuleFilter(f types.ModuleFilter) (query squirrel.SelectBuilder, err error) {
	if f.Sort == "" {
		f.Sort = "id ASC"
	}

	query = s.QueryComposeModules()

	if f.NamespaceID > 0 {
		query = query.Where("cmd.rel_namespace = ?", f.NamespaceID)
	}

	if f.Query != "" {
		q := "%" + strings.ToLower(f.Query) + "%"
		query = query.Where(squirrel.Or{
			squirrel.Like{"LOWER(cmd.name)": q},
			squirrel.Like{"LOWER(cmd.handle)": q},
		})
	}

	if f.Name != "" {
		query = query.Where(squirrel.Eq{"LOWER(cmd.name)": strings.ToLower(f.Name)})
	}

	if f.Handle != "" {
		query = query.Where(squirrel.Eq{"LOWER(cmd.handle)": strings.ToLower(f.Handle)})
	}

	if f.IsReadable != nil {
		query = query.Where(f.IsReadable)
	}

	var orderBy []string
	if orderBy, err = rh.ParseOrder(f.Sort, s.ComposeModuleColumns()...); err != nil {
		return
	} else {
		query = query.OrderBy(orderBy...)
	}

	return
}