package cluster

import (
	"fmt"
	"github.com/bytedance/dbatman/config"
	"github.com/bytedance/dbatman/database/sql"
	"github.com/ngaut/log"
	"sync"
	"time"
)

//TODO: check alive/remove unreachable/ close idle connection when timeout/config update
var (
	clustersMu   sync.RWMutex
	clusterConns = make(map[string]*Cluster)
	cfgHandler   *config.Conf
)

type Cluster struct {
	masterDB   *sql.DB
	slavesDB   []*sql.DB
	slaveNum   int
	cluserName string
}

func (c *Cluster) Master() (*sql.DB, error) {
	if c.masterDB == nil {
		return nil, fmt.Errorf("MasterConn error c.masterDB==nil")
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

func (c *Cluster) Probe(idleTimeout int) error {
	if c.masterDB == nil || c.slavesDB == nil {
		return fmt.Errorf("Probe error c.masterDB==nil")
	}
	for {
		time.Sleep(time.Second * 5)
		log.Infof("Cluster %s probe", c.cluserName)
		err := c.masterDB.ProbeIdleConnection(idleTimeout)
		if err != nil {
			log.Errorf("Master node probe error msg:%s", err.Error())
		}
		for _, db := range c.slavesDB {
			err := db.ProbeIdleConnection(idleTimeout)
			if err != nil {
				log.Errorf("Slave node probe error msg:%s", err.Error())
			}
		}
	}
	return nil
}

func Init(cfg *config.Conf) error {
	if cfg == nil {
		err := fmt.Errorf("config is nil")
		return err
	}

	cfgHandler = cfg
	proxyConfig := cfg.GetConfig()
	allClusterConfigs, _ := proxyConfig.GetAllClusters()
	serverTimeout := proxyConfig.ServerTimeout()

	for clusterName, clusterCfg := range allClusterConfigs {

		master := clusterCfg.GetMasterNode()
		slaves := clusterCfg.GetSlaveNodes()
		slaveNum := len(slaves)
		oneCluster := Cluster{nil, make([]*sql.DB, slaveNum), slaveNum, clusterName}

		db, err := openDBFromNode(master)
		if err != nil {
			return err
		}
		db.SetMaxOpenConns(master.MaxConnections)
		db.SetMaxIdleConns(master.MaxConnectionPoolSize)
		oneCluster.masterDB = db

		for i, slave := range slaves {
			db, err := openDBFromNode(slave)
			if err != nil {
				return err
			}
			db.SetMaxOpenConns(slave.MaxConnections)
			db.SetMaxIdleConns(slave.MaxConnectionPoolSize)
			oneCluster.slavesDB[i] = db
		}
		clusterConns[clusterName] = &oneCluster
		go oneCluster.Probe(serverTimeout)
	}

	return nil

}

func New(clusterName string) (*Cluster, error) {
	if cluster, ok := clusterConns[clusterName]; ok {
		return cluster, nil
	} else {
		return nil, fmt.Errorf("clusterName[%s] not exists", clusterName)
	}
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

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&timeout=%dms",
		node.Username,
		node.Password,
		node.Host,
		node.Port,
		node.DBName,
		node.Charset,
		node.ConnectTimeout)
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
