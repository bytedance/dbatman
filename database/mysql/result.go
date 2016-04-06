// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2012 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

type MySQLResult struct {
	status       statusFlag
	warnings     uint16
	affectedRows int64
	insertId     int64
}

func (r *MySQLResult) Status() uint16 {
	return uint16(r.status)
}

func (r *MySQLResult) Warnings() uint16 {
	return r.warnings
}

func (res *MySQLResult) LastInsertId() (int64, error) {
	return res.insertId, nil
}

func (res *MySQLResult) RowsAffected() (int64, error) {
	return res.affectedRows, nil
}
