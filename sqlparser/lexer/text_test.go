package lexer

import (
	"testing"
)

func TestSingleQuoteString(t *testing.T) {
	testMatchReturn(t, "'single Quoted string'", TEXT_STRING, false)
}

func TestDoubleQuoteString(t *testing.T) {
	testMatchReturn(t, "\"double quoted string\"", TEXT_STRING, false)
}

func TestAnsiQuotesSQLModeString(t *testing.T) {
	lexer, lval := getLexer("'a' ' ' 'string'")
	lexer.sqlMode.MODE_ANSI_QUOTES = true

	if lexer.Lex(lval) != TEXT_STRING {
		t.Fatalf("parse ansi quotes string failed!")
	}

}

func TestSingleQuoteString3(t *testing.T) {
	lexer, lval := getLexer("'afasgasdgasg'")
	if lexer.Lex(lval) != TEXT_STRING {
		t.Fatalf("parse ansi quotes string failed!")
	}
}
