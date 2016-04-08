// Copyright 2013 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// The MIT License (MIT)
//
// Copyright (c) 2014 siddontang
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// Copyright 2016 PinCAP, Insession.
// Copyright 2016 ByteDance, Insession.
//
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
	"github.com/bytedance/dbatman/Godeps/_workspace/src/github.com/ngaut/log"
	. "github.com/bytedance/dbatman/database/mysql"
	"github.com/bytedance/dbatman/database/sql"
	"github.com/bytedance/dbatman/database/sql/driver"
	"github.com/bytedance/dbatman/hack"
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
		log.Debugf("ComQuery: %s", hack.String(data))
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
	if _, err := session.config.GetClusterByDBName(db); err != nil {
		return NewDefaultError(ER_BAD_DB_ERROR, db)
	} else {
		session.db = db
	}
	return nil
}

func (session *Session) IsAutoCommit() bool {
	return session.status&uint16(StatusInAutocommit) > 0
}

func (session *Session) checkDB() error {

	if session.db != "" {
		return session.useDB(session.db)
	}

	return NewDefaultError(ER_NO_DB_ERROR)
}

func (session *Session) WriteRows(rs *sql.Rows) error {
	var cols []driver.RawPayload
	var err error
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
