package proxy

import (
	"github.com/bytedance/dbatman/parser"
	"github.com/ngaut/log"
)

func (session *Session) handleQuery(stmt parser.IStatement, sqlstmt string) error {

	if err := session.checkDB(stmt); err != nil {
		log.Debugf("check db error: %s", err.Error())
		return err
	}

	isread := false

	if s, ok := stmt.(parser.ISelect); ok {
		isread = !s.IsLocked()
	} else if _, sok := stmt.(parser.IShow); sok {
		isread = true
	}

	rs, err := session.Executor(isread).Query(sqlstmt)
	// TODO here should handler error
	if err != nil {
		return session.handleMySQLError(err)
	}

	defer rs.Close()

	return session.writeRows(rs)
}
