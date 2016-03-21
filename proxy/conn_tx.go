package proxy

import (
	"github.com/bytedance/dbatman/database/mysql"
)

func (c *frontConn) isInTransaction() bool {
	return c.status&SERVER_STATUS_IN_TRANS > 0
}

func (c *frontConn) isAutoCommit() bool {
	return c.status&SERVER_STATUS_AUTOCOMMIT > 0
}

func (c *frontConn) handleBegin() error {
	c.status |= SERVER_STATUS_IN_TRANS
	return c.writeOK(nil)
}

func (c *frontConn) handleCommit() (err error) {
	if err := c.commit(); err != nil {
		return err
	} else {
		return c.writeOK(nil)
	}
}

func (c *frontConn) handleRollback() (err error) {
	if err := c.rollback(); err != nil {
		return err
	}

	return c.writeOK(nil)
}

func (c *frontConn) commit() (err error) {
	c.status &= ^SERVER_STATUS_IN_TRANS

	for _, co := range c.txConns {
		if e := co.Commit(); e != nil {
			err = e
		}
		co.Close()
	}

	c.txConns = map[*Node]*backend.SqlConn{}

	return
}

func (c *frontConn) rollback() (err error) {
	c.status &= ^SERVER_STATUS_IN_TRANS

	for _, co := range c.txConns {
		if e := co.Rollback(); e != nil {
			err = e
		}
		co.Close()
	}

	c.txConns = map[*Node]*backend.SqlConn{}

	return
}

//if status is in_trans, need
//else if status is not autocommit, need
//else no need
func (c *frontConn) needBeginTx() bool {
	return c.isInTransaction() || !c.isAutoCommit()
}
