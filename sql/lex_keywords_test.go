package sql

import (
	"testing"
)

func TestKeywords(t *testing.T) {
	testMatchReturn(t, `SELECT`, SELECT_SYM, false)
}

func TestFunctions(t *testing.T) {
	testMatchReturn(t, `CURTIME()`, CURTIME, false)
}

func TestCharsetName(t *testing.T) {
	testMatchReturn(t, `_utf8_general_cli`, UNDERSCORE_CHARSET, false)
}

func TestIdent(t *testing.T) {
	testMatchReturn(t, `thisisaident`, IDENT, false)
}

func TestBoolOp(t *testing.T) {
	testMatchReturn(t, `&&`, AND_AND_SYM, false)
	testMatchReturn(t, `||`, OR_OR_SYM, false)
	testMatchReturn(t, `<`, LT, false)
	testMatchReturn(t, `<=`, LE, false)
	testMatchReturn(t, `<>`, NE, false)
	testMatchReturn(t, `!=`, NE, false)
	testMatchReturn(t, `=`, EQ, false)
	testMatchReturn(t, `>`, GT_SYM, false)
	testMatchReturn(t, `>=`, GE, false)
	testMatchReturn(t, `<<`, SHIFT_LEFT, false)
	testMatchReturn(t, `>>`, SHIFT_RIGHT, false)
	testMatchReturn(t, `<=>`, EQUAL_SYM, false)

	testMatchReturn(t, `:=`, SET_VAR, false)
}

func TestChar(t *testing.T) {
	testMatchReturn(t, `& `, '&', false)
}

func TestMultiKeywords(t *testing.T) {
	lexer, lval := getLexer(`SELECT SHOW Databases SELECT `)

	lexExpect(t, lexer, lval, SELECT_SYM)
	lexExpect(t, lexer, lval, SHOW)
	lexExpect(t, lexer, lval, DATABASES)

	lexExpect(t, lexer, lval, SELECT_SYM)
}
