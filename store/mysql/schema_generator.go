package mysql

import (
	"fmt"
	"github.com/cortezaproject/corteza-server/store/rdbms/schema"
	"strings"
)

type (
	mysqlTableCreator struct {
		t *schema.Table
	}
)

func NewMysqlTableCreator(t *schema.Table) *mysqlTableCreator {
	return &mysqlTableCreator{t: t}
}

func (tc *mysqlTableCreator) Make() []string {
	var (
		cc = make([]string, 0)
	)

	for _, c := range tc.t.Columns {
		cc = append(cc, fmt.Sprintf(
			"%-20s %-20s %-20s %-20s",
			c.Name,
			tc.mysqlColumnType(c.Type),
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
			"CREATE TABLE %s (\n   %s\n) ENGINE=InnoDB DEFAULT CHARSET=utf8\n",
			tc.t.Name,
			strings.Join(cc, ",\n   "),
		),
	}
}

func (ct *mysqlTableCreator) mysqlColumnType(c schema.ColumnType) string {
	switch c.Type {
	case schema.ColumnTypeIdentifier:
		return "BIGINT UNSIGNED"
	case schema.ColumnTypeText:
		if c.Length > 0 {
			return fmt.Sprintf("VARCHAR(%d)", c.Length)
		}
		return "TEXT"
	case schema.ColumnTypeTimestamp:
		return "DATETIME"
	case schema.ColumnTypeInteger:
		return "INTEGER"
	case schema.ColumnTypeJson:
		return "JSON"
	case schema.ColumnTypeBoolean:
		return "TINYINT(1)"
	}

	panic(fmt.Sprintf("unhandled column type: %d on table %s", c.Type, ct.t.Name))
}

func (c *mysqlTableCreator) createTablePrimaryKey(pk *schema.Index) string {
	return fmt.Sprintf("PRIMARY KEY (%s)", strings.Join(pk.Columns, ","))
}
