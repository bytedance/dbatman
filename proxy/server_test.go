package proxy

import (
	"github.com/wangjild/go-mysql-proxy/client"
	"github.com/wangjild/go-mysql-proxy/config"
	"sync"
	"testing"
	"time"
)

var testServerOnce sync.Once
var testServer *Server
var testDBOnce sync.Once
var testDB *client.DB

var testConfigData = []byte(`
addr : 127.0.0.1:4000
user : root
password : root 

nodes :
- 
    name : node1 
    down_after_noalive : 300
    idle_conns : 16
    rw_split: false
    user: root
    password: 
    master : 127.0.0.1:4306
    slave : 
- 
    name : node2
    down_after_noalive : 300
    idle_conns : 16
    rw_split: false
    user: root
    password: 
    master : 127.0.0.1:4307

- 
    name : node3 
    down_after_noalive : 300
    idle_conns : 16
    rw_split: false
    user: root
    password: 
    master : 127.0.0.1:4308

schemas :
-
    db : proxy_test 
    nodes: [node1, node2, node3]
    rules:
        default: node1 
        shard:
            -   
                table: proxy_test_shard_hash
                key: id
                nodes: [node2, node3]
                type: hash

            -   
                table: proxy_test_shard_range
                key: id
                nodes: [node2, node3]
                range: -10000-
                type: range
`)

func newTestServer(t *testing.T) *Server {
	f := func() {
		cfg, err := config.ParseConfigData(testConfigData)
		if err != nil {
			t.Fatal(err.Error())
		}

		testServer, err = NewServer(cfg)
		if err != nil {
			t.Fatal(err)
		}

		go testServer.Run()

		time.Sleep(1 * time.Second)
	}

	testServerOnce.Do(f)

	return testServer
}

func newTestDB(t *testing.T) *client.DB {
	newTestServer(t)

	f := func() {
		var err error
		testDB, err = client.Open("127.0.0.1:4000", "root", "", "go_proxy")

		if err != nil {
			t.Fatal(err)
		}

		testDB.SetMaxIdleConnNum(4)
	}

	testDBOnce.Do(f)
	return testDB
}

func newTestDBConn(t *testing.T) *client.SqlConn {
	db := newTestDB(t)

	c, err := db.GetConn()

	if err != nil {
		t.Fatal(err)
	}

	return c
}

func TestServer(t *testing.T) {
	newTestServer(t)
}
