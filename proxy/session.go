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
	"github.com/bytedance/dbatman/cmd/version"
	"github.com/bytedance/dbatman/config"
	"github.com/bytedance/dbatman/database/cluster"
	. "github.com/bytedance/dbatman/database/mysql"
	"github.com/bytedance/dbatman/database/sql/driver"
	"github.com/bytedance/dbatman/hack"
	"github.com/ngaut/log"
	"net"
	"sync/atomic"
)

var DEFAULT_CAPABILITY uint32 = uint32(ClientLongPassword | ClientLongFlag |
	ClientConnectWithDB | ClientProtocol41 | ClientTransactions | ClientSecureConn)

var baseConnId uint32 = 10000

type Session struct {
	server *Server
	config *config.ProxyConfig
	user   *config.UserConfig

	connID     uint32
	status     uint16
	collation  CollationId
	charset    string
	capability uint32

	salt []byte

	cluster *cluster.Cluster
	fc      *MySQLServerConn

	closed bool
	db     string
}

func (s *Server) newSession(conn net.Conn) *Session {
	session := new(Session)

	session.server = s
	session.config = s.cfg.GetConfig()

	session.connID = atomic.AddUint32(&baseConnId, 1)
	session.status = uint16(StatusInAutocommit)
	session.capability = DEFAULT_CAPABILITY
	session.salt, _ = RandomBuf(20)

	session.collation = DEFAULT_COLLATION_ID
	session.charset = DEFAULT_CHARSET

	session.fc = NewMySQLServerConn(session, conn)

	return session
}

func (session *Session) HandshakeWithFront() error {

	return session.fc.Handshake()

	// TODO set cluster with auth info
	// session.cluster = cluster.New(db)
}

func (session *Session) Run() error {

	for {
		data, err := session.fc.ReadPacket()
		if err != nil {
			return err
		}

		if err := session.dispatch(data); err != nil {
			if err != driver.ErrBadConn {
				// TODO handle error
				// session.writeError(err)
				return nil
			}

			log.Warningf("con[%d], dispatch error %s", session.connID, err.Error())
			return err
		}

		if session.closed {
			return nil
		}

		session.ResetSequence()
	}

	return nil
}

func (session *Session) ResetSequence() {
	session.connID = 0
}

func (session *Session) Cap() uint32 {
	return session.capability
}

func (session *Session) SetCap(c uint32) {
	session.capability = c
}

func (session *Session) Status() uint16 {
	return session.status
}

func (session *Session) Collation() CollationId {
	return session.collation
}

func (session *Session) ConnID() uint32 {
	return session.connID
}

func (session *Session) DefaultDB() string {
	return session.db
}

func (session *Session) Salt() []byte {
	return session.salt
}

func (session *Session) ServerName() []byte {
	return hack.Slice(version.Version)
}
