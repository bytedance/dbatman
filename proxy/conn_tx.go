package proxy

import (
	"github.com/bytedance/dbatman/database/mysql"
)

func (c *Session) isInTransaction() bool {
	return c.status&mysql.SERVER_STATUS_IN_TRANS > 0
}

func (c *Session) isAutoCommit() bool {
	return c.status&SERVER_STATUS_AUTOCOMMIT > 0
}

func (c *Session) handleBegin() error {
	c.status |= SERVER_STATUS_IN_TRANS
	return c.writeOK(nil)
}

func (c *Session) handleCommit() (err error) {
	if err := c.commit(); err != nil {
		return err
	} else {
		return c.writeOK(nil)
	}
}

func (c *Session) handleRollback() (err error) {
	if err := c.rollback(); err != nil {
		return err
	}

	return c.writeOK(nil)
}

func (c *Session) commit() (err error) {
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

func (c *Session) rollback() (err error) {
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
func (c *Session) needBeginTx() bool {
	return c.isInTransaction() || !c.isAutoCommit()
}
