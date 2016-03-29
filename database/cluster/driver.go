package cluster

import (
	"github.com/bytedance/dbatman/database/sql"
)

type Cluster struct {
}

func New(db string) (*Cluster, error) {
	return nil, nil
}

func (c *Cluster) Close() error {
	return nil
}

func (c *Cluster) DB(slave bool) (*sql.DB, error) {
	if slave {
		return c.Master()
	}

	return c.Slave()
}

// TODO
func (c *Cluster) Master() (*sql.DB, error) {
	return nil, nil
}

// TODO
func (c *Cluster) Slave() (*sql.DB, error) {
	return nil, nil
}
