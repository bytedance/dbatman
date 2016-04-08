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

type MySQLField struct {
	Schema    []byte
	TableName []byte
	OrgTable  []byte
	Name      []byte
	OrgName   []byte
	Charset   uint16
	ColumnLen uint32
	Flags     fieldFlag
	FieldType byte
	Decimals  byte

	DefaultValue       []byte
	DefaultValueLength uint64
}

func (f *MySQLField) Dump() []byte {

	l := len(f.Schema) + len(f.TableName) + len(f.OrgTable) + len(f.Name) +
		len(f.OrgName) + len(f.DefaultValue) + 48

	data := make([]byte, 0, l)

	data = append(data, PutLengthEncodedString([]byte("def"))...)

	data = append(data, PutLengthEncodedString(f.Schema)...)

	data = append(data, PutLengthEncodedString(f.TableName)...)
	data = append(data, PutLengthEncodedString(f.OrgTable)...)

	data = append(data, PutLengthEncodedString(f.Name)...)
	data = append(data, PutLengthEncodedString(f.OrgName)...)

	data = append(data, 0x0c)

	data = append(data, Uint16ToBytes(f.Charset)...)
	data = append(data, Uint32ToBytes(f.ColumnLen)...)
	data = append(data, Uint16ToBytes(uint16(f.FieldType))...)
	data = append(data, Uint16ToBytes(uint16(f.Flags))...)
	data = append(data, f.Decimals)
	data = append(data, 0, 0)

	if f.DefaultValue != nil {
		data = append(data, Uint64ToBytes(f.DefaultValueLength)...)
		data = append(data, f.DefaultValue...)
	}

	return data

}

type MySQLRows struct {
	mc      *MySQLConn
	columns []MySQLField
}

type BinaryRows struct {
	MySQLRows
}

type TextRows struct {
	MySQLRows
}

type emptyRows struct{}

func (rows *MySQLRows) Columns() []string {
	columns := make([]string, len(rows.columns))
	if rows.mc != nil && rows.mc.cfg.ColumnsWithAlias {
		for i := range columns {
			if tableName := rows.columns[i].TableName; len(tableName) > 0 {
				columns[i] = string(tableName) + "." + string(rows.columns[i].Name)
			} else {
				columns[i] = string(rows.columns[i].Name)
			}
		}
	} else {
		for i := range columns {
			columns[i] = string(rows.columns[i].Name)
		}
	}
	return columns
}

func (rows *MySQLRows) DumpColumns() []driver.RawPayload {
	pkgs := make([]driver.RawPayload, len(rows.columns))

	for i, column := range rows.columns {
		pkgs[i] = driver.RawPayload(column.Dump())
	}

	return pkgs
}

func (rows *MySQLRows) Close() error {
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

func (rows *MySQLRows) NextRowPayload() (driver.RawPayload, error) {
	if mc := rows.mc; mc != nil {
		if mc.netConn == nil {
			return nil, ErrInvalidConn
		}

		// Fetch next row from stream
		return rows.readRowPayload()
	}
	return nil, io.EOF
}

func (rows *BinaryRows) Next(dest []driver.Value) error {
	if mc := rows.mc; mc != nil {
		if mc.netConn == nil {
			return ErrInvalidConn
		}

		// Fetch next row from stream
		return rows.readRow(dest)
	}
	return io.EOF
}

func (rows *TextRows) Next(dest []driver.Value) error {
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

func (rows emptyRows) DumpColumns() []driver.RawPayload {
	return nil
}

func (rows emptyRows) NextRowPayload() (driver.RawPayload, error) {
	return nil, nil
}
