package rdbms

import (
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/pkg/rh"
	"strings"
)

func (s Store) convertComposePageFilter(f types.PageFilter) (query squirrel.SelectBuilder, err error) {
	if f.Sort == "" {
		f.Sort = "id ASC"
	}

	query = s.QueryComposePages()

	if f.NamespaceID > 0 {
		query = query.Where("cch.rel_namespace = ?", f.NamespaceID)
	}

	if f.Query != "" {
		q := "%" + strings.ToLower(f.Query) + "%"
		query = query.Where(squirrel.Or{
			squirrel.Like{"LOWER(cch.handle)": q},
		})
	}

	if f.Handle != "" {
		query = query.Where(squirrel.Eq{"LOWER(cch.handle)": strings.ToLower(f.Handle)})
	}

	if f.IsReadable != nil {
		query = query.Where(f.IsReadable)
	}

	var orderBy []string
	if orderBy, err = rh.ParseOrder(f.Sort, s.ComposePageColumns()...); err != nil {
		return
	} else {
		query = query.OrderBy(orderBy...)
	}

	return
}
