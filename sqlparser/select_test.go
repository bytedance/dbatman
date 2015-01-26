package sqlparser

import "testing"

func TestSelectFuncs(t *testing.T) {
	sql := "SELECT DATABASE(), CURRENT_USER(),                               CURRENT_DATE(), CURRENT_TIME()"
	testParse(t, sql)

	stmt, _ := Parse(sql)

	if db, err := GetStmtDB(stmt); err != nil {
		t.Fatal(db, err)
	}

	sql = "SELECT * FROM test_select WHERE id= ? AND CONVERT(name USING utf8) =?"
	testParse(t, sql)

	sql = "SELECT DATE_FORMAT(ts, '%Y') AS `venu` FROM test_dateformat"
	testParse(t, sql)

	sql = "SELECT DATE_FORMAT(ts, '%Y') AS 'venu' FROM test_dateformat"
	testParse(t, sql)

	sql = "SELECT int_c, var_c, date_c as date, ts_c, char_c FROM  test_prepare_field_result as t1 WHERE int_c=?"
	testParse(t, sql)

	sql = "SELECT * FROM t2 join t1 using(a)"
	testParse(t, sql)

	sql = "SELECT * FROM t2 natural join t1"
	testParse(t, sql)

	sql = "SELECT * FROM t2 natural right join t1"
	testParse(t, sql)

}
