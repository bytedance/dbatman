package lexer

import (
	"testing"
)

func getLexer(str string) (lexer *MySQLLexer, lval *yySymType) {
	lval = new(yySymType)
	lexer = NewMySQLLexer(str)

	return
}

func testMatchReturn(t *testing.T, str string, match int, dbg bool) {
	setDebug(dbg)
	lexer, lval := getLexer(str)
	ret := lexer.Lex(lval)
	if ret != match {
		t.Fatalf("test failed! expect[%d] return[%d]", match, ret)
	}
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

func TestFloatNum(t *testing.T) {
	testMatchReturn(t, " 10e-10", FLOAT_NUM, false)
	testMatchReturn(t, " 10E+10", FLOAT_NUM, false)
	testMatchReturn(t, "   10E10", FLOAT_NUM, false)
}
