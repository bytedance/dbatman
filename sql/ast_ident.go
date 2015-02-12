package sql

type IStatement interface {
	Statement()
}

func SetParseTree(yylex interface{}, stmt IStatement) {
	yylex.(*SQLLexer).ParseTree = stmt
}
