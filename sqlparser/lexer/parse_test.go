package lexer

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	sql := " select * from db.table"
	lexer := NewMySQLLexer(sql)
	fmt.Println(yyParse(lexer))
}
