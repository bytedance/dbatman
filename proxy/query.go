// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package proxy

import (
	"fmt"
	"github.com/bytedance/dbatman/database/cluster"
	. "github.com/bytedance/dbatman/database/mysql"
	"github.com/bytedance/dbatman/database/sql"
	"github.com/bytedance/dbatman/database/sql/driver"
	"github.com/bytedance/dbatman/hack"
	"github.com/ngaut/log"
)

func (session *Session) Close() error {
	if session.closed {
		return nil
	}

	session.fc.Close()

	// TODO transaction
	//	session.rollback()

	// TODO stmts
	// for _, s := range session.stmts {
	// 	s.Close()
	// }

	// session.stmts = nil

	session.closed = true

	return nil
}

func (session *Session) dispatch(data []byte) error {
	cmd := data[0]
	data = data[1:]

	switch cmd {
	case ComQuit:
		session.Close()
		return nil
	case ComQuery:
		return session.comQuery(hack.String(data))
	case ComPing:
		return session.fc.WriteOK(nil)
	case ComInitDB:
		if err := session.useDB(hack.String(data)); err != nil {
			return err
		} else {
			return session.fc.WriteOK(nil)
		}
	case ComFieldList:
		// return session.handleFieldList(data)
		// TODO
		return nil
	case ComStmtPrepare:
		// TODO
		return nil
		// return session.handleComStmtPrepare(hack.String(data))
	case ComStmtExecute:
		// TODO
		return nil
		// return session.handleComStmtExecute(data)
	case ComStmtClose:
		// TODO
		return nil
		//return session.handleComStmtClose(data)
	case ComStmtSendLongData:
		// TODO
		return nil
		//return session.handleComStmtSendLongData(data)
	case ComStmtReset:
		// TODO
		return nil
		// return session.handleComStmtReset(data)
	default:
		msg := fmt.Sprintf("command %d not supported now", cmd)
		log.Warnf(msg)
		return NewDefaultError(ER_UNKNOWN_ERROR, msg)
	}

	return nil
}

func (session *Session) useDB(db string) error {

	if session.cluster != nil {
		if session.cluster.DBName != db {
			return NewDefaultError(ER_BAD_DB_ERROR, db)
		}

		return nil
	}

	if _, err := session.config.GetClusterByDBName(db); err != nil {
		return NewDefaultError(ER_BAD_DB_ERROR, db)
	} else if session.cluster, err = cluster.New(session.user.ClusterName); err != nil {
		return err
	}

	return nil
}

func (session *Session) IsAutoCommit() bool {
	return session.status&uint16(StatusInAutocommit) > 0
}

func (session *Session) checkDB() error {

	if session.cluster == nil {
		return NewDefaultError(ER_NO_DB_ERROR)
	}

	return nil
}

func (session *Session) WriteRows(rs *sql.Rows) error {
	var cols []driver.RawPacket
	var err error
	cols, err = rs.ColumnPackets()

	if err != nil {
		return session.handleError(err)
	}

	// Send a packet contains column length
	data := make([]byte, 4, 32)
	data = AppendLengthEncodedInteger(data, uint64(len(cols)))
	if err = session.fc.WritePacket(data); err != nil {
		return err
	}

	for _, col := range cols {
		if err := session.fc.WritePacket(col); err != nil {
			return err
		}
	}

	// TODO Write a ok packet
	session.fc.WriteOK(nil)

	for {
		payload, err := rs.NextRowPacket()
		if err != nil {
			if merr, ok := err.(*MySQLError); ok {
				session.fc.WriteError(merr)
				return nil
			}
			return err
		}

		if err := session.fc.WritePacket(payload); err != nil {
			return err
		}
	}

	// TODO Write a EOF packet
	session.fc.WriteEOF()

	return nil
}

func (session *Session) handleError(err error) error {
	switch inst := err.(type) {
	case *MySQLError:
		session.fc.WriteError(inst)
		return nil
	case *MySQLWarnings:
		// TODO process warnings
		session.fc.WriteOK(nil)
		return nil
	default:
		log.Errorf("handler default error: %v", err)
		return err
	}
}
