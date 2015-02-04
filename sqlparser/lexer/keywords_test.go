package lexer

import (
	"testing"
)

func TestKeywords(t *testing.T) {
	testMatchReturn(t, `SELECT`, SELECT_SYM, true)
}
