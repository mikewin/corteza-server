package pgsql

import (
	"fmt"
	"github.com/cortezaproject/corteza-server/corteza/store/rdbms/schema"
	"strings"
)

type (
	pgsqlTableCreator struct {
		t *schema.Table
	}
)

func NewPgsqlTableCreator(t *schema.Table) *pgsqlTableCreator {
	return &pgsqlTableCreator{t: t}
}

func (tc *pgsqlTableCreator) Make() []string {
	var (
		cc = make([]string, 0)
	)

	for _, c := range tc.t.Columns {
		cc = append(cc, fmt.Sprintf(
			"%-20s %-20s %-20s %-20s",
			c.Name,
			tc.pgsqlColumnType(c.Type),
			func() string {
				if !c.IsNull {
					return "NOT NULL"
				} else {
					return ""

				}
			}(),
			func() string {
				if len(c.DefaultValue) > 0 {
					return "DEFAULT " + c.DefaultValue
				} else {
					return ""
				}
			}(),
		))
	}

	if tc.t.Primary != nil {
		cc = append(cc, tc.createTablePrimaryKey(tc.t.Primary))
	}

	return []string{
		fmt.Sprintf(
			"CREATE TABLE %s (\n   %s\n) WITHOUT OIDS\n",
			tc.t.Name,
			strings.Join(cc, ",\n   "),
		),
	}
}

func (ct *pgsqlTableCreator) pgsqlColumnType(c schema.ColumnType) string {
	switch c.Type {
	case schema.ColumnTypeIdentifier:
		return "BIGINT"
	case schema.ColumnTypeText:
		if c.Length > 0 {
			return fmt.Sprintf("VARCHAR(%d)", c.Length)
		}
		return "TEXT"
	case schema.ColumnTypeTimestamp:
		return fmt.Sprintf("TIMESTAMPTZ(%d)", c.Length)
	case schema.ColumnTypeInteger:
		return "INTEGER"
	case schema.ColumnTypeJson:
		return "JSON"
	case schema.ColumnTypeBoolean:
		return "BOOLEAN"
	}

	panic(fmt.Sprintf("unhandled column type: %d on table %s", c.Type, ct.t.Name))
}

func (c *pgsqlTableCreator) createTablePrimaryKey(pk *schema.Index) string {
	return fmt.Sprintf("PRIMARY KEY (%s)", strings.Join(pk.Columns, ","))
}
