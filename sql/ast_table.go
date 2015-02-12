package sql

/*******************************************
 * Table Interfaces and Structs
 * doc:
 *  - table_references http://dev.mysql.com/doc/refman/5.7/en/join.html
 *  - table_factor     http://dev.mysql.com/doc/refman/5.7/en/join.html
 *  - join_table       http://dev.mysql.com/doc/refman/5.7/en/join.html
 ******************************************/
type ITable interface {
	IsTable()
}

type ITables []ITable

func (*JoinTable) IsTable()    {}
func (*ParenTable) IsTable()   {}
func (*AliasedTable) IsTable() {}

type JoinTable struct {
	Left  ITable
	Join  []byte
	Right ITable
	// TODO On    BoolExpr
}

type ParenTable struct {
	Table ITable
}

type AliasedTable struct {
	TableOrSubQuery interface{} // here may be the table_ident or subquery
	As              []byte
	// TODO IndexHints
}

// SimpleTable contains only qualifier, name and a column field
func (*SimpleTable) IsSimpleTable() {}
func (*SimpleTable) IsTable()       {}

type ISimpleTable interface {
	IsSimpleTable()
	ITable
}

type SimpleTable struct {
	Qualifier []byte
	Name      []byte
	Column    []byte
}

type Spname struct {
	Qualifier []byte
	Name      []byte
}

type SchemaInfo struct {
	Name []byte
}
