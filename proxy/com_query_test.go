package proxy

import (
	"testing"
)

func TestProxy_Query(t *testing.T) {

	conn := newRawProxyConn(t)

	if _, err := conn.Exec(`
		CREATE TABLE IF NOT EXISTS go_proxy_test_proxy_conn (
          	id BIGINT(64) UNSIGNED  NOT NULL,
			str VARCHAR(256),
        		f DOUBLE,
        		e enum("test1", "test2"),
        		u tinyint unsigned,
          	i tinyint,
          	ni tinyint,
          	PRIMARY KEY (id)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8`, nil); err != nil {
		t.Fatal("create table failed: ", err)
	}
}
