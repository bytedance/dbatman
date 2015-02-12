package sql

import (
	"fmt"
	"testing"
)

func PrintTree(statement IStatement) {
	if statement == nil {
		fmt.Println(`(nil)`)
	}

	switch st := statement.(type) {
	case *Union:
		fmt.Printf("left: %+v right: %+v\n", st.Right)
	case *Select:
		fmt.Printf("From: %+v Lock: %+v\n", st.From, st.LockType)
	default:
		fmt.Println("Yet Unknow Statement:", st)
	}
}

func testParse(sql string, t *testing.T, dbg bool) IStatement {
	setDebug(dbg)
	if st, err := Parse(sql); err != nil {
		setDebug(false)
		t.Fatalf("%v", err)
		return nil
	} else {
		setDebug(false)
		return st
	}
}

func TestExplain(t *testing.T) {
	testParse("EXPLAIN SELECT f1(5)", t, false)
	testParse("EXPLAIN SELECT * FROM t1 AS a1, (SELECT BENCHMARK(1000000, MD5(NOW())));", t, false)
}

func TestParse(t *testing.T) {
	setDebug(false)
	if _, err := Parse("Select version()"); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestTokenName(t *testing.T) {
	if name := MySQLTokname(ABORT_SYM); name == "" {
		t.Fatal("get token name error")
	}
}
