package sql

import (
	"testing"
)

func testTextParse(t *testing.T, str string, mode SQLMode) {
	lexer, lval := getLexer(str)
	lexer.sqlMode = mode
	if r := lexer.Lex(lval); r != TEXT_STRING {
		t.Fatalf("parse text failed. return[%s]", TokenName(r))
	}

	if string(lval.bytes) != str {
		t.Fatalf("orgin[%s] not match parsed[%s]", str, string(lval.bytes))
	}
}

func TestSingleQuoteString(t *testing.T) {
	testMatchReturn(t, `'single Quoted string'`, TEXT_STRING, false)
}

func TestDoubleQuoteString(t *testing.T) {
	testMatchReturn(t, `"double quoted string"`, TEXT_STRING, false)
}

func TestAnsiQuotesSQLModeString(t *testing.T) {
	str := `'a' ' ' 'string'`
	lexer, lval := getLexer(str)
	lexer.sqlMode.MODE_ANSI_QUOTES = true

	if lexer.Lex(lval) != TEXT_STRING {
		t.Fatalf("parse ansi quotes string failed!")
	}

}

func TestSingleQuoteString3(t *testing.T) {
	testTextParse(t, `'afasgasdgasg'`, SQLMode{})
	testTextParse(t, `'''afasgasdgasg'`, SQLMode{})
	testTextParse(t, `''`, SQLMode{})
	testTextParse(t, `""`, SQLMode{})

	testTextParse(t, `'""hello""'`, SQLMode{})
	testTextParse(t, `'hel''lo'`, SQLMode{})
	testTextParse(t, `'\'hello'`, SQLMode{})

	testTextParse(t, `'\''`, SQLMode{})
	testTextParse(t, `'\'`, SQLMode{MODE_NO_BACKSLASH_ESCAPES: true})
}

func TestStringException(t *testing.T) {
	str := `'\'`
	lexer, lval := getLexer(str)
	if r := lexer.Lex(lval); r != ABORT_SYM {
		t.Fatalf("parse text failed. return[%s]", MySQLToknames[r-ABORT_SYM])
	}

	lexer, lval = getLexer(`"\`)
	if r := lexer.Lex(lval); r != ABORT_SYM {
		t.Fatalf("parse text failed. return[%s]", MySQLToknames[r-ABORT_SYM])
	}
}

func TestNChar(t *testing.T) {
	testMatchReturn(t, `n'some text'`, NCHAR_STRING, false)
	testMatchReturn(t, `N'some text'`, NCHAR_STRING, false)

	testMatchReturn(t, `N'`, ABORT_SYM, false)
}

func lexExpect(t *testing.T, lexer *SQLLexer, lval *MySQLSymType, expect int) {
	if ret := lexer.Lex(lval); ret != expect {
		t.Fatalf("expect[%s] return[%s]", TokenName(expect), TokenName(ret))
	}
}

func lvalExpect(t *testing.T, lval *MySQLSymType, expect string) {
	if string(lval.bytes) != expect {
		t.Fatalf("expect[%s] return[%s]", expect, string(lval.bytes))
	}
}

func TestMultiString(t *testing.T) {
	str := `"string1" 'string2'    'string3' n'string 4'    `
	lex, lval := getLexer(str)

	lexExpect(t, lex, lval, TEXT_STRING)
	lvalExpect(t, lval, `"string1"`)

	lexExpect(t, lex, lval, TEXT_STRING)
	lvalExpect(t, lval, `'string2'`)

	lexExpect(t, lex, lval, TEXT_STRING)
	lvalExpect(t, lval, `'string3'`)

	lexExpect(t, lex, lval, NCHAR_STRING)
	lvalExpect(t, lval, `n'string 4'`)

	lexExpect(t, lex, lval, END_OF_INPUT)
}
