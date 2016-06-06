package parser

import (
	"testing"
)

func TestHostName(t *testing.T) {
	setDebug(true)
	// testMatchReturn(t, `user@hostname`, LEX_HOSTNAME, true)
}

func TestSystemVariables(t *testing.T) {
	lexer, lval := testMatchReturn(t, `@@uservar`, '@', false)
	ret := lexer.Lex(lval)
	if ret != '@' {
		t.Fatalf("expect[IDENT_QUOTED] unexpect %s", MySQLSymName(ret))
	}

	ret = lexer.Lex(lval)
	if ret != IDENT {
		t.Fatalf("expect[IDENT] unexpect %s", MySQLSymName(ret))
	}
}

func TestUserDefinedVariables(t *testing.T) {
	lexer, lval := testMatchReturn(t, "@`uservar`", '@', false)
	ret := lexer.Lex(lval)
	if ret != IDENT_QUOTED {
		t.Fatalf("expect[IDENT_QUOTED] unexpect %s", MySQLSymName(ret))
	}
}

func TestSetVarIdent(t *testing.T) {

	lexer, lval := testMatchReturn(t, "set @var=1", SET, false)

	lexExpect(t, lexer, lval, '@')

	lexExpect(t, lexer, lval, IDENT)
	lvalExpect(t, lval, "var")
}
