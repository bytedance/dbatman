package cluster

import (
	"fmt"
	"github.com/bytedance/dbatman/Godeps/_workspace/src/github.com/ngaut/log"
	"github.com/bytedance/dbatman/config"
	"github.com/bytedance/dbatman/database/sql"
	"sync"
)

var (
	clustersMu   sync.RWMutex
	clusterConns = make(map[string]*Cluster)
	cfgHandler   *config.Conf
)

type Cluster struct {
	masterDB *sql.DB
	slavesDB []*sql.DB
	slaveNum int
}

func Init(cfg *config.Conf) error {
	if cfg == nil {
		err := fmt.Errorf("config is nil")
		return err
	}

	cfgHandler = cfg
	proxyConfig := cfg.GetConfig()
	allClusterConfigs := proxyConfig.GetAllClusters()

	for clusterName, clusterCfg := range allClusterConfigs {

		master := clusterCfg.GetMasterNode()
		slaves := clusterCfg.GetSlaveNodes()
		slaveNum := len(slaves)
		oneCluster := Cluster{nil, make([]*sql.DB, slaveNum), slaveNum}

		db, err := openDBFromNode(master)
		if err != nil {
			return err
		}
		oneCluster.masterDB = db

		for i, slave := range slaves {
			db, err := openDBFromNode(slave)
			if err != nil {
				return err
			}
			oneCluster.slavesDB[i] = db
		}
		clusterConns[clusterName] = &oneCluster
	}

	return nil

}

func openDBFromNode(node *config.NodeConfig) (*sql.DB, error) {
	if node == nil {
		return nil, fmt.Errorf("openDBFromNode error node==nil")
	}

	dsn := getDsnFromNode(node)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Errorf("openDBFromNod sql.open error dsn:%s msg:%s\n", dsn, err.Error())
		return nil, err
	}
	return db, nil
}

func getDsnFromNode(node *config.NodeConfig) string {
	if node == nil {
		log.Errorf("getDsnFromNode error node==nil")
		return ""
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&timeout=%dms",
		node.Username,
		node.Password,
		node.Host,
		node.Port,
		node.DBName,
		node.ConnectTimeout)
	//	Node.Charset)
}

func New(clusterName string) (*Cluster, error) {
	if cluster, ok := clusterConns[clusterName]; ok {
		return cluster, nil
	} else {
		return nil, fmt.Errorf("clusterName[%s] not exists", clusterName)
	}
}

func (c *Cluster) Master() (*sql.DB, error) {
	if c.masterDB == nil {
		return nil, fmt.Errorf("MasterConn error c.masterDb==nil")
	}
	db := c.masterDB
	stats := db.Stats()
	if stats.OpenConnections == 0 {
		// set connection using Ping
		err := makeConnection(db)
		if err != nil {
			//TODO: retry mechanism
			return nil, err

		}
	}

	return db, nil
}

func (c *Cluster) Slave() (*sql.DB, error) {
	if c.slavesDB == nil {
		return nil, fmt.Errorf("SlaveConn error c.slavesDB==nil")
	}

	var (
		freeConnections      int
		openConnections      int
		serviceSlavePosition int
	)

	//load balance mechanism
	for i, db := range c.slavesDB {
		if db != nil {
			stats := db.Stats()
			if stats.FreeConnections > freeConnections ||
				(stats.FreeConnections == freeConnections && stats.OpenConnections < openConnections) {
				freeConnections, openConnections = stats.FreeConnections, stats.OpenConnections
				serviceSlavePosition = i
			}
		}
	}

	db := c.slavesDB[serviceSlavePosition]
	if openConnections == 0 {
		err := makeConnection(db)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

func (c *Cluster) DB(isread bool) (*sql.DB, error) {
	if isread {
		return c.Master()
	}

	return c.Slave()
}

func makeConnection(db *sql.DB) error {
	if db == nil {
		return fmt.Errorf("Can not make connection to database because of db is nil")
	}

	err := db.Ping()
	//TODO: retry
	if err != nil {
		log.Errorf("makeConnection error msg:%s", err.Error())
	}
	return err
}
