package sql

import (
	"testing"
)

func TestParse(t *testing.T) {
	if _, err := Parse("Select version()"); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestTokenName(t *testing.T) {
	if name := MySQLTokname(ABORT_SYM); name == "" {
		t.Fatal("get token name error")
	}
}
