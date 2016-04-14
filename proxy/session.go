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
)

type Session struct {
	server *Server
	config *config.ProxyConfig
	user   *config.UserConfig

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
	session.salt, _ = RandomBuf(20)

	session.fc = NewMySQLServerConn(session, conn)

	return session
}

func (session *Session) Handshake() error {

	if err := session.fc.Handshake(); err != nil {
		return err
	}

	if err := session.useDB(session.user.DBName); err != nil {
		return err
	}

	return nil
}

func (session *Session) Run() error {

	for {

		data, err := session.fc.ReadPacket()

		if err != nil {
			log.Warn(err)
			return err
		}

		if err := session.dispatch(data); err != nil {
			log.Warnf("con[%d], dispatch error %s", session.fc.ConnID(), err.Error())

			if err == driver.ErrBadConn {
				// TODO handle error
				log.Warn(err)
			}

			return err
		}

		if session.closed {
			// TODO return MySQL Go Away ?
			return nil
		}

		session.fc.ResetSequence()
	}

	return nil
}

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

func (session *Session) ServerName() []byte {
	return hack.Slice(version.Version)
}

func (session *Session) Salt() []byte {
	return session.salt
}
