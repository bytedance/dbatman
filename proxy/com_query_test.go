package proxy

import (
	"github.com/bytedance/dbatman/errors"
	"testing"
)

func TestProxy_Query(t *testing.T) {

	db := newSqlDB(testProxyDSN)
	defer db.Close()

	if rs, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS go_proxy_test_proxy_conn (
          	id BIGINT(64) UNSIGNED  NOT NULL,
			str VARCHAR(256),
        		f DOUBLE,
        		e enum("test1", "test2"),
        		u tinyint unsigned,
          	i tinyint,
          	ni tinyint,
          	PRIMARY KEY (id)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8`); err != nil {
		t.Fatal("create table failed: ", err)
	} else if rows, err := rs.RowsAffected(); err != nil {
		t.Fatal("create table failed: ", err)
	} else if rows != 0 {
		t.Fatal("ddl should have no affected rows")
	}

	if rs, err := db.Exec(`
		insert into go_proxy_test_proxy_conn (id, str, f, e, u, i) values(
			1, 
			"abc", 
			3.14, 
			"test1", 
			255, 
			-127)`); err != nil {
		t.Fatal("insert failed: ", errors.ErrorStack(err))
	} else if rows, err := rs.RowsAffected(); err != nil {
		t.Fatal("insert failed: ", errors.ErrorStack(err))
	} else if rows != 1 {
		t.Fatalf("expect insert 1 rows, got %d", rows)
	}

	if rs, err := db.Exec(`
		update go_proxy_test_proxy_conn 
			set str="abcde", f=3.1415926, e="test2", u=128, i=126
			where id=1`); err != nil {
		t.Fatal("update failed: ", errors.ErrorStack(err))
	} else if rows, err := rs.RowsAffected(); err != nil {
		t.Fatal("update failed: ", errors.ErrorStack(err))
	} else if rows != 1 {
		t.Fatalf("expect update 1 rows, got %d", rows)
	}

	if rs, err := db.Exec(`
		insert into go_proxy_test_proxy_conn (id, str, f, e, u, i) values(
			2, 
			"abc", 
			3.14, 
			"test1", 
			255, 
			-127)`); err != nil {
		t.Fatal("insert failed: ", errors.ErrorStack(err))
	} else if rows, err := rs.RowsAffected(); err != nil {
		t.Fatal("insert failed: ", errors.ErrorStack(err))
	} else if rows != 1 {
		t.Fatalf("expect insert 1 rows, got %d", rows)
	}

	if rs, err := db.Exec(`delete from go_proxy_test_proxy_conn where id = 1 or id = 2`); err != nil {
		t.Fatal("delete failed: ", errors.ErrorStack(err))
	} else if rows, err := rs.RowsAffected(); err != nil {
		t.Fatal("delete failed: ", errors.ErrorStack(err))
	} else if rows != 2 {
		t.Fatalf("expect delete 2 rows, got %d", rows)
	}
}

func TestProxy_QueryFailed(t *testing.T) {

	db := newSqlDB(testProxyDSN)
	defer db.Close()

	if _, err := db.Exec(`
		update go_proxy_test_proxy_conn 
			set str="abcde", f=3.1415926, e="test2", u=128, i=255
			when id=1`); err == nil {
		t.Fatal("syntax error sql expect error, but go ok")
	}
}

func TestProxy_Use(t *testing.T) {

	db := newSqlDB(testProxyDSN)
	defer db.Close()

	if _, err := db.Exec("use dbatman_test"); err != nil {
		t.Fatalf("use dbatman_test failed: %s", err.Error())
	}

	if _, err := db.Exec("use mysql"); err == nil {
		t.Fatalf("use mysql for this user expect deny, got pass")
	}
}
