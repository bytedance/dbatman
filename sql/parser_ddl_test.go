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
		`CREATE DEFINER = 'admin'@'localhost' PROCEDURE account_count()
            SQL SECURITY INVOKER
            BEGIN
                SELECT 'Number of accounts:', COUNT(*) FROM mysql.user;
            END;`, t, false)
	matchSchemas(t, st)
}
