package config

import (
	"fmt"
	"github.com/ngaut/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

var conf Conf

type Conf struct {
	path             string
	lastModifiedTime time.Time
	mu               sync.RWMutex
	proxyConfig      *ProxyConfig
}

type ProxyConfig struct {
	Global   *GlobalConfig             `yaml:"global"`
	Clusters map[string]*ClusterConfig `yaml:"clusters"`
	Users    map[string]*UserConfig    `yaml:"users"`
}

type GlobalConfig struct {
	Port              int
	ManagePort        int      `yaml:"manage_port"`
	MaxConnections    int      `yaml:"max_connections"`
	LogFilename       string   `yaml:"log_filename"`
	LogLevel          int      `yaml:"log_level"`
	LogMaxSize        int      `yaml:"log_maxsize"`
	ClientTimeout     int      `yaml:"client_timeout"`
	ServerTimeout     int      `yaml:"server_timeout"`
	WriteTimeInterval int      `yaml:"write_time_interval"`
	ConfAutoload      int      `yaml:"conf_autoload"`
	AuthIPs           []string `yaml:"auth_ips,omitempty"`
}

type ClusterConfig struct {
	Master *NodeConfig
	Slaves []*NodeConfig
}

type NodeConfig struct {
	Host                  string
	Port                  int
	Username              string
	Password              string
	DBName                string
	Charset               string
	Weight                int
	MaxConnections        int `yaml:"max_connections"`
	MaxConnectionPoolSize int `yaml:"max_connection_pool_size"`
	ConnectTimeout        int `yaml:"connect_timeout"`
	TimeReconnectInterval int `yaml:"time_reconnect_interval"`
}

type UserConfig struct {
	Username       string
	Password       string
	DBName         string
	Charset        string
	MaxConnections int      `yaml:"max_connections"`
	MinConnections int      `yaml:"min_connections"`
	ClusterName    string   `yaml:"cluster_name"`
	AuthIPs        []string `yaml:"auth_ips,omitempty"`
	BlackListIPs   []string `yaml:"black_list_ips,omitempty"`
}

func (p *ProxyConfig) GetAllClusters() (map[string]*ClusterConfig, error) {
	if p.Clusters == nil {
		err := fmt.Errorf("GetAllClusters p.Clusters==nil")
		return nil, err
	}
	return p.Clusters, nil
}

// GetClusterByDBName return all cluster by given dbname
func (p *ProxyConfig) GetClusterByDBName(dbName string) (*ClusterConfig, error) {
	if p.Clusters == nil {
		err := fmt.Errorf("GetClusterByDBName p.Clusters==nil")
		return nil, err
	}

	for _, cluster := range p.Clusters {
		if cluster.Master.DBName == dbName {
			return cluster, nil
		}
	}

	err := fmt.Errorf("GetClusterByDBName DB %s not exists", dbName)
	return nil, err
}

func (p *ProxyConfig) GetMasterNodefromClusterByName(clusterName string) (*NodeConfig, error) {
	if p.Clusters == nil {
		err := fmt.Errorf("GetMasterNodefromClusterByName p.Clusters==nil")
		return nil, err
	}
	node := p.Clusters[clusterName]
	if node == nil || node.Master == nil {
		err := fmt.Errorf("GetMasterNodefromClusterByName cluster %s do not exist", clusterName)
		return nil, err
	}
	return node.Master, nil
}

func (p *ProxyConfig) GetSlaveNodesfromClusterByName(clusterName string) ([]*NodeConfig, error) {
	if p.Clusters == nil {
		err := fmt.Errorf("GetSlaveNodesfromCluster p.Clusters==nil")
		return nil, err
	}
	node := p.Clusters[clusterName]
	if node == nil {
		err := fmt.Errorf("GetSlaveNodesfromCluster cluster %s do not exist", clusterName)
		return nil, err
	}
	return node.Slaves, nil
}

func (p *ProxyConfig) GetUserByName(username string) (*UserConfig, error) {
	if p.Users == nil {
		err := fmt.Errorf("p.Users==nil")
		return nil, err
	}
	user := p.Users[username]
	if user == nil {
		err := fmt.Errorf("GetUserByName user %s do not exist", username)
		return nil, err
	}
	return user, nil
}

func (p *ProxyConfig) ServerTimeout() int {
	return p.Global.ServerTimeout
}

func (cc *ClusterConfig) GetMasterNode() *NodeConfig {
	return cc.Master
}

func (cc *ClusterConfig) GetSlaveNodes() []*NodeConfig {
	return cc.Slaves
}

func (c *Conf) parseConfigFile(proxyConfig *ProxyConfig) error {
	data, err := ioutil.ReadFile(c.path)
	if err == nil {
		err = yaml.Unmarshal([]byte(data), proxyConfig)
		if err == nil {
			if !validateConfig(proxyConfig) {
				err = fmt.Errorf("config is invalidate")
			}
		}

	}
	return err
}

func (c *Conf) GetConfig() *ProxyConfig {
	c.mu.RLock()
	proxyConfig := c.proxyConfig
	c.mu.RUnlock()
	return proxyConfig
}

func (c *Conf) CheckConfigUpdate(notifyChans ...chan bool) {
	if c.proxyConfig.Global.ConfAutoload == 1 {
		for {
			time.Sleep(time.Second * 60)
			log.Infof("CheckConfigUpdate checking")
			fileinfo, err := os.Stat(c.path)
			if err != nil {
				log.Errorf("CheckConfigUpdate error %s", err.Error())
				continue
			}
			//config been modified
			if c.lastModifiedTime.Before(fileinfo.ModTime()) {
				log.Infof("CheckConfigUpdate config change and load new config")
				defaultProxyConfig := getDefaultProxyConfig()
				err = c.parseConfigFile(defaultProxyConfig)
				if err != nil {
					log.Errorf("CheckConfigUpdate error %s", err.Error())
					continue
				}
				c.lastModifiedTime = fileinfo.ModTime()
				//goroutine need mutex lock
				c.mu.Lock()
				c.proxyConfig = defaultProxyConfig
				c.mu.Unlock()
				log.Infof("CheckConfigUpdate new config load success")
				for _, notifyChan := range notifyChans {
					notifyChan <- true
				}
			}
		}

	}
}

func LoadConfig(path string) (*Conf, error) {
	fileinfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	conf.path = path
	defaultProxyConfig := getDefaultProxyConfig()
	err = conf.parseConfigFile(defaultProxyConfig)
	if err != nil {
		return nil, err
	}
	conf.lastModifiedTime = fileinfo.ModTime()
	conf.proxyConfig = defaultProxyConfig
	return &conf, err
}

func validateConfig(cfg *ProxyConfig) bool {
	if cfg == nil {
		return false
	}

	if len(cfg.Clusters) == 0 {
		log.Errorf("ValidateConfig 0 cluster")
		return false
	}

	if len(cfg.Users) == 0 {
		log.Errorf("ValidateConfig 0 user")
		return false
	}

	for username, user := range cfg.Users {
		clusterName := user.ClusterName
		if _, ok := cfg.Clusters[clusterName]; !ok {
			log.Errorf("ValidateConfig cluster %s belong to user %s do not exist", clusterName, username)
			return false
		}
	}

	for clusterName, cluster := range cfg.Clusters {
		if cluster.Master == nil {
			log.Errorf("ValidateConfig cluster %s do not have master node", clusterName)
			return false
		}

		master := cluster.Master
		if master.MaxConnections < master.MaxConnectionPoolSize {
			log.Errorf("ValidateConfig cluster %s master MaxConnectionPoolSize more than MaxConnections", clusterName)
			return false
		}

		if cluster.Slaves != nil {
			for _, slave := range cluster.Slaves {
				if slave.MaxConnections < slave.MaxConnectionPoolSize {
					log.Errorf("ValidateConfig cluster %s slave MaxConnectionPoolSize more than MaxConnections", clusterName)
					return false
				}
			}
		}
	}

	return true
}

func getDefaultProxyConfig() *ProxyConfig {
	cfg := ProxyConfig{
		Global: &GlobalConfig{
			Port:              3306,
			ManagePort:        3307,
			MaxConnections:    2000,
			LogLevel:          1,
			LogFilename:       "./log/dbatman.log",
			LogMaxSize:        2014,
			ClientTimeout:     1800,
			ServerTimeout:     1800,
			WriteTimeInterval: 10,
			ConfAutoload:      1,
			AuthIPs:           []string{"127.0.0.1"},
		},
	}
	return &cfg
}
