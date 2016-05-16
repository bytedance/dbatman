package cluster

import (
	"fmt"
	"github.com/bytedance/dbatman/config"
	"github.com/bytedance/dbatman/database/sql"
	"github.com/ngaut/log"
	"sync"
	"time"
)

var (
	clustersMu            sync.RWMutex
	clusterConns          = make(map[string]*Cluster)
	cfgHandler            *config.Conf
	currentClusterVersion = 1
	NotifyChan            = make(chan bool)
)

type Cluster struct {
	masterNode *sql.DB
	slaveNodes map[string]*sql.DB
	cluserName string
	DBName     string
	version    int
}

func (c *Cluster) Master() (*sql.DB, error) {
	if c.masterNode == nil {
		return nil, fmt.Errorf("MasterConn error c.masterNode==nil")
	}
	db := c.masterNode
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
	if len(c.slaveNodes) == 0 {
		log.Warnf("Slave no exsits, use master in instead")
		return c.Master()
	}

	var (
		freeConnections int = -1
		openConnections int = -1
		serviceSlaveKey string
	)

	//load balance mechanism
	for key, db := range c.slaveNodes {
		if db != nil {
			stats := db.Stats()
			if stats.FreeConnections > freeConnections ||
				(stats.FreeConnections == freeConnections && stats.OpenConnections < openConnections) {
				freeConnections, openConnections = stats.FreeConnections, stats.OpenConnections
				serviceSlaveKey = key
			}
		}
	}
	db := c.slaveNodes[serviceSlaveKey]
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

func (c *Cluster) String() string {
	masterNodeDsn := c.masterNode.Dsn()
	slaveNodeDsns := make([]string, 0)
	for _, db := range c.slaveNodes {
		slaveNodeDsns = append(slaveNodeDsns, db.Dsn())
	}
	s := fmt.Sprintf("[clusterName:%s, version:%d, master:%s, slaves:%s]", c.cluserName, c.version, masterNodeDsn, slaveNodeDsns)
	return s
}

func Init(cfg *config.Conf) error {
	if cfg == nil {
		err := fmt.Errorf("config is nil")
		return err
	}

	cfgHandler = cfg
	allClusterConfigs, _ := cfg.GetConfig().GetAllClusters()

	for clusterName, clusterCfg := range allClusterConfigs {
		cluster, err := makeCluster(clusterName, clusterCfg)
		if err != nil {
			log.Errorf("Make cluster error clusterName:%s msg:%s", clusterName, err.Error())
			return err
		}
		clusterConns[clusterName] = cluster
	}

	go monitor()
	return nil
}

func monitor() {
	for {
		timeout := time.After(time.Second * 10)
		select {
		case <-NotifyChan:
			reload()
		case <-timeout:
			probe()
		}
	}
}

func probe() error {
	log.Info("Cluster probing")
	idleTimeout := cfgHandler.GetConfig().ServerTimeout()
	for _, c := range clusterConns {
		log.Infof("Cluster[%s] probe", c.cluserName)
		err := c.masterNode.ProbeIdleConnection(idleTimeout)
		if err != nil {
			log.Errorf("Master node probe error msg:%s", err.Error())
		}
		for _, db := range c.slaveNodes {
			err := db.ProbeIdleConnection(idleTimeout)
			if err != nil {
				log.Errorf("Slave node probe error msg:%s", err.Error())
			}
		}
	}
	return nil
}

func reload() error {
	log.Info("Cluster reloading because of config update")
	if cfgHandler == nil {
		err := fmt.Errorf("cfgHandler is nil")
		return err
	}
	proxyConfig := cfgHandler.GetConfig()
	allClusterConfigs, _ := proxyConfig.GetAllClusters()

	clustersMu.Lock()
	defer clustersMu.Unlock()
	currentClusterVersion += 1
	dbsWaitToBeClosed := []*sql.DB{}
	for clusterName, clusterCfg := range allClusterConfigs {
		cluster, ok := clusterConns[clusterName]
		if !ok { // new cluster
			addedCluster, err := makeCluster(clusterName, clusterCfg)
			if err == nil {
				clusterConns[clusterName] = addedCluster
				log.Infof("Cluster realod make new cluster[%s]", clusterName)
			} else {
				log.Errorf("Cluster reload make new cluster[%s] error msg:%s", clusterName, err.Error())
			}
		} else { // update exist cluster
			newMasterNodeCfg := clusterCfg.GetMasterNode()
			newMasterNodeDsn := getDsnFromNodeCfg(newMasterNodeCfg)
			oldMasterNodeDsn := cluster.masterNode.Dsn()
			if newMasterNodeDsn != oldMasterNodeDsn {
				newMasterNode, err := openDBFromNodeCfg(newMasterNodeCfg)
				if err != nil {
					log.Errorf("Cluster reload modify cluster[%s] master node error dsn:%s msg:%s", clusterName, newMasterNodeDsn, err.Error())
				} else {
					dbsWaitToBeClosed = append(dbsWaitToBeClosed, cluster.masterNode)
					cluster.masterNode = newMasterNode
				}
				log.Infof("Cluster reload cluster[%s] master node change from %s to %s", clusterName, oldMasterNodeDsn, newMasterNodeDsn)
			} else {
				setDBProperty(cluster.masterNode, newMasterNodeCfg)
			}

			// slave nodes
			allSlaveNodeDsns := make(map[string]bool)
			for _, newSlaveNodeCfg := range clusterCfg.GetSlaveNodes() {
				newSlaveNodeDsn := getDsnFromNodeCfg(newSlaveNodeCfg)
				node, ok := cluster.slaveNodes[newSlaveNodeDsn]
				if ok {
					setDBProperty(node, newSlaveNodeCfg)
				} else {
					newSlaveNode, err := openDBFromNodeCfg(newSlaveNodeCfg)
					if err == nil {
						cluster.slaveNodes[newSlaveNodeDsn] = newSlaveNode
						log.Infof("Cluster reload cluster[%s] add slave node %s", clusterName, newSlaveNodeDsn)
					}
				}
				allSlaveNodeDsns[newSlaveNodeDsn] = true
			}
			// remove old slave
			for dsn, slaveNode := range cluster.slaveNodes {
				_, ok := allSlaveNodeDsns[dsn]
				if !ok {
					dbsWaitToBeClosed = append(dbsWaitToBeClosed, slaveNode)
					delete(cluster.slaveNodes, dsn)
					log.Infof("Cluster reload cluster[%s] remove slave node %s", clusterName, dsn)
				}
			}

			cluster.version = currentClusterVersion
		}

	}

	// remove old cluster
	for clusterName, cluster := range clusterConns {
		if cluster.version < currentClusterVersion {
			log.Infof("Cluster reload remove cluster[%s]", clusterName)
			//append waitToCLose
			dbsWaitToBeClosed = append(dbsWaitToBeClosed, cluster.masterNode)
			for _, slaveNode := range cluster.slaveNodes {
				dbsWaitToBeClosed = append(dbsWaitToBeClosed, slaveNode)
			}
			delete(clusterConns, clusterName)
		} else {
			log.Infof("Cluster reload now cluster:%s", cluster)
		}
	}

	go closeClusterDBConns(dbsWaitToBeClosed)
	return nil
}

func closeClusterDBConns(dbsWaitToBeClosed []*sql.DB) {
	for _, db := range dbsWaitToBeClosed {
		err := db.Close()
		if err != nil {
			log.Errorf("Close cluster DB conns error msg:%s", err.Error())
		}
	}
}

func makeCluster(clusterName string, clusterCfg *config.ClusterConfig) (*Cluster, error) {
	master := clusterCfg.GetMasterNode()
	DBName := master.DBName
	cluster := Cluster{nil, make(map[string]*sql.DB), clusterName, DBName, currentClusterVersion}

	db, err := openDBFromNodeCfg(master)
	if err != nil {
		return nil, err
	}
	cluster.masterNode = db

	for _, slave := range clusterCfg.GetSlaveNodes() {
		db, err := openDBFromNodeCfg(slave)
		if err != nil {
			return nil, err
		}
		cluster.slaveNodes[db.Dsn()] = db
	}
	return &cluster, nil
}

func New(clusterName string) (*Cluster, error) {
	clustersMu.RLock()
	defer clustersMu.RUnlock()
	if cluster, ok := clusterConns[clusterName]; ok {
		return cluster, nil
	} else {
		return nil, fmt.Errorf("ClusterName[%s] not exists", clusterName)
	}
}

func openDBFromNodeCfg(nodeCfg *config.NodeConfig) (*sql.DB, error) {
	if nodeCfg == nil {
		return nil, fmt.Errorf("OpenDBFromNodeCfg error nodeCfg==nil")
	}

	dsn := getDsnFromNodeCfg(nodeCfg)
	db, err := sql.Open("dbatman", dsn)
	if err != nil {
		log.Errorf("OpenDBFromNodeCfg sql.open error dsn:%s msg:%s", dsn, err.Error())
		return nil, err
	}
	setDBProperty(db, nodeCfg)
	return db, nil
}

func setDBProperty(db *sql.DB, nodeCfg *config.NodeConfig) {
	db.SetMaxOpenConns(nodeCfg.MaxConnections)
	db.SetMaxIdleConns(nodeCfg.MaxConnectionPoolSize)
}

func getDsnFromNodeCfg(nodeCfg *config.NodeConfig) string {
	if nodeCfg == nil {
		log.Errorf("Get dsn from NodeCfg error nodeCfg==nil")
		return ""
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&timeout=%dms",
		nodeCfg.Username,
		nodeCfg.Password,
		nodeCfg.Host,
		nodeCfg.Port,
		nodeCfg.DBName,
		nodeCfg.Charset,
		nodeCfg.ConnectTimeout)
}

func makeConnection(db *sql.DB) error {
	if db == nil {
		return fmt.Errorf("Can not make connection to database because of db is nil")
	}

	err := db.Ping()
	//TODO: retry
	if err != nil {
		log.Errorf("Make connection error msg:%s", err.Error())
	}
	return err
}
