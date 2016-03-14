package sql

import (
	"testing"
)

func TestDesc(t *testing.T) {
	st := testParse(` DESCRIBE db1.tb1;`, t, false)
	matchType(t, st, &DescribeTable{})
	matchSchemas(t, st, `db1`)

	st = testParse(`explain select * from db1.table1`, t, false)
	matchSchemas(t, st, `db1`)
}

func TestHelp(t *testing.T) {
	st := testParse(`help 'help me'`, t, false)
	matchType(t, st, &Help{})
}

func TestUse(t *testing.T) {
	st := testParse(`use mydb`, t, false)
	matchType(t, st, &Use{})

	if string(st.(*Use).DB) != `mydb` {
		t.Fatalf("expect [mydb] match[%s]", string(st.(*Use).DB))
	}
}
