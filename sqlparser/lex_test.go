package sql

import (
	"testing"
)

func getLexer(str string) (lexer *SQLLexer, lval *MySQLSymType) {
	lval = new(MySQLSymType)
	lexer = NewSQLLexer(str)

	return
}

func testMatchReturn(t *testing.T, str string, match int, dbg bool) (*SQLLexer, *MySQLSymType) {
	setDebug(dbg)
	lexer, lval := getLexer(str)
	ret := lexer.Lex(lval)
	if ret != match {
		t.Fatalf("test failed! expect[%s] return[%s]", TokenName(match), TokenName(ret))
	}

	return lexer, lval
}

func TestNULLEscape(t *testing.T) {
	lexer, lval := getLexer("\\N")
	if lexer.Lex(lval) != NULL_SYM {
		t.Fatal("test failed")
	}
}

func TestSingleComment(t *testing.T) {
	lexer, lval := getLexer(" -- Single Line Comment. \r\n")

	if lexer.Lex(lval) != END_OF_INPUT {
		t.Fatal("test failed")
	}
}

func TestSingleComment2(t *testing.T) {
}
