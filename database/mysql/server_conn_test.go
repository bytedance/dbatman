package mysql

import (
	//	"github.com/bytedance/dbatman/database/sql"
	"testing"
)

func TestWriteCommandFieldList(t *testing.T) {

	runTests(t, dsn, func(dbt *DBTest) {
		dbt.mustExec("CREATE TABLE `test` (`id` int(11) NOT NULL, `value` int(11) NOT NULL) ")

		var rows Rows
		var err error
		if rows, err = dbt.db.FieldList("test", ""); err != nil {
			t.Fatal(err)
		}

		cols, err := rows.ColumnPackets()
		if err != nil {
			t.Fatal(err)
		}

		if len(cols) != 2 {
			t.Fatalf("expect 2 rows, got %d", len(cols))
		}

	})
}
