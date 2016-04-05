package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
	conf, err := LoadConfig("./proxy.yml")
	if err != nil {
		t.Fatal(err)
	}
	cfg := conf.GetConfig()

	globalConfig := GlobalConfig{
		Port:              3306,
		ManagePort:        3307,
		MaxConnections:    10,
		LogFilename:       "./log/dbatman.log",
		LogLevel:          1,
		LogMaxSize:        1024,
		ClientTimeout:     1800,
		ServerTimeout:     1800,
		WriteTimeInterval: 10,
		ConfAutoload:      1,
		AuthIPs:           []string{"10.4.64.1", "10.4.64.2"},
	}

	masterNode := NodeConfig{
		Host:                  "10.4.4.4",
		Port:                  3307,
		Username:              "pgc",
		Password:              "pgc",
		DBName:                "pgc",
		Charset:               "utf8mb4",
		MaxConnections:        100,
		MaxConnectionPoolSize: 10,
		ConnectTimeout:        10,
		TimeReconnectInterval: 10,
		Weight:                1,
	}

	userNode := UserConfig{
		Username:       "proxy_pgc_user",
		Password:       "pgc",
		MaxConnections: 1000,
		MinConnections: 100,
		DBName:         "pgc",
		Charset:        "utf8mb4",
		ClusterName:    "pgc_cluster",
		AuthIPs:        []string{"10.1.1.1", "10.1.1.2"},
		BlackListIPs:   []string{"10.1.1.3", "10.1.1.4"},
	}

	if !reflect.DeepEqual(cfg.Global, &globalConfig) {
		fmt.Printf("%v\n", globalConfig)
		t.Fatal("global must equal")
	}

	if !reflect.DeepEqual(cfg.GetMasterNodefromClusterByName("pgc_cluster"), &masterNode) {
		fmt.Printf("%v\n", masterNode)
		t.Fatal("master must equal")
	}

	if !reflect.DeepEqual(cfg.GetUserByName("proxy_pgc_user"), &userNode) {
		fmt.Printf("%v\n", userNode)
		t.Fatal("user must equal")
	}
}
