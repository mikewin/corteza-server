package schema

import "fmt"

type (
	tableGenerator    func(*Table) (string, error)
	tableManipulator  func(*Table)
	columnManipulator func(*Column)

	Table struct {
		Name    string
		Columns []*Column
		Indexes []*Index
		Primary *Index
	}

	Index struct {
		Name      string
		Columns   []string
		Unique    bool
		Condition string
	}

	columnType uint
	ColumnType struct {
		Type   columnType
		Length int
	}

	Column struct {
		Name         string
		Type         ColumnType
		IsNull       bool
		DefaultValue string
		Comment      string
	}
)

const (
	// Subset of SQL types
	//
	// Just the ones we use in Corteza
	// Each type should have it's own formatter inside

	ColumnTypeIdentifier columnType = iota
	ColumnTypeText
	ColumnTypeTimestamp
	ColumnTypeInteger
	ColumnTypeJson
	ColumnTypeBoolean
)

func TableDef(name string, mm ...tableManipulator) *Table {
	var t = &Table{Name: name}
	return t.Apply(mm...)
}

func (t *Table) Apply(mm ...tableManipulator) *Table {
	for _, m := range mm {
		m(t)
	}

	return t
}

// Adds ID column and sets primary key
func AddID() tableManipulator {
	return func(t *Table) {
		t.Columns = append(t.Columns, &Column{
			Name:    "id",
			Type:    ColumnType{Type: ColumnTypeIdentifier},
			Comment: fmt.Sprintf("Unique ID for %s", t.Name),
		})
		t.Primary = &Index{Columns: []string{"id"}}
	}
}

// Adds created_at/updated_at/deleted-at columns
func CUDTimestamps(t *Table) {
	t.Apply(
		ColumnDef("created_at", ColumnTypeTimestamp),
		ColumnDef("updated_at", ColumnTypeTimestamp, Null),
		ColumnDef("deleted_at", ColumnTypeTimestamp, Null),
	)
}

// Adds created_at/updated_at/deleted-by columns
func CUDUsers(t *Table) {
	t.Apply(
		ColumnDef("created_by", ColumnTypeIdentifier),
		ColumnDef("updated_by", ColumnTypeIdentifier, DefaultValue("0")),
		ColumnDef("deleted_by", ColumnTypeIdentifier, DefaultValue("0")),
	)
}

func ColumnDef(name string, cType columnType, mm ...columnManipulator) tableManipulator {
	return func(t *Table) {
		c := &Column{Name: name, Type: ColumnType{Type: cType}}

		for _, m := range mm {
			m(c)
		}

		t.Columns = append(t.Columns, c)
	}
}

func Null(c *Column) {
	c.IsNull = true
}

func ColumnTypeLength(l int) columnManipulator {
	return func(c *Column) {
		c.Type.Length = l
	}
}

func DefaultValue(v string) columnManipulator {
	return func(c *Column) {
		c.DefaultValue = v
	}
}

func SetPrimaryKey(columns ...string) tableManipulator {
	return func(t *Table) {
		t.Primary = &Index{Columns: columns}
	}
}
