package lexer

import (
	"testing"
)

func TestIdentifier(t *testing.T) {
	testMatchReturn(t, "`test ` ", IDENT_QUOTED, false)
}

func TestMultiIdentifier(t *testing.T) {
	str := "SELECT INSERT 'string     ' UPDATE DELEte `SELECT` `Update`"
	lex, lval := getLexer(str)

	lexExpect(t, lex, lval, SELECT_SYM)
	lexExpect(t, lex, lval, INSERT)

	lexExpect(t, lex, lval, TEXT_STRING)
	lvalExpect(t, lval, "'string     '")

	lexExpect(t, lex, lval, UPDATE_SYM)
	lexExpect(t, lex, lval, DELETE_SYM)

}
