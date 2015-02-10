package sql

type IStatement interface {
	Statement()
}

type (
	TableInfo struct {
		Quolifier []byte
		Name      []byte
	}
)

func SetParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}
