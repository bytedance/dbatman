package lexer

import (
	"testing"
)

func TestParse(t *testing.T) {
	sql := " select * from tablename;"
	lexer := NewMySQLLexer(sql)
	if ret := yyParse(lexer); ret != 0 {
		t.Fatalf("yyParse return[%d]", ret)
	}
}

func TestUpdateSQL(t *testing.T) {
}

func TestSetSQL(t *testing.T) {
	testyyParse(`set global sysvar = 1`, t)
	testyyParse(`set global autocommit = 1`, t)
	testyyParse(`set session sysvar = 123`, t)
	testyyParse(`set @@sysvar = 1`, t)
	testyyParse(`set @@global.sysvar = 1`, t)
	testyyParse(`set @@global. sysvar = 1`, t)
}

func TestShowSQL(t *testing.T) {
	testyyParse(`show tables like '%tablename%'`, t)
	testyyParse(`show databases`, t)
}

func TestSelectSQL(t *testing.T) {
	testyyParse(`select version() ;`, t)
}

func testyyParse(sql string, t *testing.T) {
	lexer := NewMySQLLexer(sql)
	if ret := yyParse(lexer); ret != 0 {
		t.Fatalf("yyParse return[%d]", ret)
	}
}
