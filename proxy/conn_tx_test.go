package proxy

import (
	"testing"
)

var inAutoCommit bool = true

func TestProxy_Tx(t *testing.T) {
	db := newSqlDB(testProxyDSN)
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS dbatman_test_tx (
          	id BIGINT(64) UNSIGNED  NOT NULL,
			str VARCHAR(256),
          	PRIMARY KEY (id)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8`); err != nil {
		t.Fatal("create tx table failed: ", err)
	}

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("start transaction failed: %s", err)
	}

	if rs, err := tx.Exec(`insert into dbatman_test_tx values(
			1,
			"abc")`); err != nil {
		tx.Rollback(inAutoCommit)
		t.Fatalf("insert in transaction failed: %s", err)
	} else if rn, err := rs.RowsAffected(); err != nil {
		tx.Rollback(inAutoCommit)
		t.Fatalf("insert failed: %s", err)
	} else if rn != 1 {
		tx.Rollback(inAutoCommit)
		t.Fatalf("expect 1 rows, got %d", rn)
	}
	if _, err := tx.Exec(`savepoint a1`); err != nil {
		t.Fatalf("save point faied", err)
	}
	if rs, err := tx.Query("select * from dbatman_test_tx"); err != nil {
		t.Fatalf("select in trans failed: %s", err)
	} else {
		var row int
		for rs.Next() {
			row += 1
		}

		if row != 1 {
			t.Fatalf("expect 1 rows after transaction, got %d", row)
		}
	}

	if rs, err := tx.Exec(`insert into dbatman_test_tx values(
			2,'def')`); err != nil {
		tx.Rollback(inAutoCommit)
		t.Fatalf("insert in transaction failed: %s", err)
	} else if rn, err := rs.RowsAffected(); err != nil {
		tx.Rollback(inAutoCommit)
		t.Fatalf("insert failed: %s", err)
	} else if rn != 1 {
		tx.Rollback(inAutoCommit)
		t.Fatalf("expect 1 rows, got %d", rn)
	}
	if rs, err := tx.Query("select * from dbatman_test_tx"); err != nil {
		t.Fatalf("select in trans failed: %s", err)
	} else {
		var row int
		for rs.Next() {
			row += 1
		}

		if row != 2 {
			t.Fatalf("expect 2 rows after transaction, got %d", row)
		}
	}

	if _, err := tx.Exec(`rollback to a1`); err != nil {
		t.Fatalf("rollback to faild", err)
	}

	if rs, err := tx.Query("select * from dbatman_test_tx"); err != nil {
		t.Fatalf("select in trans failed: %s", err)
	} else {
		var row int
		for rs.Next() {
			row += 1
		}

		if row != 1 {
			t.Fatalf("expect 0 rows after transaction, got %d", row)
		}
	}
	// add savepoint
	if err := tx.Rollback(inAutoCommit); err != nil {
		t.Fatalf("rollback in trans failed: %s", err)
	}
	tx, err = db.Begin()

	if err := tx.Rollback(inAutoCommit); err != nil {
		t.Fatalf("rollback in trans failed: %s", err)
	}

	if rs, err := db.Query("select * from dbatman_test_tx"); err != nil {
		t.Fatalf("select after trans failed: %s", err)
	} else {
		var row int
		for rs.Next() {
			row += 1
		}

		if row > 0 {
			t.Fatalf("expect none rows after transaction, got %d", row)
		}
	}

	tx, err = db.Begin()
	if err != nil {
		t.Fatalf("start transaction failed: %s", err)
	}

	if rs, err := tx.Exec(`insert into dbatman_test_tx values(
			1,
			"abc")`); err != nil {
		tx.Rollback(inAutoCommit)
		t.Fatalf("insert in transaction failed: %s", err)
	} else if rn, err := rs.RowsAffected(); err != nil {
		tx.Rollback(inAutoCommit)
		t.Fatalf("insert failed: %s", err)
	} else if rn != 1 {
		tx.Rollback(inAutoCommit)
		t.Fatalf("expect 1 rows, got %d", rn)
	}

	if rs, err := tx.Query("select * from dbatman_test_tx"); err != nil {
		t.Fatalf("select in trans failed: %s", err)
	} else {
		var row int
		for rs.Next() {
			row += 1
		}

		if row != 1 {
			t.Fatalf("expect 1 rows after transaction, got %d", row)
		}
	}

	if err := tx.Commit(inAutoCommit); err != nil {
		t.Fatalf("commit in trans failed: %s", err)
	}

	if rs, err := db.Query("select * from dbatman_test_tx"); err != nil {
		t.Fatalf("select after trans failed: %s", err)
	} else {
		var row int
		for rs.Next() {
			row += 1
		}

		if row != 1 {
			t.Fatalf("expect 1 rows after transaction, got %d", row)
		}
	}
}
