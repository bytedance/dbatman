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
	"github.com/bytedance/dbatman/hack"
	"github.com/bytedance/dbatman/mysql"
	"github.com/juju/errors"
	"github.com/ngaut/log"
	"net"
	"sync"
	"sync/atomic"
	"runtime"
)

var DEFAULT_CAPABILITY uint32 = mysql.CLIENT_LONG_PASSWORD | mysql.CLIENT_LONG_FLAG |
	mysql.CLIENT_CONNECT_WITH_DB | mysql.CLIENT_PROTOCOL_41 |
	mysql.CLIENT_TRANSACTIONS | mysql.CLIENT_SECURE_CONNECTION

type frontConn struct {
	sync.Mutex

	pkg        *PacketIO
	conn       net.Conn
	server     *Server
	capability uint32
	connID     uint32

	status    uint16
	collation uint8
	charset   string

	user    string
	possdbs []string

	db   string
	salt []byte

	schema *Schema

	txConns map[*Node]*client.SqlConn

	closed bool

	lastInsertId int64
	affectedRows int64

	stmtId uint32

	stmts map[uint32]*Stmt
}

var baseConnId uint32 = 10000

func (s *Server) newConn(co net.Conn) *Conn {
	c := new(Conn)

	c.c = co

	c.pkg = NewPacketIO(co)

	c.server = s

	c.c = co
	c.pkg.Sequence = 0

	c.connectionId = atomic.AddUint32(&baseConnId, 1)

	c.status = SERVER_STATUS_AUTOCOMMIT

	c.salt, _ = RandomBuf(20)

	c.txConns = make(map[*Node]*client.SqlConn)

	c.closed = false

	c.collation = DEFAULT_COLLATION_ID
	c.charset = DEFAULT_CHARSET

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

func (c *Conn) Close() error {
	if c.closed {
		return nil
	}

	c.c.Close()

	c.rollback()
	for _, s := range c.stmts {
		s.Close()
	}

	c.stmts = nil

	c.closed = true

	return nil
}

func (c *Conn) writeInitialHandshake() error {

	// preserved for write head
	data := make([]byte, 4, 128)

	// min version 10
	data = append(data, 10)

	// server version[00]
	data = append(data, ServerVersion...)
	data = append(data, 0)

	// connection id
	data = append(data, byte(c.connectionId), byte(c.connectionId>>8), byte(c.connectionId>>16), byte(c.connectionId>>24))

	// auth-plugin-data-part-1
	data = append(data, c.salt[0:8]...)

	// filter [00]
	data = append(data, 0)

	// capability flag lower 2 bytes, using default capability here
	data = append(data, byte(DEFAULT_CAPABILITY), byte(DEFAULT_CAPABILITY>>8))

	// charset, utf-8 default
	data = append(data, uint8(DEFAULT_COLLATION_ID))

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

func (c *Conn) readPacket() ([]byte, error) {
	return c.pkg.ReadPacket()
}

func (c *Conn) writePacket(data []byte) error {
	return c.pkg.WritePacket(data)
}

func (c *Conn) readHandshakeResponse() error {
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

	if c.capability&CLIENT_CONNECT_WITH_DB == 0 {
		if err := c.checkAuth(auth); err != nil {
			return errors.Trace(err）
		}
	} else {
		// connect with db
		if len(data[pos:]) == 0 {
			return errors.Trace(
				NewDefaultError(ER_ACCESS_DENIED_ERROR, c.c.RemoteAddr().String(), c.user, "Yes")
				)
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

func (c *Conn) Run() {
	defer func() {
		r := recover()
		if r != nil {
			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Errorf("lastCmd %s, %v, %s", cc.lastCmd, r, buf)
		}
		c.Close()
	}()

	for {
		data, err := c.readPacket()
		if err != nil {
			return
		}

		if err := c.dispatch(data); err != nil {
			AppLog.Warn("con[%d], dispatch error %s", c.connectionId, err.Error())
			if err != ErrBadConn {
				c.writeError(err)
			}
		}

		if c.closed {
			return
		}

		c.pkg.Sequence = 0
	}
}

func (c *Conn) dispatch(data []byte) error {
	cmd := data[0]
	data = data[1:]

	switch cmd {
	case COM_QUIT:
		c.Close()
		return nil
	case COM_QUERY:
		return c.handleQuery(hack.String(data))
	case COM_PING:
		return c.writeOK(nil)
	case COM_INIT_DB:
		if err := c.useDB(hack.String(data)); err != nil {
			return err
		} else {
			return c.writeOK(nil)
		}
	case COM_FIELD_LIST:
		return c.handleFieldList(data)
	case COM_STMT_PREPARE:
		return c.handleComStmtPrepare(hack.String(data))
	case COM_STMT_EXECUTE:
		return c.handleComStmtExecute(data)
	case COM_STMT_CLOSE:
		return c.handleComStmtClose(data)
	case COM_STMT_SEND_LONG_DATA:
		return c.handleComStmtSendLongData(data)
	case COM_STMT_RESET:
		return c.handleComStmtReset(data)
	default:
		msg := fmt.Sprintf("command %d not supported now", cmd)
		AppLog.Warn(msg)
		return NewError(ER_UNKNOWN_ERROR, msg)
	}

	return nil
}

func (c *Conn) useDB(db string) error {
	if s := c.server.getSchema(db); s == nil {
		return NewDefaultError(ER_BAD_DB_ERROR, db)
	} else {
		c.schema = s
		c.db = db
	}
	return nil
}

func (c *Conn) writeOK(r *Result) error {
	if r == nil {
		r = &Result{Status: c.status}
	}
	data := make([]byte, 4, 32)

	data = append(data, OK_HEADER)

	data = append(data, PutLengthEncodedInt(r.AffectedRows)...)
	data = append(data, PutLengthEncodedInt(r.InsertId)...)

	if c.capability&CLIENT_PROTOCOL_41 > 0 {
		data = append(data, byte(r.Status), byte(r.Status>>8))
		data = append(data, byte(r.Warnings), byte(r.Warnings>>8))
	}

	return c.writePacket(data)
}

func (c *Conn) writeError(e error) error {
	var m *SqlError
	var ok bool
	if m, ok = e.(*SqlError); !ok {
		m = NewError(ER_UNKNOWN_ERROR, e.Error())
	}

	data := make([]byte, 4, 16+len(m.Message))

	data = append(data, ERR_HEADER)
	data = append(data, byte(m.Code), byte(m.Code>>8))

	if c.capability&CLIENT_PROTOCOL_41 > 0 {
		data = append(data, '#')
		data = append(data, m.State...)
	}

	data = append(data, m.Message...)

	return c.writePacket(data)
}

func (c *Conn) writeEOF(status uint16) error {
	data := make([]byte, 4, 9)

	data = append(data, EOF_HEADER)
	if c.capability&CLIENT_PROTOCOL_41 > 0 {
		data = append(data, 0, 0)
		data = append(data, byte(status), byte(status>>8))
	}

	return c.writePacket(data)
}

func (c *Conn) IsAutoCommit() bool {
	return c.status&SERVER_STATUS_AUTOCOMMIT > 0
}

func (c *Conn) checkDB() error {
	if c.schema != nil {
		return nil
	}

	if c.db != "" {
		return c.useDB(c.db)
	}

	return NewDefaultError(ER_NO_DB_ERROR)
}
