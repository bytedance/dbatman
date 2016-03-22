// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2012 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

import (
	"github.com/bytedance/dbatman/database/sql/driver"
	"io"
)

type mysqlField struct {
	schema    []byte
	tableName []byte
	orgTable  []byte
	name      []byte
	orgName   []byte
	charset   uint16
	columnLen uint32
	flags     fieldFlag
	fieldType byte
	decimals  byte

	defaultValue       []byte
	defaultValueLength uint64
}

func (f *mysqlField) Dump() []byte {

	l := len(f.schema) + len(f.tableName) + len(f.orgTable) + len(f.name) +
		len(f.orgName) + len(f.defaultValue) + 48

	data := make([]byte, 0, l)

	data = append(data, PutLengthEncodedString([]byte("def"))...)

	data = append(data, PutLengthEncodedString(f.schema)...)

	data = append(data, PutLengthEncodedString(f.tableName)...)
	data = append(data, PutLengthEncodedString(f.orgTable)...)

	data = append(data, PutLengthEncodedString(f.name)...)
	data = append(data, PutLengthEncodedString(f.orgName)...)

	data = append(data, 0x0c)

	data = append(data, Uint16ToBytes(f.charset)...)
	data = append(data, Uint32ToBytes(f.columnLen)...)
	data = append(data, Uint16ToBytes(uint16(f.fieldType))...)
	data = append(data, Uint16ToBytes(uint16(f.flags))...)
	data = append(data, f.decimals)
	data = append(data, 0, 0)

	if f.defaultValue != nil {
		data = append(data, Uint64ToBytes(f.defaultValueLength)...)
		data = append(data, f.defaultValue...)
	}

	return data

}

type mysqlRows struct {
	mc      *mysqlConn
	columns []mysqlField
}

type binaryRows struct {
	mysqlRows
}

type textRows struct {
	mysqlRows
}

type emptyRows struct{}

func (rows *mysqlRows) Columns() []string {
	columns := make([]string, len(rows.columns))
	if rows.mc != nil && rows.mc.cfg.ColumnsWithAlias {
		for i := range columns {
			if tableName := rows.columns[i].tableName; len(tableName) > 0 {
				columns[i] = string(tableName) + "." + string(rows.columns[i].name)
			} else {
				columns[i] = string(rows.columns[i].name)
			}
		}
	} else {
		for i := range columns {
			columns[i] = string(rows.columns[i].name)
		}
	}
	return columns
}

func (rows *mysqlRows) DumpColumns() []driver.RawPacket {
	pkgs := make([]driver.RawPacket, len(rows.columns))

	for i, column := range rows.columns {
		pkgs[i] = driver.RawPacket(column.Dump())
	}

	return pkgs
}

func (rows *mysqlRows) Close() error {
	mc := rows.mc
	if mc == nil {
		return nil
	}
	if mc.netConn == nil {
		return ErrInvalidConn
	}

	// Remove unread packets from stream
	err := mc.readUntilEOF()
	if err == nil {
		if err = mc.discardResults(); err != nil {
			return err
		}
	}

	rows.mc = nil
	return err
}

func (rows *mysqlRows) NextRowPacket() (driver.RawPacket, error) {
	if mc := rows.mc; mc != nil {
		if mc.netConn == nil {
			return nil, ErrInvalidConn
		}

		// Fetch next row from stream
		// dest = rows.readRowPacket(dest)
		return nil, nil
	}
	return nil, io.EOF
}

func (rows *binaryRows) Next(dest []driver.Value) error {
	if mc := rows.mc; mc != nil {
		if mc.netConn == nil {
			return ErrInvalidConn
		}

		// Fetch next row from stream
		return rows.readRow(dest)
	}
	return io.EOF
}

func (rows *textRows) Next(dest []driver.Value) error {
	if mc := rows.mc; mc != nil {
		if mc.netConn == nil {
			return ErrInvalidConn
		}

		// Fetch next row from stream
		return rows.readRow(dest)
	}
	return io.EOF
}

func (rows emptyRows) Columns() []string {
	return nil
}

func (rows emptyRows) Close() error {
	return nil
}

func (rows emptyRows) Next(dest []driver.Value) error {
	return io.EOF
}

func (rows emptyRows) DumpColumns() []driver.RawPacket {
	return nil
}

func (rows emptyRows) NextRowPacket() (driver.RawPacket, error) {
	return nil, nil
}
