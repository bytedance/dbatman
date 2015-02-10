package ast

type IStatement interface {
	Statement()
}

type (
	TableInfo struct {
		Quolifier []byte
		Name      []byte
	}
)
