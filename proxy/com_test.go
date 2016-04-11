package proxy

import (
	"github.com/bytedance/dbatman/database/sql"
	"testing"
)

func TestDB_Handshake(t *testing.T) {
	cls := newTestCluster(t, "mysql_cluster")

	var db *sql.DB
	var err error

	if db, err = cls.Master(); err != nil {
		t.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		t.Fatal(err)
	}

}
