package sql

import (
	"fmt"
	"reflect"
	"testing"
)

func fmtimport() {
	fmt.Println()
}

func matchType(t *testing.T, st IStatement, ref interface{}) {
	if reflect.TypeOf(st) != reflect.TypeOf(ref) {
		t.Fatalf("expect type[%v] not match[%v]", reflect.TypeOf(ref), reflect.TypeOf(st))
	}
}

func matchSchemas(t *testing.T, st IStatement, tables ...string) {
	var ts []string

	switch ast := st.(type) {
	case *Select:
		ts = ast.GetSchemas()
	case *Union:
		ts = ast.GetSchemas()
	case *Insert:
		ts = ast.GetSchemas()
	case *Delete:
		ts = ast.GetSchemas()
	case *Update:
		ts = ast.GetSchemas()
	case *Replace:
		ts = ast.GetSchemas()
	case *AlterView:
		ts = ast.GetSchemas()
	case IDDLSchemas:
		ts = ast.GetSchemas()
	case *Lock:
		ts = ast.GetSchemas()
	case *DescribeTable:
		ts = ast.GetSchemas()
	case *DescribeStmt:
		ts = ast.GetSchemas()
	case ITableMtStmt:
		ts = ast.GetSchemas()
	case *CacheIndex:
		ts = ast.GetSchemas()
	case *LoadIndex:
		ts = ast.GetSchemas()
	case *FlushTables:
		ts = ast.GetSchemas()
	case IShowSchemas:
		ts = ast.GetSchemas()
	default:
		t.Fatalf("unknow statement type: %T", ast)
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
	matchSchemas(t, st)

	st = testParse("SELECT t1.* FROM (select * from db1.table1) as t1;", t, false)
	matchSchemas(t, st, "db1")

	st = testParse("SELECT sb1,sb2,sb3 \n FROM (SELECT s1 AS sb1, s2 AS sb2, s3*2 AS sb3 FROM db1.t1) AS sb \n    WHERE sb1 > 1;", t, false)
	matchSchemas(t, st, "db1")

	st = testParse("SELECT AVG(SUM(column1)) FROM t1 GROUP BY column1;", t, false)
	matchSchemas(t, st)

	st = testParse("SELECT REPEAT('a',1) UNION SELECT REPEAT('b',10);", t, false)
	matchSchemas(t, st)

	st = testParse(`(SELECT a FROM db1.t1 WHERE a=10 AND B=1 ORDER BY a LIMIT 10)
		    UNION
		    (SELECT a FROM db2.t2 WHERE a=11 AND B=2 ORDER BY a LIMIT 10);`, t, false)
	matchSchemas(t, st, "db1", "db2")

	st = testParse(`SELECT funcs(s)
    FROM db1.table1
    LEFT OUTER JOIN db2.table2
    ON db1.table1.column_name=db2.table2.column_name;`, t, false)
	matchSchemas(t, st, "db1", "db2")

	st = testParse("SELECT * FROM db1.table1 LEFT JOIN db2.table2 ON table1.id=table2.id LEFT JOIN db3.table3 ON table2.id = table3.id for update", t, false)
	matchSchemas(t, st, "db1", "db2", "db3")

	if st.(*Select).LockType != LockType_ForUpdate {
		t.Fatalf("lock type is not For Update")
	}

	st = testParse(`select last_insert_id() as a`, t, false)
	st = testParse(`SELECT substr('''a''bc',0,3) FROM dual`, t, false)
	testParse(`SELECT /*mark for picman*/ * FROM filterd limit 1;`, t, false)

	testParse(`SELECT ?,?,? from t1;`, t, false)
}

func TestInsert(t *testing.T) {
	st := testParse(`INSERT INTO db1.tbl_temp2 (fld_id)
        SELECT tempdb.tbl_temp1.fld_order_id
        FROM tempdb.tbl_temp1 WHERE tbl_temp1.fld_order_id > 100;`, t, false)
	matchSchemas(t, st, "db1", "tempdb")
}

func TestUpdate(t *testing.T) {
	st := testParse(`UPDATE t1 SET col1 = col1 + 1, col2 = col1;`, t, false)
	matchSchemas(t, st)

	st = testParse("UPDATE `Table A`,`Table B` SET `Table A`.`text`=concat_ws('',`Table A`.`text`,`Table B`.`B-num`,\" from \",`Table B`.`date`,'/') WHERE `Table A`.`A-num` = `Table B`.`A-num`", t, false)
	matchSchemas(t, st)

	st = testParse(`UPDATE db1.items,db2.month SET items.price=month.price
    WHERE items.id=month.id;`, t, false)
	matchSchemas(t, st, "db1", "db2")
}

func TestDelete(t *testing.T) {
	st := testParse(`DELETE FROM db.somelog WHERE user = 'jcole'
    ORDER BY timestamp_column LIMIT 1;`, t, false)
	matchSchemas(t, st, "db")

	st = testParse(`DELETE FROM db1.t1, db2.t2 USING t1 INNER JOIN t2 INNER JOIN db3.t3
    WHERE t1.id=t2.id AND t2.id=t3.id;`, t, false)
	matchSchemas(t, st, "db1", "db2", "db3")

	st = testParse(`DELETE FROM a1, a2 USING db1.t1 AS a1 INNER JOIN t2 AS a2
    WHERE a1.id=a2.id;`, t, false)
	matchSchemas(t, st, "db1")
}

func TestReplace(t *testing.T) {
	st := testParse(`REPLACE INTO test2 VALUES (1, 'Old', '2014-08-20 18:47:00');`, t, false)
	matchSchemas(t, st)

	st = testParse(`REPLACE INTO dbname2.test2 VALUES (1, 'Old', '2014-08-20 18:47:00');`, t, false)
	matchSchemas(t, st, "dbname2")
}
