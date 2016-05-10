package mysql

import (
	"testing"
)

func TestWriteCommandFieldList(t *testing.T) {

	runTests(t, dsn, func(dbt *DBTest) {
		dbt.db.FieldList("test", "")
	})
}
