package proxy

import (
	"github.com/bytedance/dbatman/database/mysql"
	"testing"
)

func TestProxy_ComPing(t *testing.T) {
	conn := newRawProxyConn(t)

	if err := conn.WriteCommandPacket(mysql.ComPing); err != nil {
		t.Fatal(err)
	}

	// Test ComQuit
	defer conn.Close()

	_, err := conn.ReadPacket()
	if err != nil {
		t.Fatal(err)
	}
}

func TestProxy_ComInitDB(t *testing.T) {
	conn := newRawProxyConn(t)

	if err := conn.WriteCommandPacketStr(mysql.ComInitDB, "dbatman_test"); err != nil {
		t.Fatal(err)
	}

	// should receive an ok result
	if err := conn.ReadResultOK(); err != nil {
		t.Fatal(err)
	}

	if err := conn.WriteCommandPacketStr(mysql.ComInitDB, "db_no_exist"); err != nil {
		t.Fatal(err)
	}

	if err := conn.ReadResultOK(); err == nil {
		t.Fatal("expect an error result packet")
	} else if e, ok := err.(*mysql.MySQLError); !ok {
		t.Fatal(err)
	} else if e.Number != 1049 {
		t.Fatal("expect an Unknow DB error")
	}
}
