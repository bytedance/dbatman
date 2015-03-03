package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
	var testConfigData = []byte(
		`
addr : 127.0.0.1:4000
user : root
password : 
log_level : error

nodes :
- 
  name : node1 
  down_after_noalive : 300
  idle_conns : 16
  rw_split: true
  user: root
  password:
  master : 127.0.0.1:3306
  slave : 127.0.0.1:4306
- 
  name : node2
  user: root
  master : 127.0.0.1:3307

- 
  name : node3 
  down_after_noalive : 300
  idle_conns : 16
  rw_split: false
  user: root
  password:
  master : 127.0.0.1:3308

schemas :
-
  db : go_proxy 
  node: node1
  auths :
  -   
    user: xm_test
    passwd: xiaomi
  -   
    user: xm_test1
    passwd: xiaomi

`)

	cfg, err := ParseConfigData(testConfigData)
	if err != nil {
		t.Fatal(err)
	}

	if len(cfg.Nodes) != 3 {
		t.Fatal(len(cfg.Nodes))
	}

	if len(cfg.Schemas) != 1 {
		t.Fatal(len(cfg.Schemas))
	}

	testNode := NodeConfig{
		Name:             "node1",
		DownAfterNoAlive: 300,
		IdleConns:        16,
		RWSplit:          true,

		User:     "root",
		Password: "",

		Master: "127.0.0.1:3306",
		Slave:  "127.0.0.1:4306",
	}

	if !reflect.DeepEqual(cfg.Nodes[0], testNode) {
		fmt.Printf("%v\n", cfg.Nodes[0])
		t.Fatal("node1 must equal")
	}

	testNode_2 := NodeConfig{
		Name:   "node2",
		User:   "root",
		Master: "127.0.0.1:3307",
	}

	if !reflect.DeepEqual(cfg.Nodes[1], testNode_2) {
		t.Fatal("node2 must equal")
	}

	testSchema := SchemaConfig{
		DB:    "go_proxy",
		Node:  "node1",
		Auths: []Auth{Auth{User: "xm_test", Passwd: "xiaomi"}, Auth{User: "xm_test1", Passwd: "xiaomi"}},
	}

	fmt.Println(cfg.Schemas[0])
	if !reflect.DeepEqual(cfg.Schemas[0], testSchema) {
		t.Fatal("schema must equal")
	}

	if cfg.LogLevel != "error" || cfg.User != "root" || cfg.Password != "" || cfg.Addr != "127.0.0.1:4000" {
		t.Fatal("Top Config not equal.")
	}
}
