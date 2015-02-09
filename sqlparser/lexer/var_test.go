package lexer

import (
	. "github.com/wangjild/go-mysql-proxy/sqlparser/parser"
	"testing"
)

func TestHostName(t *testing.T) {
	// testMatchReturn(t, `user@hostname`, LEX_HOSTNAME, true)
}

func TestSystemVariables(t *testing.T) {
	lexer, lval := testMatchReturn(t, `@@uservar`, '@', false)
	ret := lexer.Lex(lval)
	if ret != '@' {
		t.Fatalf("expect[IDENT_QUOTED] unexpect %s", tokenName(ret))
	}

	ret = lexer.Lex(lval)
	if ret != IDENT {
		t.Fatalf("expect[IDENT] unexpect %s", tokenName(ret))
	}
}

func TestUserDefinedVariables(t *testing.T) {
	lexer, lval := testMatchReturn(t, "@`uservar`", '@', false)
	ret := lexer.Lex(lval)
	if ret != IDENT_QUOTED {
		t.Fatalf("expect[IDENT_QUOTED] unexpect %s", tokenName(ret))
	}
}
