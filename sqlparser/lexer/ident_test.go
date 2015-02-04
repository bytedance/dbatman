package lexer

import (
	"testing"
)

func TestIdentifier(t *testing.T) {
	testMatchReturn(t, "`test ` ", IDENT_QUOTED, false)
}
