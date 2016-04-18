package proxy

import (
	"testing"
)

func TestProxy_Show(t *testing.T) {

	db := newSqlDB(testProxyDSN)
	defer db.Close()

	if q, err := db.Query("show databases"); err != nil {
		t.Fatalf("show databases failed: %s", err.Error())
	} else {
		q.Next()
		var database string
		if err := q.Scan(&database); err != nil {
			t.Fatalf("show databases got error %s", err)
		} else if database != "dbatman_test" {
			t.Fatalf("expect %s, got %s", "dbatman", database)
		}
	}

	if _, err := db.Query("show tables"); err != nil {
		t.Fatalf("show tables failed: %s", err.Error())
	}
}
