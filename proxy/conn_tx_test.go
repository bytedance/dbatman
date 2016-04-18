package proxy

import (
	"testing"
)

func TestProxy_Tx(t *testing.T) {
	db := newSqlDB(testProxyDSN)
	defer db.Close()

	if _, err := db.Begin(); err != nil {
		t.Fatalf("start transaction failed: %s", err)
	}
}
