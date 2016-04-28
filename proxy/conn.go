package proxy

import (
	"errors"
	"github.com/bytedance/dbatman/database/sql"
	juju "github.com/bytedance/dbatman/errors"
)

// Wrap the connection
type SqlConn struct {
	master *sql.DB
	slave  *sql.DB
	stmts  map[uint32]*sql.Stmt
	tx     *sql.Tx

	session *Session
}

func (bc *SqlConn) begin() error {
	if bc.tx != nil {
		return juju.Trace(errors.New("duplicate begin"))
	}

	var err error
	bc.tx, err = bc.master.Begin()
	if err != nil {
		return juju.Trace(err)
	}

	return nil
}

func (bc *SqlConn) commit() error {
	if bc.tx == nil {
		return juju.Trace(errors.New("unexpect commit"))
	}

	defer func() {
		bc.tx = nil
	}()

	if err := bc.tx.Commit(); err != nil {
		return juju.Trace(err)
	}

	return nil
}

func (bc *SqlConn) rollback() error {
	if bc.tx == nil {
		return juju.Trace(errors.New("unexpect rollback"))
	}

	defer func() {
		bc.tx = nil
	}()

	if err := bc.tx.Rollback(); err != nil {
		return juju.Trace(err)
	}

	return nil
}

func (session *Session) Executor(isread bool) sql.Executor {

	// TODO set autocommit
	if session.isInTransaction() {
		return session.bc.tx
	}

	if isread {
		return session.bc.slave
	}

	return session.bc.master
}
