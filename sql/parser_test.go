package sql

import (
	"testing"
)

func testParse(sql string, t *testing.T, dbg bool) {
	setDebug(dbg)
	if _, err := Parse(sql); err != nil {
		setDebug(false)
		t.Fatalf("%v", err)
	}
	setDebug(false)
}

func TestSelect(t *testing.T) {
	testParse("SELECT * FROM table1;", t, false)
	testParse("SELECT t1.* FROM (select * from table1) as t1;", t, false)
	testParse("SELECT sb1,sb2,sb3 \n FROM (SELECT s1 AS sb1, s2 AS sb2, s3*2 AS sb3 FROM t1) AS sb \n    WHERE sb1 > 1;", t, false)

	testParse("SELECT AVG(SUM(column1)) FROM t1 GROUP BY column1;", t, false)

	testParse("SELECT REPEAT('a',1) UNION SELECT REPEAT('b',10);", t, true)
	// testParse("SELECT * FROM table1 LEFT JOIN table2 ON table1.id=table2.id LEFT JOIN table3 ON table2.id = table3.id ", t, true)
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
