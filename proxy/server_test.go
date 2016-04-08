package proxy

import (
	"fmt"
	"github.com/bytedance/dbatman/config"
	"github.com/bytedance/dbatman/database/cluster"
	"github.com/bytedance/dbatman/database/mysql"
	"github.com/bytedance/dbatman/database/sql"
	"github.com/ngaut/log"
	"os"
	"sync"
	"testing"
	"time"
)

var testServerOnce sync.Once
var testServer *Server
var testClusterOnce sync.Once
var testCluster *cluster.Cluster
var proxyConfig *config.ProxyConfig

var testConfigData = []byte(`
global:
  port: 3307
  manage_port: 3308
  max_connections: 10
  log_filename: ./log/dbatman.log
  log_level: 16
  log_maxsize: 1024
  log_query_min_time: 0
  client_timeout: 1800
  server_timeout: 1800
  write_time_interval: 10
  conf_autoload: 1
  auth_ips:

clusters:
    mysql_cluster:
        master:
            host: 127.0.0.1
            port: 3306
            username: root
            password: 
            dbname: mysql
            charset: utf8mb4
            max_connections: 100
            max_connection_pool_size: 10
            connect_timeout: 10
            time_reconnect_interval: 10
            weight: 1
        slaves:

users:
    proxy_mysql_user:
        username: proxy_mysql_user
        password: proxy_mysql_passwd
        max_connections: 1000
        min_connections: 100
        dbname: mysql
        charset: utf8mb4
        cluster_name: mysql_cluster
        auth_ips:
            - 127.0.0.1
        black_list_ips:
            - 10.1.1.3
            - 10.1.1.4
`)

func newTestServer(t *testing.T) *Server {
	f := func() {

		path, err := tmpFile(testConfigData)
		if err != nil {
			t.Fatal(err)
		}

		defer os.Remove(path) // clean up tmp file

		cfg, err := config.LoadConfig(path)
		if err != nil {
			t.Fatal(err)
		}

		if err := cluster.Init(cfg); err != nil {
			t.Fatal(err)
		}

		log.SetLevel(log.LogLevel(cfg.GetConfig().Global.LogLevel))
		mysql.SetLogger(log.Logger())

		testServer, err = NewServer(cfg)
		if err != nil {
			t.Fatal(err)
		}

		go testServer.Serve()

		time.Sleep(1 * time.Second)
	}

	testServerOnce.Do(f)

	return testServer
}

func newTestCluster(t *testing.T, cluster_name string) *cluster.Cluster {
	newTestServer(t)

	f := func() {
		var err error
		testCluster, err = cluster.New(cluster_name)

		if err != nil {
			t.Fatal(err)
		}
	}

	testClusterOnce.Do(f)
	return testCluster
}

func newTestDB(t *testing.T) *sql.DB {
	cls := newTestCluster(t, "mysql_cluster")

	db, err := cls.Master()

	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
	return db
}

func TestServer(t *testing.T) {
	newTestServer(t)

	// Open Proxy
	_, err := sql.Open("mysql", fmt.Sprintf("proxy_mysql_user:proxy_mysql_passwd@tcp(127.0.0.1:3307)/mysql"))
	if err != nil {
		t.Fatal(err)
	}

	// TODO
	/*
		if err := proxy.Ping(); err != nil {
			t.Fatal(err)
		}
	*/

}
