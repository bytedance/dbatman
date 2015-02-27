package sql

type IDDLSchemas interface {
	GetSchemas() []string
	HasDDLSchemas()
}

type RenameTable struct {
	ToList []*TableToTable
}

func (*RenameTable) IStatement() {}

func (*TruncateTable) IStatement()    {}
func (*TruncateTable) HasDDLSchemas() {}
func (t *TruncateTable) GetSchemas() []string {
	return t.Table.GetSchemas()
}

type TruncateTable struct {
	Table ISimpleTable
}
