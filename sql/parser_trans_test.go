package sql

import (
	"testing"
)

func TestTransaction(t *testing.T) {
	st := testParse(`Start Transaction WITH CONSISTENT SNAPSHOT`, t, false)
	matchType(t, st, &StartTrans{})

	st = testParse(`BEGIN`, t, false)
	matchType(t, st, &Begin{})

	st = testParse(`COMMIT WORk NO RELEASE`, t, false)
	matchType(t, st, &Commit{})

	st = testParse(`rollback`, t, false)
	matchType(t, st, &Rollback{})
}

func TestSavePoint(t *testing.T) {
	st := testParse(`Savepoint identifier`, t, false)
	matchType(t, st, &SavePoint{})

	st = testParse(`rollback to identifier`, t, false)
	matchType(t, st, &Rollback{})

	st = testParse(`release savepoint identifier`, t, false)
	matchType(t, st, &Release{})
}

func TestLockTables(t *testing.T) {
	st := testParse(`LOCK TABLES tb1 AS alias1 read, db2.tb2 low_priority write`, t, false)
	matchType(t, st, &Lock{})
	matchSchemas(t, st, `db2`)

	st = testParse(`UNLOCK TABLES`, t, false)
	matchType(t, st, &Unlock{})
}
