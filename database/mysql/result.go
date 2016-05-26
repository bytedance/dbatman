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
	warnings     []error
	affectedRows int64
	insertId     int64
	status_info  string
}

func (r *MySQLResult) Status() (int64, error) {
	return int64(r.status), nil
}

func (r *MySQLResult) Warnings() []error {
	return r.warnings
}

func (r *MySQLResult) LastInsertId() (int64, error) {
	return r.insertId, nil
}

func (r *MySQLResult) RowsAffected() (int64, error) {
	return r.affectedRows, nil
}

func (r *MySQLResult) Info() (string, error) {
	return r.status_info, nil
}
