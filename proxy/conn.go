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
	"fmt"
	"github.com/bytedance/dbatman/Godeps/_workspace/src/github.com/ngaut/log"
	"github.com/bytedance/dbatman/database/mysql"
	"github.com/bytedance/dbatman/hack"
)

func (ctx *Context) Close() error {
	if ctx.closed {
		return nil
	}

	ctx.conn.Close()

	ctx.rollback()
	for _, s := range c.stmts {
		s.Close()
	}

	c.stmts = nil

	c.closed = true

	return nil
}

func (c *Context) dispatch(data []byte) error {
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

func (c *Context) useDB(db string) error {
	if s := c.server.getSchema(db); s == nil {
		return mysql.NewDefaultError(mysql.ER_BAD_DB_ERROR, db)
	} else {
		c.schema = s
		c.db = db
	}
	return nil
}

func (c *Context) IsAutoCommit() bool {
	return c.status&mysql.SERVER_STATUS_AUTOCOMMIT > 0
}

func (c *Context) checkDB() error {
	if c.schema != nil {
		return nil
	}

	if c.db != "" {
		return c.useDB(c.db)
	}

	return mysql.NewDefaultError(mysql.ER_NO_DB_ERROR)
}
