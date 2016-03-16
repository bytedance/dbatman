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

import ()

var DEFAULT_CAPABILITY uint32 = mysql.CLIENT_LONG_PASSWORD | mysql.CLIENT_LONG_FLAG |
	mysql.CLIENT_CONNECT_WITH_DB | mysql.CLIENT_PROTOCOL_41 |
	mysql.CLIENT_TRANSACTIONS | mysql.CLIENT_SECURE_CONNECTION

var baseConnId uint32 = 10000

type Context struct {
	s         *Server
	connID    uint32
	status    uint32
	collation uint32
	charset   uint32

	salt []byte
}

func (s *Server) newCtx() *Context {
	c := new(Context)

	c.server = s

	c.connID = atomic.AddUint32(&baseConnId, 1)

	c.status = mysql.SERVER_STATUS_AUTOCOMMIT

	c.salt, _ = mysql.RandomBuf(20)

	c.collation = mysql.DEFAULT_COLLATION_ID
	c.charset = mysql.DEFAULT_CHARSET

	return c
}

func (c *Context) Run() error {

	for {
		data, err := c.front.ReadPacket()
		if err != nil {
			return err
		}

		if err := c.dispatch(data); err != nil {
			if err != mysql.ErrBadConn {
				c.writeError(err)
				return nil
			}

			log.Warning("con[%d], dispatch error %s", c.connID, err.Error())
			return err
		}

		if c.closed {
			return
		}

		c.ResetSequence()
	}

	return nil
}
