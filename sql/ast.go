package sql

type IStatement interface {
	Statement()
}

type (
	TableInfo struct {
		Qualifier []byte
		Name      []byte
	}
)

func SetParseTree(yylex interface{}, stmt IStatement) {
	yylex.(*SQLLexer).ParseTree = stmt
}
