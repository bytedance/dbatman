package lexer

import (
	"testing"
)

func TestHostName(t *testing.T) {
	// testMatchReturn(t, `user@hostname`, LEX_HOSTNAME, true)
}

/*
func TestUserDefinedVariables(t *testing.T) {
	lexer, lval := testMatchReturn(t, `@'my-var'`, '@', true)
	ret := lexer.Lex(lval)
	if ret != IDENT_QUOTED {
		t.Fatalf("expect[IDENT_QUOTED] unexpect %s", tokenName(ret))
	}
}
*/
