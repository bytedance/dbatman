package sql

import (
	"testing"
)

func testSelectSchemas(t *testing.T, st IStatement, tables ...string) {
	var ts []string
	if is, ok := st.(ISelect); !ok {
		t.Fatal("not select statement")
	} else {
		ts = is.GetSchemas()
	}

	if len(tables) == 0 && len(ts) == 0 {
		return
	} else if len(tables) != len(ts) {
		t.Fatalf("expect table number[%d] not match return[%d]", len(tables), len(ts))
	}

	for k, v := range ts {
		if v != tables[k] {
			t.Fatalf("expect table[%s] not match return[%s]", tables[k], v)
		}
	}

}

func TestSelect(t *testing.T) {
	st := testParse("SELECT * FROM table1;", t, false)
	testSelectSchemas(t, st)

	st = testParse("SELECT t1.* FROM (select * from db1.table1) as t1;", t, false)
	testSelectSchemas(t, st, "db1")

	st = testParse("SELECT sb1,sb2,sb3 \n FROM (SELECT s1 AS sb1, s2 AS sb2, s3*2 AS sb3 FROM db1.t1) AS sb \n    WHERE sb1 > 1;", t, false)
	testSelectSchemas(t, st, "db1")

	st = testParse("SELECT AVG(SUM(column1)) FROM t1 GROUP BY column1;", t, false)
	testSelectSchemas(t, st)

	st = testParse("SELECT REPEAT('a',1) UNION SELECT REPEAT('b',10);", t, false)
	testSelectSchemas(t, st)

	st = testParse(`(SELECT a FROM db1.t1 WHERE a=10 AND B=1 ORDER BY a LIMIT 10)
		    UNION
		    (SELECT a FROM db2.t2 WHERE a=11 AND B=2 ORDER BY a LIMIT 10);`, t, false)
	testSelectSchemas(t, st, "db1", "db2")

	st = testParse(`SELECT column_name(s)
    FROM table1
    LEFT OUTER JOIN table2
    ON table1.column_name=table2.column_name;`, t, false)

	st = testParse("SELECT * FROM table1 LEFT JOIN table2 ON table1.id=table2.id LEFT JOIN table3 ON table2.id = table3.id ", t, false)
}
