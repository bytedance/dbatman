package proxy

import (
	"testing"
)

func TestProxy_Query(t *testing.T) {

	db := newSqlDB(testProxyDSN)
	defer db.Close()

	if _, err := db.Exec(`
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
	}

	if rs, err := db.Exec(`
		insert into go_proxy_test_proxy_conn (id, str, f, e, u, i) values(
			1, 
			"abc", 
			3.14, 
			"test1", 
			255, 
			-127)`); err != nil {
		t.Fatal("insert failed: ", err)
	} else if id, err := rs.LastInsertId(); err != nil {
		t.Fatal("insert failed: ", err)
	} else if id == 0 {
		t.Fatalf("expect insert 1 rows, got %d", id)
	}
}
