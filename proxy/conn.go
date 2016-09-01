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

func (bc *SqlConn) commit(inAutoCommit bool) error {
	if bc.tx == nil {
		return errors.New("unexpect commit")
	}

	defer func() {
		if inAutoCommit {
			bc.tx = nil
		}
	}()

	if err := bc.tx.Commit(inAutoCommit); err != nil {
		// fmt.Println("commit err :", err)
		return err
	}

	return nil
}

func (bc *SqlConn) rollback(inAutoCommit bool) error {
	if bc.tx == nil {
		return errors.New("unexpect rollback")
	}

	defer func() {
		if inAutoCommit {
			bc.tx = nil
		}
	}()

	if err := bc.tx.Rollback(inAutoCommit); err != nil {
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
