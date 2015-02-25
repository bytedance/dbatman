package sql

type IDDLSchemas interface {
	GetSchemas() []string
	HasDDLSchemas()
}

type RenameTable struct{}

func (*RenameTable) Statement() {}

type TruncateTable struct{}

func (*TruncateTable) Statement() {}
