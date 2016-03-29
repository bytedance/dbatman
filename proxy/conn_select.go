package proxy

import (
	"fmt"
	. "github.com/bytedance/dbatman/database/mysql"
	"github.com/bytedance/dbatman/database/sql"
	"github.com/bytedance/dbatman/database/sql/driver"
	"github.com/bytedance/dbatman/parser"
)

func (session *Session) handleQuery(stmt parser.IStatement, sqlstmt string) error {

	if err := session.checkDB(); err != nil {
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
		// r := c.newEmptyResultset(stmt)
		// return c.writeResultset(c.status, r)
		return fmt.Errorf("no available backend db")
	}

	var rs *sql.Rows
	rs, err = db.Query(sqlstmt)

	// TODO here should handler error
	if err != nil {
		return err
	}

	defer rs.Close()

	var cols []driver.RawPayload
	cols, err = rs.ColumnPackets()
	if err != nil {
		return err
	}

	for _, col := range cols {
		if err := session.fc.WritePacket(col); err != nil {
			return err
		}
	}

	// TODO Write a ok packet

	var payload driver.RawPayload
	for {
		payload, err := rs.NextRowPayload()
		if err != nil {
			if merr, ok := err.(*MySQLError); ok {
				session.fc.WriteError(merr)
			}
			return err
		}

		if err := session.fc.WritePacket(payload); err != nil {
			return err
		}
	}

	// TODO Write a EOF packet

	return nil
}
