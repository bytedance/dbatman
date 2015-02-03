package lexer

import (
	"testing"
)

func testTextParse(t *testing.T, str string, mode SQLMode) {
	lexer, lval := getLexer(str)
	lexer.sqlMode = mode
	if r := lexer.Lex(lval); r != TEXT_STRING {
		t.Fatalf("parse text failed. return[%s]", yyToknames[r-ABORT_SYM])
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
		t.Fatalf("parse text failed. return[%s]", yyToknames[r-ABORT_SYM])
	}

	lexer, lval = getLexer(`"\`)
	if r := lexer.Lex(lval); r != ABORT_SYM {
		t.Fatalf("parse text failed. return[%s]", yyToknames[r-ABORT_SYM])
	}
}
