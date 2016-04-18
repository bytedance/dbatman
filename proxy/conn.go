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
	stmt   *sql.Stmt
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

func (session *Session) DB(isread bool) *sql.DB {
	if isread {
		return session.bc.slave
	}

	return session.bc.master
}
