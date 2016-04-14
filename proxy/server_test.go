package proxy

import (
	"fmt"
	"github.com/bytedance/dbatman/config"
	"github.com/bytedance/dbatman/database/cluster"
	"github.com/bytedance/dbatman/database/mysql"
	"github.com/bytedance/dbatman/database/sql"
	"github.com/bytedance/dbatman/errors"
	"github.com/ngaut/log"

	gosql "database/sql"
	_ "github.com/go-sql-driver/mysql"

	"os"
	"sync"
	"testing"
	"time"
)

var testServerOnce sync.Once
var testServer *Server
var testServerError error

var testClusterOnce sync.Once
var testCluster *cluster.Cluster
var testClusterError error

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
    dbatman_test_cluster:
        master:
            host: 127.0.0.1
            port: 3306
            username: root
            password: 
            dbname: dbatman_test
            charset: utf8mb4
            max_connections: 100
            max_connection_pool_size: 10
            connect_timeout: 10
            time_reconnect_interval: 10
            weight: 1
        slaves:
          - host: 127.0.0.1
            port: 3306
            username: root
            password: 
            dbname: dbatman_test
            charset: utf8mb4
            max_connections: 100
            max_connection_pool_size: 10
            connect_timeout: 10
            time_reconnect_interval: 10
            weight: 1

users:
    proxy_mysql_user:
        username: proxy_mysql_user
        password: proxy_mysql_passwd
        max_connections: 1000
        min_connections: 100
        dbname: dbatman_test
        charset: utf8mb4
        cluster_name: dbatman_test_cluster
        auth_ips:
            - 127.0.0.1
        black_list_ips:
            - 10.1.1.3
            - 10.1.1.4
`)

var testDBDSN = "root:@tcp(127.0.0.1:3306)/mysql"
var testProxyDSN = "proxy_mysql_user:proxy_mysql_passwd@tcp(127.0.0.1:3307)/dbatman_test"

func newTestServer() (*Server, error) {
	f := func() {

		path, err := tmpFile(testConfigData)
		if err != nil {
			testServer, testServerError = nil, err
			return
		}

		defer os.Remove(path) // clean up tmp file

		cfg, err := config.LoadConfig(path)
		if err != nil {
			testServer, testServerError = nil, err
			return
		}

		if err := cluster.Init(cfg); err != nil {
			testServer, testServerError = nil, err
			return
		}

		log.SetLevel(log.LogLevel(cfg.GetConfig().Global.LogLevel))
		mysql.SetLogger(log.Logger())

		testServer, err = NewServer(cfg)
		if err != nil {
			testServer, testServerError = nil, err
			return
		}

		go testServer.Serve()

		time.Sleep(1 * time.Second)
	}

	testServerOnce.Do(f)

	return testServer, testServerError
}

func newTestCluster(cluster_name string) (*cluster.Cluster, error) {
	if _, err := newTestServer(); err != nil {
		testCluster, testClusterError = nil, err
	}

	f := func() {
		testCluster, testClusterError = cluster.New(cluster_name)
	}

	testClusterOnce.Do(f)
	return testCluster, testClusterError
}

func newTestDB(t *testing.T) *sql.DB {
	cls, err := newTestCluster("dbatman_test_cluster")
	if err != nil {
		t.Fatal(err)
	}

	db, err := cls.Master()

	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
	return db
}

// return a direct connection to proxy server, this is a
func newRawProxyConn(t *testing.T) *mysql.MySQLConn {
	newTestServer()

	d := mysql.MySQLDriver{}

	if conn, err := d.Open(testProxyDSN); err != nil {
		t.Fatal(err)
	} else if c, ok := conn.(*mysql.MySQLConn); !ok {
		t.Fatal("connection is not MySQLConn type")
	} else {
		return c
	}

	return nil
}

// return a direct connection to proxy server, this is a
func newSqlDB(dsn string) *gosql.DB {

	db, err := gosql.Open("mysql", dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s is unavailable", dsn)
		os.Exit(2)
	}

	if err := db.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "%s is unreacheable", dsn)
		os.Exit(2)
	}

	return db
}

func TestMain(m *testing.M) {
	// Init dbatman_test database

	errors.SetTrace(true)

	db := newSqlDB(testDBDSN)

	// Create DataBase dbatman_test
	if _, err := db.Exec("CREATE DATABASE IF NOT EXISTS `dbatman_test`"); err != nil {
		fmt.Fprintln(os.Stderr, "create database `dbatman_test` failed: ", err.Error())
		os.Exit(2)
	}

	if _, err := newTestServer(); err != nil {
		fmt.Fprintln(os.Stderr, "setup proxy server failed: ", err.Error())
		os.Exit(2)
	}

	if _, err := newTestCluster("dbatman_test_cluster"); err != nil {
		fmt.Fprintln(os.Stderr, "setup proxy -> cluster failed: ", err.Error())
		os.Exit(2)
	}

	exit := m.Run()

	// Clear Up Database

	if _, err := db.Exec("DROP DATABASE IF EXISTS `dbatman_test`"); err != nil {
		fmt.Fprintln(os.Stderr, "drop database `dbatman_test` failed: ", err.Error())
		os.Exit(2)
	}

	os.Exit(exit)
}
