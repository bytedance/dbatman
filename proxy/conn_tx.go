package proxy

import (
	. "github.com/bytedance/dbatman/database/mysql"
)

func (c *Session) isInTransaction() bool {
	return c.fc.Status()&uint16(StatusInTrans) > 0
}

func (c *Session) isAutoCommit() bool {
	return c.fc.Status()&uint16(StatusInAutocommit) > 0
}

func (c *Session) handleBegin() error {

	// We already in transaction
	if c.isInTransaction() {
		return c.fc.WriteOK(nil)
	}

	c.fc.XORStatus(uint16(StatusInTrans))
	if err := c.bc.begin(); err != nil {
		return c.handleMySQLError(err)
	}

	return c.fc.WriteOK(nil)
}

func (c *Session) handleCommit() (err error) {

	if !c.isInTransaction() {
		return c.fc.WriteOK(nil)
	}

	if err := c.bc.commit(); err != nil {
		return c.handleMySQLError(err)
	} else {
		return c.fc.WriteOK(nil)
	}
}

func (c *Session) handleRollback() (err error) {
	if !c.isInTransaction() {
		return c.fc.WriteOK(nil)
	}

	if err := c.bc.rollback(); err != nil {
		return c.handleMySQLError(err)
	}

	return c.fc.WriteOK(nil)
}
