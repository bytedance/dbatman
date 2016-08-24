package proxy

import . "github.com/bytedance/dbatman/database/mysql"

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

	defer func() {
		if c.isInTransaction() {
			if c.isAutoCommit() {
				c.fc.AndStatus(uint16(^StatusInTrans))
				// fmt.Println("close the proxy tx")
			}
		}
	}()

	// fmt.Println("commit")
	// fmt.Println("this is a autocommit tx:", !c.isAutoCommit())
	if err := c.bc.commit(c.isAutoCommit()); err != nil {
		return c.handleMySQLError(err)
	} else {
		return c.fc.WriteOK(nil)
	}
}

func (c *Session) handleRollback() (err error) {
	if !c.isInTransaction() {
		return c.fc.WriteOK(nil)
	}

	defer func() {
		if c.isInTransaction() {
			if c.isAutoCommit() {
				c.fc.AndStatus(uint16(^StatusInTrans))
				// fmt.Println("close the proxy tx")
			}
		}
	}()
	// fmt.Println("rollback")
	// fmt.Println("this is a autocommit tx:", !c.isAutoCommit())
	if err := c.bc.rollback(c.isAutoCommit()); err != nil {
		return c.handleMySQLError(err)
	}

	return c.fc.WriteOK(nil)
}
