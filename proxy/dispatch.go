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
	"github.com/bytedance/dbatman/database/mysql"
	"github.com/bytedance/dbatman/database/sql/driver"
	"github.com/bytedance/dbatman/hack"
	"github.com/ngaut/log"
	"io"
)

func (session *Session) dispatch(data []byte) (err error) {
	cmd := data[0]
	data = data[1:]

	defer func() {
		err = session.fc.Flush()
	}()

	switch cmd {
	case mysql.ComQuit:
		session.Close()
		err = nil
	case mysql.ComQuery:
		err = session.comQuery(hack.String(data))
	case mysql.ComPing:
		err = session.fc.WriteOK(nil)
	case mysql.ComInitDB:
		if err := session.useDB(hack.String(data)); err != nil {
			err = session.handleMySQLError(err)
		} else {
			err = session.fc.WriteOK(nil)
		}
	case mysql.ComFieldList:
		err = session.handleFieldList(data)
	case mysql.ComStmtPrepare:
		err = session.handleComStmtPrepare(hack.String(data))
	case mysql.ComStmtExecute:
		err = session.handleComStmtExecute(data)
	case mysql.ComStmtClose:
		err = session.handleComStmtClose(data)
	case mysql.ComStmtSendLongData:
		// TODO
		//return session.handleComStmtSendLongData(data)
	case mysql.ComStmtReset:
		// TODO
		// return session.handleComStmtReset(data)
	default:
		msg := fmt.Sprintf("command %d not supported now", cmd)
		log.Warnf(msg)
		err = mysql.NewDefaultError(mysql.ER_UNKNOWN_ERROR, msg)
	}

	return
}

func (session *Session) useDB(db string) error {

	if session.cluster != nil {
		if session.cluster.DBName != db {
			return mysql.NewDefaultError(mysql.ER_BAD_DB_ERROR, db)
		}

		return nil
	}

	if _, err := session.config.GetClusterByDBName(db); err != nil {
		return mysql.NewDefaultError(mysql.ER_BAD_DB_ERROR, db)
	} else if session.cluster, err = cluster.New(session.user.ClusterName); err != nil {
		return err
	}

	if session.bc == nil {
		master, err := session.cluster.Master()
		if err != nil {
			return mysql.NewDefaultError(mysql.ER_BAD_DB_ERROR, db)
		}
		slave, err := session.cluster.Slave()
		if err != nil {
			slave = master
		}
		session.bc = &SqlConn{
			master:  master,
			slave:   slave,
			stmts:   make(map[uint32]*mysql.Stmt),
			tx:      nil,
			session: session,
		}
	}

	return nil
}

func (session *Session) IsAutoCommit() bool {
	return session.fc.Status()&uint16(mysql.StatusInAutocommit) > 0
}

func (session *Session) writeRows(rs mysql.Rows) error {
	var cols []driver.RawPacket
	var err error
	cols, err = rs.ColumnPackets()

	if err != nil {
		return session.handleMySQLError(err)
	}

	// Send a packet contains column length
	data := make([]byte, 4, 32)
	data = mysql.AppendLengthEncodedInteger(data, uint64(len(cols)))
	if err = session.fc.WritePacket(data); err != nil {
		return err
	}

	// Write Columns Packet
	for _, col := range cols {
		if err := session.fc.WritePacket(col); err != nil {
			log.Debugf("write columns packet error %v", err)
			return err
		}
	}

	// TODO Write a ok packet
	if err = session.fc.WriteEOF(); err != nil {
		return err
	}

	for {
		packet, err := rs.NextRowPacket()

		// Handle Error

		if err != nil {
			if err == io.EOF {
				return session.fc.WriteEOF()
			} else {
				return session.handleMySQLError(err)
			}
		}

		if err := session.fc.WritePacket(packet); err != nil {
			return err
		}
	}

	return nil
}

func (session *Session) handleMySQLError(e error) error {

	switch inst := e.(type) {
	case *mysql.MySQLError:
		session.fc.WriteError(inst)
		return nil
	default:
		log.Warnf("default error: %T %s", e, e)
		return e
	}
}
