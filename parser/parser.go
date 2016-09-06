package parser

import (
	"errors"
)

func Parse(sql string) (IStatement, error) {
	//TODO MEM used 70%in total
	lexer := NewSQLLexer(sql)
	if MySQLParse(lexer) != 0 {
		return nil, errors.New(lexer.LastError)
	}

	return lexer.ParseTree, nil
}
