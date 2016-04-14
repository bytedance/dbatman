package proxy

import (
	"fmt"
	"github.com/bytedance/dbatman/database/sql"
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

	db, err := session.cluster.DB(isread)

	// TODO here if db is nil, then we should return a error?
	if err != nil {
		return err
	} else if db == nil {
		// TODO error process
		return fmt.Errorf("no available backend db")
	}

	var rs *sql.Rows
	rs, err = db.Query(sqlstmt)

	// TODO here should handler error
	if err != nil {
		return err
	}

	defer rs.Close()

	return session.WriteRows(rs)
}
