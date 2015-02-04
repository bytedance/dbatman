package lexer

import (
	"fmt"
	"testing"
)

func getLexer(str string) (lexer *MySQLLexer, lval *yySymType) {
	lval = new(yySymType)
	lexer = NewMySQLLexer(str)

	return
}

func tokenName(tok int) string {
	if (tok-ABORT_SYM) < 0 || (tok-ABORT_SYM) > len(yyToknames) {
		return fmt.Sprintf("Unknown Token:%d", tok)
	}

	return yyToknames[tok-ABORT_SYM]
}

func testMatchReturn(t *testing.T, str string, match int, dbg bool) (*MySQLLexer, *yySymType) {
	setDebug(dbg)
	lexer, lval := getLexer(str)
	ret := lexer.Lex(lval)
	if ret != match {
		t.Fatalf("test failed! expect[%s] return[%s]", tokenName(match), tokenName(ret))
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
