package sql

import (
	"errors"
)

func Parse(sql string) (IStatement, error) {
	lexer := NewSQLLexer(sql)
	if MySQLParse(lexer) != 0 {
		return nil, errors.New(lexer.LastError)
	}

	return lexer.ParseTree, nil
}
