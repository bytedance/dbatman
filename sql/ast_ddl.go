package sql

type IDDLStatement interface {
	IDDLStatement()
	IStatement()
}

type IDDLSchemas interface {
	GetSchemas() []string
	HasDDLSchemas()
}

type RenameTable struct {
	ToList []*TableToTable
}

func (*RenameTable) IStatement()    {}
func (*RenameTable) IDDLStatement() {}

func (*TruncateTable) IStatement()    {}
func (*TruncateTable) IDDLStatement() {}
func (*TruncateTable) HasDDLSchemas() {}
func (t *TruncateTable) GetSchemas() []string {
	return t.Table.GetSchemas()
}

type TruncateTable struct {
	Table ISimpleTable
}
