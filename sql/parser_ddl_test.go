package sql

import (
	"testing"
)

func TestAlter(t *testing.T) {
	st := testParse(`alter view v1 as select * from t2;`, t, false)
	matchSchemas(t, st)
}

func TestCreate(t *testing.T) {
	st := testParse(`CREATE DATABASE IF NOT EXISTS my_db default charset utf8 COLLATE utf8_general_ci;`, t, false)

	st = testParse(`CREATE EVENT mydb.myevent
                        ON SCHEDULE AT CURRENT_TIMESTAMP + INTERVAL 1 HOUR
                        DO
                            UPDATE myschema.mytable SET mycol = mycol + 1;`, t, false)
	matchSchemas(t, st, `mydb`)

	st = testParse(`CREATE FUNCTION thisdb.hello (s CHAR(20)) RETURNS CHAR(50) DETERMINISTIC RETURN CONCAT('Hello, ',s,'!');`, t, false)
	matchSchemas(t, st, `thisdb`)

	st = testParse(
		`CREATE DEFINER = 'admin'@'localhost' PROCEDURE db1.account_count()
            SQL SECURITY INVOKER
            BEGIN
                SELECT 'Number of accounts:', COUNT(*) FROM mysql.user;
            END;`, t, false)
	matchSchemas(t, st, `db1`)

	st = testParse(`CREATE INDEX part_of_name ON customer (name(10));`, t, false)
	st = testParse(`CREATE INDEX id_index ON lookup (id) USING BTREE;`, t, false)
	st = testParse(`CREATE INDEX id_index ON t1 (id) COMMENT 'MERGE_THRESHOLD=40';`, t, false)

	st = testParse(
		`CREATE SERVER s FOREIGN DATA WRAPPER mysql
            OPTIONS (USER 'Remote', HOST '192.168.1.106', DATABASE 'test');`, t, false)
}

func TestCreateTable(t *testing.T) {
	st := testParse(`CREATE TABLE db1.t1 (col1 INT, col2 CHAR(5))
        PARTITION BY HASH(col1);`, t, false)
	matchSchemas(t, st, `db1`)

	testParse(`CREATE TABLE t1 (col1 INT, col2 CHAR(5), col3 DATETIME)
        PARTITION BY HASH ( YEAR(col3) );`, t, false)
	testParse(`CREATE /*!32302 TEMPORARY */ TABLE t (a INT);`, t, false)

	testParse(`SELECT /*! STRAIGHT_JOIN */ col1 FROM table1,table2`, t, false)
}
