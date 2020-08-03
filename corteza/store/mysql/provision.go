package mysql

// MySQL specific prefixes, sql
// templates, functions and other helpers

import (
	"fmt"
	. "github.com/cortezaproject/corteza-server/corteza/store/provisioner"
	"github.com/cortezaproject/corteza-server/corteza/store/rdbms/schema"
	_ "github.com/go-sql-driver/mysql"
)

type (
	// Holds table structure
	tableColumn struct {
		Field   string  `db:"Field"`
		Type    string  `db:"Type"`
		Null    string  `db:"Null"`
		Key     string  `db:"Key"`
		Default *string `db:"Default"`
		Extra   string  `db:"Extra"`
	}

	// storeProvision groups all provisioning functions
	storeProvision struct {
		*Store
	}
)

// Engine, charset are used on every mysql table
const (
	pfxCreateTable = `ENGINE=InnoDB DEFAULT CHARSET=utf8`
	sqlTableExists = `SELECT COUNT(*) FROM information_schema.TABLES WHERE (TABLE_SCHEMA = ?) AND (TABLE_NAME = ?)`
	fmtDropColumn  = `ALTER TABLE %s DROP COLUMN %s`
	fmtAddColumn   = `ALTER TABLE %s ADD COLUMN %s %s`
)

// utility to simplify table creation
func (s storeProvision) createTable(def *schema.Table, ifFalse ...Executor) Executor {
	var mysqlMaker = NewMysqlTableCreator(def)

	return Do(
		Label("provisioning mysql database table "+def.Name),
		IfElse(
			s.tableMissing(def.Name),
			Do(s.execSql(mysqlMaker.Make()...), Label("created\n")),
			Do(ifFalse...),
		),
	)
}

// Returns Tester fn that will
// verify if table is present or missing
func (s storeProvision) tableMissing(table string) Tester {
	return func(p *Provisioner) (bool, error) {
		// @todo implement
		var count int
		if err := s.DB().Get(&count, sqlTableExists, s.Config().DBName, table); err != nil {
			return false, err
		} else {
			return count == 0, nil
		}
	}
}

// Returns Executor fn that removes column (if exists) from a table
func (s storeProvision) dropColumn(table, column string) Executor {
	return func(p *Provisioner) error {
		if tt, err := s.getTableColumns(table); err != nil || s.getColumn(tt, column) == nil {
			return err
		}

		if _, err := s.DB().Exec(fmt.Sprintf(fmtDropColumn, table, column)); err != nil {
			return err
		}

		p.Log("column %s.%s dropped\n", table, column)
		return nil
	}
}

// Returns Executor fn that adds column
func (s storeProvision) addColumn(table, column, spec string) Executor {
	return func(p *Provisioner) error {
		if tt, err := s.getTableColumns(table); err != nil || s.getColumn(tt, column) != nil {
			return err
		}

		if _, err := s.DB().Exec(fmt.Sprintf(fmtAddColumn, table, column, spec)); err != nil {
			return err
		}

		p.Log("column %s.%s added\n", table, column)
		return nil
	}
}

// Returns all table's columns
func (s storeProvision) getTableColumns(name string) ([]*tableColumn, error) {
	tt := make([]*tableColumn, 0)

	if err := s.DB().Select(&tt, "DESCRIBE "+name); err != nil {
		return nil, err
	}

	return tt, nil
}

// Searches for a column by it's name in the list of columns
func (s storeProvision) getColumn(tt []*tableColumn, name string) *tableColumn {
	for _, t := range tt {
		if t.Field == name {
			return t
		}
	}

	return nil
}

// Executes one or more SQL commands
func (s storeProvision) rawSqlExec(ss ...string) error {
	for _, sql := range ss {
		if _, err := s.DB().Exec(sql); err != nil {
			return err
		}
	}

	return nil
}

// Returns Executor fn that calls rawSqlExec
func (s storeProvision) execSql(ss ...string) Executor {
	return func(p *Provisioner) error {
		return s.rawSqlExec(ss...)
	}
}
