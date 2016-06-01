package proxy

import (
	"errors"
	"github.com/bytedance/dbatman/database/mysql"
)

// Wrap the connection
type SqlConn struct {
	master *mysql.DB
	slave  *mysql.DB
	stmts  map[uint32]*mysql.Stmt
	tx     *mysql.Tx

	session *Session
}

func (bc *SqlConn) begin() error {
	if bc.tx != nil {
		return errors.New("duplicate begin")
	}

	var err error
	bc.tx, err = bc.master.Begin()
	if err != nil {
		return err
	}

	return nil
}

func (bc *SqlConn) commit() error {
	if bc.tx == nil {
		return errors.New("unexpect commit")
	}

	defer func() {
		bc.tx = nil
	}()

	if err := bc.tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (bc *SqlConn) rollback() error {
	if bc.tx == nil {
		return errors.New("unexpect rollback")
	}

	defer func() {
		bc.tx = nil
	}()

	if err := bc.tx.Rollback(); err != nil {
		return err
	}

	return nil
}

func (session *Session) Executor(isread bool) mysql.Executor {

	// TODO set autocommit
	if session.isInTransaction() {
		return session.bc.tx
	}

	if isread {
		return session.bc.slave
	}

	return session.bc.master
}
