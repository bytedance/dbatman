package sql

type IStatement interface {
	IStatement()
}

func SetParseTree(yylex interface{}, stmt IStatement) {
	yylex.(*SQLLexer).ParseTree = stmt
}
