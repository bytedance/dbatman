package sqlparser

import (
	"fmt"
	"testing"
)

var f interface{} = fmt.Printf

func testParse(tb testing.TB, sql string) {
	_, err := Parse(sql)
	if err != nil {
		tb.Fatal(err, "sql["+sql+"]")
	}

}

func TestUse(t *testing.T) {
	sql := "USE client_test_db"
	testParse(t, sql)
}

func TestCreate(t *testing.T) {
	sql := "create database if not exists test"
	testParse(t, sql)

	sql = "create function f1 () returns int return 5"
	testParse(t, sql)

	//sql = "CREATE TABLE D1.T1"
	//testParse(t, sql)

	sql = "CREATE TABLE b1.t1 (a INT, b INT, c INT)"
	testParse(t, sql)
	s, _ := Parse(sql)
	if string(s.(*CreateTable).Table.Qualifier) != "b1" {
		t.Fatal(s.(*CreateTable).Table.Qualifier, "is not \"b1\"")
	}

	sql = "create TABLE t1 (a INT, b INT, c INT, UNIQUE (A), UNIQUE(B))"
	testParse(t, sql)
	s, _ = Parse(sql)
	if s.(*CreateTable).Table.Qualifier != nil {
		t.Fatal(s.(*CreateTable).Table.Qualifier, "is not \"\"")
	}

	sql = "create TABLE s1 ( s1 char, s2 char)"
	testParse(t, sql)

	sql = "CREATE TABLE t1(id int primary key auto_increment, name varchar(20))"
	testParse(t, sql)
}

func TestInsert(t *testing.T) {
	sql := "INSERT INTO T (F1, F2, F3, F4, F5) VALUES (?, ?, ?, ?, ?), (?, ?, ?, ?, ?), (?, ?, ?, ?, ?)"
	testParse(t, sql)

	sql = "insert into test_decimal_bug value(8), (10.22), (5.61)"
	testParse(t, sql)

	sql = "insert into test_free_result values(), (), ()"
	testParse(t, sql)

	sql = "INSERT INTO test_piping VALUES(?||?)"
	testParse(t, sql)

}

func TestSet(t *testing.T) {
	sql := "set names gbk"
	testParse(t, sql)

	//sql = "set @var=(1 in (select * from t1))"
	//testParse(t, sql)

	sql = "set names default"
	testParse(t, sql)

	sql = "set global sysvar=sysval"
	testParse(t, sql)

	sql = "set autocommit =1"
	testParse(t, sql)

	sql = "set character set 'utf8'"
	testParse(t, sql)
}

func TestSimpleSelect(t *testing.T) {
	sql := "select last_insert_id() as a"
	testParse(t, sql)

	sql = "SELECT * FROM t1"
	testParse(t, sql)
	s, _ := Parse(sql)
	if db, err := GetStmtDB(s); db != "" || err != nil {
		t.Fatalf("db[%s] error[%s]", db, err.Error())
	}
}

func TestCommentedSelect(t *testing.T) {
	sql := "SELECT /*mark for picman*/ * FROM WP_ALBUM WHERE MEMBER_ID = ? AND ID IN (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	_, err := Parse(sql)
	if err != nil {
		t.Fatal(err)
	}

}

func TestFuncSelect(t *testing.T) {
	sql := "SELECT substr('''a''bc',0,3) FROM dual"
	testParse(t, sql)
}

func TestShow(t *testing.T) {
	//sql := `admin upnode("node1", "master", "127.0.0.1")`
	//testParse(t, sql)

	sql := "show databases"
	testParse(t, sql)

	sql = "show tables from abc"
	testParse(t, sql)

	sql = "show tables from abc like a"
	testParse(t, sql)

	sql = "show tables from abc where a = 1"
	testParse(t, sql)

	sql = "show proxy abc"
	testParse(t, sql)

	sql = "show table status from abc like a"
	testParse(t, sql)

	sql = "show create table t"
	testParse(t, sql)

	sql = "show variables like 'have_innodb'"
	testParse(t, sql)

	sql = "show keys from test_show"
	testParse(t, sql)

	sql = "show warnings"
	testParse(t, sql)

	sql = "show errors"
	testParse(t, sql)

	sql = "show status like 'qcache_hits'"
	testParse(t, sql)
}

func TestDrop(t *testing.T) {
	sql := "drop function if exists sp_name"
	testParse(t, sql)

	sql = "drop procedure sp_name"
	testParse(t, sql)
}

func TestExplain(t *testing.T) {
	sql := "explain select id, name FROM test_explain"
	testParse(t, sql)
}
