package mysql

import (
	"testing"
)

func TestWriteCommandFieldList(t *testing.T) {

	d := MySQLDriver{}

	conn, err := d.Open(dsn)

	if err != nil {
		t.Fatal(err)
	}

	conn.(*MySQLConn).WriteCommandFieldList("test", "")

	conn.(*MySQLConn).WriteCommandFieldList("test", "%%")

}
