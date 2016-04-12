package proxy

import (
	"github.com/bytedance/dbatman/database/mysql"
	"testing"
)

func TestProxy_ComPing(t *testing.T) {
	conn := newTestProxyConn(t)

	if err := conn.WriteCommandPacket(mysql.ComPing); err != nil {
		t.Fatal(err)
	}

	_, err := conn.ReadPacket()
	if err != nil {
		t.Fatal(err)
	}
}
