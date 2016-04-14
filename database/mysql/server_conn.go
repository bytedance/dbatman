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

package mysql

import (
	"bytes"
	"github.com/bytedance/dbatman/database/sql/driver"
	"github.com/bytedance/dbatman/errors"
	"github.com/ngaut/log"
	"net"
	"sync/atomic"
)

// MySQLServerCtx is a server-side interface of 1-time-connection
// context
type MySQLServerCtx interface {
	Salt() []byte
	CheckAuth(username string, auth []byte, db string) error

	ServerName() []byte
}

// Connection between mysql client <-> mysql server
// here we wrap the go-mysql-driver.MySQLConn
type MySQLServerConn struct {
	*MySQLConn
	ctx       MySQLServerCtx
	collation CollationId
	connID    uint32
}

var baseConnId uint32 = 10000

func NewMySQLServerConn(s MySQLServerCtx, conn net.Conn) *MySQLServerConn {
	c := new(MySQLServerConn)

	c.ctx = s

	c.MySQLConn = &MySQLConn{
		maxPacketAllowed: maxPacketSize,
		maxWriteSize:     maxPacketSize - 1,
		netConn:          conn,
	}

	c.buf = newBuffer(c.netConn)

	// Default Capacity Flags
	c.flags = ClientLongPassword | ClientLongFlag |
		ClientConnectWithDB | ClientProtocol41 | ClientTransactions | ClientSecureConn

	// Set default connection id
	c.connID = atomic.AddUint32(&baseConnId, 1)
	c.status = StatusInAutocommit

	c.collation = DEFAULT_COLLATION_ID

	return c
}

// Hnadshake init handshake package to the client, wait for client autheticate
// response.
func (mc *MySQLServerConn) Handshake() error {
	var err error = nil

	// Handeshake
	if err = mc.writeInitPacket(); err != nil {
		mc.cleanup()
		return errors.Trace(err)
	}

	if err = mc.readHandshakeResponse(); err != nil {
		if e, ok := err.(*MySQLError); ok {
			mc.WriteError(e)
			err = nil
		}

		mc.cleanup()
		return errors.Trace(err)
	}

	// TODO here we should proceed PROTOCOL41 ?
	if err = mc.WriteOK(nil); err != nil {
		mc.cleanup()
		return errors.Trace(err)
	}

	mc.sequence = 0
	return nil
}

func (mc *MySQLServerConn) RemoteAddr() net.Addr {
	return mc.MySQLConn.netConn.RemoteAddr()
}

func (mc *MySQLServerConn) ResetSequence() {
	mc.sequence = 0
}

func (mc *MySQLServerConn) XORStatus(status uint16) {
	mc.status |= statusFlag(status)
}

func (mc *MySQLServerConn) AndStatus(status uint16) {
	mc.status &= statusFlag(status)
}

func (mc *MySQLServerConn) Status() uint16 {
	return uint16(mc.status)
}

func (mc *MySQLServerConn) SetFlags(flags uint32) {
	mc.flags = clientFlag(flags)
}

func (mc *MySQLServerConn) Flags() uint32 {
	return mc.Flags()
}

func (mc *MySQLServerConn) SetConnID(id uint32) {
	mc.connID = id
}

func (mc *MySQLServerConn) ConnID() uint32 {
	return mc.connID
}

func (mc *MySQLServerConn) SetCollation(id CollationId) {
	mc.collation = id
}

/******************************************************************************
*                   Server-Side Initialisation Process                        *
******************************************************************************/

// Handshake Initialization Packet
// http://dev.mysql.com/doc/internals/en/connection-phase-packets.html#packet-Protocol::Handshake
func (mc *MySQLServerConn) writeInitPacket() error {
	// preserved for write head
	data := make([]byte, 4, 128)

	// min version 10
	data = append(data, 10)

	// server version[00]
	data = append(data, mc.ctx.ServerName()...)
	data = append(data, 0)

	// connection id
	data = append(data, byte(mc.connID), byte(mc.connID), byte(mc.connID), byte(mc.connID))

	// auth-plugin-data-part-1
	data = append(data, mc.ctx.Salt()[0:8]...)

	// filter [00]
	data = append(data, 0)

	// capability flag lower 2 bytes, using default capability here
	data = append(data, byte(mc.flags), byte(mc.flags>>8))

	// collation, utf-8 default
	data = append(data, uint8(mc.collation))

	// status
	data = append(data, byte(mc.status), byte(mc.status>>8))

	// below 13 byte may not be used
	// capability flag upper 2 bytes, using default capability here
	data = append(data, byte(mc.flags>>16), byte(mc.flags>>24))

	// filter [0x15], for wireshark dump, value is 0x15
	data = append(data, 0x00)

	// reserved 10 [00]
	data = append(data, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)

	// auth-plugin-data-part-2
	data = append(data, mc.ctx.Salt()[8:]...)

	// filter [00]
	data = append(data, 0)

	if err := mc.writePacket(data); err != nil {
		return err
	}

	return nil
}

// ReadHandshakeResponse read the client handshake response, set the collations and
// capability, check the authetication info.
// for futher infomation, read the doc:
// http://dev.mysql.com/doc/internals/en/connection-phase-packets.html#packet-Protocol::Handshake
func (mc *MySQLServerConn) readHandshakeResponse() error {
	data, err := mc.readPacket()
	if err != nil {
		return err
	}

	pos := 0

	//capability
	mc.flags = clientFlag(endian.Uint32(data[:4]))
	pos += 4

	//skip max packet size
	pos += 4

	//charset, skip, if you want to use another charset, use set names
	//c.collation = CollationId(data[pos])
	pos++

	//skip reserved 23[00]
	pos += 23

	//user name
	user := string(data[pos : pos+bytes.IndexByte(data[pos:], 0)])
	pos += len(user) + 1

	//auth length and auth
	authLen := int(data[pos])
	pos++
	auth := data[pos : pos+authLen]
	pos += authLen

	if mc.flags&clientConnectWithDB == 0 {
		if err := mc.ctx.CheckAuth(user, auth, ""); err != nil {
			return err
		}
	} else {
		// connect must with db, otherwise it will deny the access
		if len(data[pos:]) == 0 {
			return errors.Trace(NewDefaultError(ER_ACCESS_DENIED_ERROR, mc.netConn.RemoteAddr().String(), user, "Yes"))
		}

		db := string(data[pos : pos+bytes.IndexByte(data[pos:], 0)])
		pos += len(db) + 1

		// check with user
		if err := mc.ctx.CheckAuth(user, auth, db); err != nil {
			log.Debugf("mysql check auth fail!")
			return err
		}
	}

	return nil
}

/******************************************************************************
*                   Function Send Packets to front client                     *
******************************************************************************/

// WriteError write error package to the client
func (mc *MySQLServerConn) WriteError(e *MySQLError) error {

	data := make([]byte, 4, 16+len(e.Message))

	data = append(data, ERR)
	data = append(data, byte(e.Number), byte(e.Number>>8))

	if mc.flags&clientProtocol41 > 0 {
		data = append(data, '#')
		data = append(data, e.State...)
	}

	data = append(data, e.Message...)

	return mc.writePacket(data)
}

// WriteOk write ok package to the client
func (mc *MySQLServerConn) WriteOK(r driver.Result) error {
	if r == nil {
		r = &MySQLResult{status: mc.status}
	}

	// Reserve 4 byte for packet header
	data := make([]byte, 4, 32)

	data = append(data, iOK)

	rows, _ := r.RowsAffected()
	insertId, _ := r.LastInsertId()

	data = appendLengthEncodedInteger(data, uint64(rows))
	data = appendLengthEncodedInteger(data, uint64(insertId))

	warnings := len(r.Warnings())
	status, _ := r.Status()

	if mc.flags&clientProtocol41 > 0 {
		data = append(data, byte(status), byte(status>>8))
		data = append(data, byte(warnings), byte(warnings>>8))
	}

	return mc.writePacket(data)
}

func (mc *MySQLServerConn) WriteEOF() error {
	data := make([]byte, 4, 9)

	data = append(data, iEOF)
	if mc.flags&ClientProtocol41 > 0 {
		data = append(data, 0, 0)
		data = append(data, byte(mc.status), byte(mc.status>>8))
	}

	return mc.writePacket(data)
}

func (mc *MySQLConn) WritePacket(data []byte) error {
	return mc.writePacket(data)
}

func (mc *MySQLConn) ReadPacket() ([]byte, error) {
	return mc.readPacket()
}

/******************************************************************************
*                   Function Wrapper for Export Visiable                      *
******************************************************************************/

func (mc *MySQLConn) HandleOkPacket(data []byte) error {
	return mc.handleOkPacket(data)
}

func (mc *MySQLConn) HandleErrorPacket(data []byte) error {
	return errors.Trace(mc.handleErrorPacket(data))
}
