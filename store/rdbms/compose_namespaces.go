package rdbms

import (
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/compose/types"
)

func (s Store) convertComposeNamespaceFilter(f types.NamespaceFilter) (query squirrel.SelectBuilder, err error) {
	query = s.QueryCredentials()

	if f.Slug != "" {
		query = query.Where(squirrel.Eq{"cns.slug": f.Slug})
	}

	return
}
