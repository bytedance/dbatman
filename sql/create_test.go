package sqlparser

import "testing"

func TestCreate1(t *testing.T) {
	sql := "CREATE TABLE test_fetch_null( col1 tinyint, col2 smallint,  col3 int, col4 bigint,  col5 float, col6 double,  col7 date, col8 time,  col9 varbinary(10),  col10 varchar(50),  col11 char(20))"
	testParse(t, sql)

	sql = "CREATE TABLE test_long_data(col1 int,       col2 long varchar, col3 long varbinary)"
	testParse(t, sql)

	sql = "CREATE TABLE my_demo_transaction( col1 int , col2 varchar(30)) ENGINE= BDB"
	testParse(t, sql)

	sql = "CREATE TABLE test_prepare_ext( c1  tinyint, c2  smallint, c3  mediumint, c4  int, c5  integer, c6  bigint, c7  float, c8  double, c9  double precision, c10 real, c11 decimal(7, 4), c12 numeric(8, 4), c13 date, c14 datetime, c15 timestamp, c16 time, c17 year, c18 bit, c19 bool, c20 char, c21 char(10), c22 varchar(30), c23 tinyblob, c24 tinytext, c25 blob, c26 text, c27 mediumblob, c28 mediumtext, c29 longblob, c30 longtext, c31 enum('one', 'two', 'three'), c32 set('monday', 'tuesday', 'wednesday'))"
	testParse(t, sql)

	sql = "CREATE TABLE test_sshort(a smallint signed,                                                   b smallint signed,                                                   c smallint unsigned,                                                   d smallint unsigned)"
	testParse(t, sql)

	sql = "CREATE TABLE test_bg1500 (s VARCHAR(25), FULLTEXT(s)) engine=MyISAM"
	testParse(t, sql)

	sql = "create table t2 select * from t1"
	testParse(t, sql)
}
