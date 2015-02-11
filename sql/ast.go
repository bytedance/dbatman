package sql

type IStatement interface {
	Statement()
}

type (
	TableInfo struct {
		Qualifier []byte
		Name      []byte
	}

	Spname struct {
		Qualifier []byte
		Name      []byte
	}

	SchemaInfo struct {
		Name []byte
	}
)

func SetParseTree(yylex interface{}, stmt IStatement) {
	yylex.(*SQLLexer).ParseTree = stmt
}
