package sql

import (
	"fmt"
)

func (*Help) IStatement()          {}
func (*DescribeTable) IStatement() {}
func (*DescribeStmt) IStatement()  {}
func (*Use) IStatement()           {}

type Help struct{}

func (d *DescribeTable) GetSchemas() []string {
	return d.Table.GetSchemas()
}

type DescribeTable struct {
	Table ISimpleTable
}

func (d *DescribeStmt) GetSchemas() []string {
	switch st := d.Stmt.(type) {
	case *Select:
		return st.GetSchemas()
	case *Insert:
		return st.GetSchemas()
	case *Update:
		return st.GetSchemas()
	case *Replace:
		return st.GetSchemas()
	case *Delete:
		return st.GetSchemas()
	default:
		panic(fmt.Sprintf("statement type %T is not explainable", st))
	}
}

type DescribeStmt struct {
	Stmt IStatement
}

type Use struct {
	DB []byte
}
