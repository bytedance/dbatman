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
	"encoding/binary"
)

type MySQLServer interface {
	ConnID() uint32
	Salt() []byte
	Collation() uint8
	Status() uint16
	Cap() uint32
}

type MySQLServerConn struct {
	*MySQLConn
	MySQLServer
}

func NewMySQLConn(s MySQLServer) *MySQLServerConn {
	return &MySQLServerConn{
		MySQLServer: s,
	}
}

/******************************************************************************
*                           Initialisation Process                            *
******************************************************************************/

const serverVersion = "dbatman mysql-proxy 1.0"

// Handshake Initialization Packet
// http://dev.mysql.com/doc/internals/en/connection-phase-packets.html#packet-Protocol::Handshake
func (mc *MySQLServerConn) writeInitPacket() error {
	// preserved for write head
	data := make([]byte, 4, 128)

	// min version 10
	data = append(data, 10)

	// server version[00]
	data = append(data, serverVersion...)
	data = append(data, 0)

	// connection id
	data = append(data, byte(mc.ConnID()), byte(mc.ConnID()>>8), byte(mc.ConnID()>>16), byte(mc.ConnID()>>24))

	// auth-plugin-data-part-1
	data = append(data, mc.Salt()[0:8]...)

	// filter [00]
	data = append(data, 0)

	// capability flag lower 2 bytes, using default capability here
	data = append(data, byte(mc.Cap()), byte(mc.Cap()>>8))

	// charset, utf-8 default
	data = append(data, uint8(mc.Collation()))

	// status
	data = append(data, byte(mc.Status()), byte(mc.Status()>>8))

	// below 13 byte may not be used
	// capability flag upper 2 bytes, using default capability here
	data = append(data, byte(mc.Cap()>>16), byte(mc.Cap()>>24))

	// filter [0x15], for wireshark dump, value is 0x15
	data = append(data, 0x00)

	// reserved 10 [00]
	data = append(data, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)

	// auth-plugin-data-part-2
	data = append(data, mc.Salt()[8:]...)

	// filter [00]
	data = append(data, 0)

	if err := mc.writePacket(data); err != nil {
		return err
	}

	return nil
}

func (mc *MySQLConn) readHandshakeResponse() error {
	data, err := mc.readPacket()
	if err != nil {
		return err
	}

	pos := 0

	//capability
	mc.capability = binary.LittleEndian.Uint32(data[:4])
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
			return err
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

func (mc *MySQLConn) HandleOkPacket(data []byte) error {
	return mc.handleOkPacket(data)
}

func (mc *MySQLConn) HandleErrorPacket(data []byte) error {
	return mc.handleErrorPacket(data)
}
