package sql

type IStatement interface {
	Statement()
}

/*******************************************
 * Table Interfaces and Structs
 ******************************************/
type ITable interface {
	Table()
}

func (*JoinTable) ITable()    {}
func (*ParenTable) ITable()   {}
func (*AliasedTable) ITable() {}

type JoinTable struct {
	Left  ITable
	Join  []byte
	Right ITable
	// TODO On    BoolExpr
}

type ParenTable struct {
	Table ISimpleTable
}

type AliasedTable struct {
	Table ISimpleTable
	As    []byte
	// TODO IndexHints
}

// SimpleTable
func (*SimpleTable) SimpleTable() {}

type ISimpleTable interface {
	SimpleTable()
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

func SetParseTree(yylex interface{}, stmt IStatement) {
	yylex.(*SQLLexer).ParseTree = stmt
}
