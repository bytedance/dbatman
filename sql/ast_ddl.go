package sql

type IDDLSchemas interface {
	GetSchemas() []string
	HasDDLSchemas()
}

type RenameTable struct {
	ToList []*TableToTable
}

func (*RenameTable) Statement() {}

func (*TruncateTable) Statement()     {}
func (*TruncateTable) HasDDLSchemas() {}
func (t *TruncateTable) GetSchemas() []string {
	return t.Table.GetSchemas()
}

type TruncateTable struct {
	Table ISimpleTable
}
