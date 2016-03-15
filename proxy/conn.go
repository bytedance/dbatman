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

// Copyright 2016 PinCAP, Inc.
// Copyright 2016 ByteDance, Inc.
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
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/bytedance/dbatman/Godeps/_workspace/src/github.com/juju/errors"
	"github.com/bytedance/dbatman/Godeps/_workspace/src/github.com/ngaut/log"
	"github.com/bytedance/dbatman/database/sql/driver/mysql"
	"github.com/bytedance/dbatman/hack"
	"net"
	"runtime"
	"sync"
	"sync/atomic"
)

var DEFAULT_CAPABILITY uint32 = mysql.CLIENT_LONG_PASSWORD | mysql.CLIENT_LONG_FLAG |
	mysql.CLIENT_CONNECT_WITH_DB | mysql.CLIENT_PROTOCOL_41 |
	mysql.CLIENT_TRANSACTIONS | mysql.CLIENT_SECURE_CONNECTION

type frontConn struct {
	sync.Mutex

	pkg        *mysql.PacketIO
	conn       net.Conn
	server     *Server
	capability uint32
	connID     uint32

	status    uint16
	collation mysql.CollationId
	charset   string

	user    string
	possdbs []string

	db   string
	salt []byte

	schema *Schema

	txConns map[*Node]*backend.SqlConn

	closed bool

	lastInsertId int64
	affectedRows int64
	lastCmd      string

	stmtId uint32

	stmts map[uint32]*Stmt
}

var baseConnId uint32 = 10000

func (s *Server) newConn(co net.Conn) *frontConn {
	c := new(frontConn)

	c.conn = co

	c.pkg = mysql.NewPacketIO(co)

	c.server = s

	c.conn = co
	c.pkg.Sequence = 0

	c.connID = atomic.AddUint32(&baseConnId, 1)

	c.status = mysql.SERVER_STATUS_AUTOCOMMIT

	c.salt, _ = mysql.RandomBuf(20)

	c.txConns = make(map[*Node]*backend.SqlConn)

	c.closed = false

	c.collation = mysql.DEFAULT_COLLATION_ID
	c.charset = mysql.DEFAULT_CHARSET

	c.stmtId = 0
	c.stmts = make(map[uint32]*Stmt)

	return c
}

func (c *frontConn) handshake() error {
	if err := c.writeInitialHandshake(); err != nil {
		return errors.Trace(err)
	}

	if err := c.readHandshakeResponse(); err != nil {
		c.writeError(err)
		return errors.Trace(err)
	}

	// TODO here we should proceed PROTOCOL41 ?

	if err := c.writeOK(nil); err != nil {
		return errors.Trace(err)
	}

	c.pkg.Sequence = 0

	return errors.Trace(c.flush())
}

func (c *frontConn) flush() error {
	return c.pkg.Flush()
}

func (c *frontConn) Close() error {
	if c.closed {
		return nil
	}

	c.conn.Close()

	c.rollback()
	for _, s := range c.stmts {
		s.Close()
	}

	c.stmts = nil

	c.closed = true

	return nil
}

func (c *frontConn) writeInitialHandshake() error {

	// preserved for write head
	data := make([]byte, 4, 128)

	// min version 10
	data = append(data, 10)

	// server version[00]
	data = append(data, mysql.ServerVersion...)
	data = append(data, 0)

	// connection id
	data = append(data, byte(c.connID), byte(c.connID>>8), byte(c.connID>>16), byte(c.connID>>24))

	// auth-plugin-data-part-1
	data = append(data, c.salt[0:8]...)

	// filter [00]
	data = append(data, 0)

	// capability flag lower 2 bytes, using default capability here
	data = append(data, byte(DEFAULT_CAPABILITY), byte(DEFAULT_CAPABILITY>>8))

	// charset, utf-8 default
	data = append(data, uint8(mysql.DEFAULT_COLLATION_ID))

	// status
	data = append(data, byte(c.status), byte(c.status>>8))

	// below 13 byte may not be used
	// capability flag upper 2 bytes, using default capability here
	data = append(data, byte(DEFAULT_CAPABILITY>>16), byte(DEFAULT_CAPABILITY>>24))

	// filter [0x15], for wireshark dump, value is 0x15
	data = append(data, 0x00)

	// reserved 10 [00]
	data = append(data, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)

	// auth-plugin-data-part-2
	data = append(data, c.salt[8:]...)

	// filter [00]
	data = append(data, 0)

	if err := c.writePacket(data); err != nil {
		return errors.Trace(err)
	}

	return errors.Trace(c.flush())
}

func (c *frontConn) readPacket() ([]byte, error) {
	return c.pkg.ReadPacket()
}

func (c *frontConn) writePacket(data []byte) error {
	return c.pkg.WritePacket(data)
}

func (c *frontConn) readHandshakeResponse() error {
	data, err := c.readPacket()
	if err != nil {
		return errors.Trace(err)
	}

	pos := 0

	//capability
	c.capability = binary.LittleEndian.Uint32(data[:4])
	pos += 4

	//skip max packet size
	pos += 4

	//charset, skip, if you want to use another charset, use set names
	//c.collation = CollationId(data[pos])
	pos++

	//skip reserved 23[00]
	pos += 23

	//user name
	c.user = string(data[pos : pos+bytes.IndexByte(data[pos:], 0)])
	pos += len(c.user) + 1

	//auth length and auth
	authLen := int(data[pos])
	pos++
	auth := data[pos : pos+authLen]
	pos += authLen

	if c.capability&mysql.CLIENT_CONNECT_WITH_DB == 0 {
		if err := c.checkAuth(auth); err != nil {
			return errors.Trace(err)
		}
	} else {
		// connect with db
		if len(data[pos:]) == 0 {
			return errors.Trace(mysql.NewDefaultError(mysql.ER_ACCESS_DENIED_ERROR, c.conn.RemoteAddr().String(), c.user, "Yes"))
		}

		db := string(data[pos : pos+bytes.IndexByte(data[pos:], 0)])
		pos += len(c.db) + 1

		// check with db multi-user
		if err := c.checkAuthWithDB(auth, db); err != nil {
			return errors.Trace(err)
		}
	}

	return nil
}

func (c *frontConn) Run() {
	defer func() {
		r := recover()
		if r != nil {
			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Errorf("lastCmd %s, %v, %s", c.lastCmd, r, buf)
		}
		c.Close()
	}()

	for {
		data, err := c.readPacket()
		if err != nil {
			return
		}

		if err := c.dispatch(data); err != nil {
			log.Warnf("con[%d], dispatch error %s", c.connID, err.Error())
			if err != mysql.ErrBadConn {
				c.writeError(err)
			}
		}

		if c.closed {
			return
		}

		c.pkg.Sequence = 0
	}
}

func (c *frontConn) dispatch(data []byte) error {
	cmd := data[0]
	data = data[1:]

	switch cmd {
	case mysql.COM_QUIT:
		c.Close()
		return nil
	case mysql.COM_QUERY:
		return c.handleQuery(hack.String(data))
	case mysql.COM_PING:
		return c.writeOK(nil)
	case mysql.COM_INIT_DB:
		if err := c.useDB(hack.String(data)); err != nil {
			return err
		} else {
			return c.writeOK(nil)
		}
	case mysql.COM_FIELD_LIST:
		return c.handleFieldList(data)
	case mysql.COM_STMT_PREPARE:
		return c.handleComStmtPrepare(hack.String(data))
	case mysql.COM_STMT_EXECUTE:
		return c.handleComStmtExecute(data)
	case mysql.COM_STMT_CLOSE:
		return c.handleComStmtClose(data)
	case mysql.COM_STMT_SEND_LONG_DATA:
		return c.handleComStmtSendLongData(data)
	case mysql.COM_STMT_RESET:
		return c.handleComStmtReset(data)
	default:
		msg := fmt.Sprintf("command %d not supported now", cmd)
		log.Warnf(msg)
		return mysql.NewError(mysql.ER_UNKNOWN_ERROR, msg)
	}

	return nil
}

func (c *frontConn) useDB(db string) error {
	if s := c.server.getSchema(db); s == nil {
		return mysql.NewDefaultError(mysql.ER_BAD_DB_ERROR, db)
	} else {
		c.schema = s
		c.db = db
	}
	return nil
}

func (c *frontConn) writeOK(r *mysql.Result) error {
	if r == nil {
		r = &mysql.Result{Status: c.status}
	}
	data := make([]byte, 4, 32)

	data = append(data, mysql.OK_HEADER)

	data = append(data, mysql.PutLengthEncodedInt(r.AffectedRows)...)
	data = append(data, mysql.PutLengthEncodedInt(r.InsertId)...)

	if c.capability&mysql.CLIENT_PROTOCOL_41 > 0 {
		data = append(data, byte(r.Status), byte(r.Status>>8))
		data = append(data, byte(r.Warnings), byte(r.Warnings>>8))
	}

	return c.writePacket(data)
}

func (c *frontConn) writeError(e error) error {
	var m *mysql.SqlError
	var ok bool
	if m, ok = e.(*mysql.SqlError); !ok {
		m = mysql.NewError(mysql.ER_UNKNOWN_ERROR, e.Error())
	}

	data := make([]byte, 4, 16+len(m.Message))

	data = append(data, mysql.ERR_HEADER)
	data = append(data, byte(m.Code), byte(m.Code>>8))

	if c.capability&mysql.CLIENT_PROTOCOL_41 > 0 {
		data = append(data, '#')
		data = append(data, m.State...)
	}

	data = append(data, m.Message...)

	return c.writePacket(data)
}

func (c *frontConn) writeEOF(status uint16) error {
	data := make([]byte, 4, 9)

	data = append(data, mysql.EOF_HEADER)
	if c.capability&mysql.CLIENT_PROTOCOL_41 > 0 {
		data = append(data, 0, 0)
		data = append(data, byte(status), byte(status>>8))
	}

	return c.writePacket(data)
}

func (c *frontConn) IsAutoCommit() bool {
	return c.status&mysql.SERVER_STATUS_AUTOCOMMIT > 0
}

func (c *frontConn) checkDB() error {
	if c.schema != nil {
		return nil
	}

	if c.db != "" {
		return c.useDB(c.db)
	}

	return mysql.NewDefaultError(mysql.ER_NO_DB_ERROR)
}
