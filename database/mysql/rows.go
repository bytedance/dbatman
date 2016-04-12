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
	"github.com/bytedance/dbatman/errors"
	"io"
)

type MySQLField struct {
	Catalog   []byte
	Database  []byte
	Table     []byte
	OrgTable  []byte
	Name      []byte
	OrgName   []byte
	Charset   uint16
	Length    uint32
	Flags     fieldFlag
	FieldType byte
	Decimals  byte

	DefaultValue       []byte
	DefaultValueLength uint64
}

func (f *MySQLField) Dump() []byte {

	l := len(f.Database) + len(f.Table) + len(f.OrgTable) + len(f.Name) +
		len(f.OrgName) + len(f.DefaultValue) + 52

	data := make([]byte, 4, l)

	data = appendLengthEncodedString(data, f.Catalog)
	data = appendLengthEncodedString(data, f.Database)

	data = appendLengthEncodedString(data, f.Table)
	data = appendLengthEncodedString(data, f.OrgTable)

	data = appendLengthEncodedString(data, f.Name)
	data = appendLengthEncodedString(data, f.OrgName)

	// Filler always be 0x0c
	data = append(data, 0x0c)

	data = append(data, Uint16ToBytes(f.Charset)...)
	data = append(data, Uint32ToBytes(f.Length)...)
	data = append(data, Uint16ToBytes(uint16(f.FieldType))...)
	data = append(data, Uint16ToBytes(uint16(f.Flags))...)
	data = append(data, f.Decimals)

	// Filler always be 2 bytes 0
	data = append(data, 0, 0)

	if f.DefaultValue != nil {
		data = append(data, uint64ToBytes(f.DefaultValueLength)...)
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
			if tableName := rows.columns[i].Table; len(tableName) > 0 {
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

func (rows *MySQLRows) DumpColumns() []driver.RawPacket {
	pkgs := make([]driver.RawPacket, len(rows.columns))

	for i, column := range rows.columns {
		pkgs[i] = driver.RawPacket(column.Dump())
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

func (rows *BinaryRows) NextRowPacket() (driver.RawPacket, error) {
	if mc := rows.mc; mc != nil {
		if mc.netConn == nil {
			return nil, errors.Trace(ErrInvalidConn)
		}

		// Fetch next row from stream
		return rows.readRowPacket()
	}
	return nil, errors.Trace(io.EOF)
}

func (rows *BinaryRows) readRowPacket() (driver.RawPacket, error) {
	data, err := rows.mc.readPacket()
	if err != nil {
		return nil, errors.Trace(err)
	}

	// packet indicator [1 byte]
	if data[0] != iOK {
		// EOF Packet
		if data[0] == iEOF && len(data) == 5 {
			rows.mc.status = readStatus(data[3:])
			if err := rows.mc.discardResults(); err != nil {
				return nil, errors.Trace(err)
			}
			rows.mc = nil
			return nil, errors.Trace(io.EOF)
		}
		rows.mc = nil

		// Error otherwise
		return nil, errors.Trace(rows.mc.handleErrorPacket(data))
	}

	return data, nil
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

func (rows *TextRows) NextRowPacket() (driver.RawPacket, error) {
	if mc := rows.mc; mc != nil {
		if mc.netConn == nil {
			return nil, errors.Trace(ErrInvalidConn)
		}

		// Fetch next row from stream
		return rows.readRowPacket()
	}
	return nil, errors.Trace(io.EOF)
}

func (rows *TextRows) readRowPacket() (driver.RawPacket, error) {
	data, err := rows.mc.readPacket()
	if err != nil {
		return nil, errors.Trace(err)
	}

	// EOF Packet
	if data[0] == iEOF && len(data) == 5 {
		// server_status [2 bytes]
		rows.mc.status = readStatus(data[3:])
		if err := rows.mc.discardResults(); err != nil {
			return nil, errors.Trace(err)
		}
		rows.mc = nil
		return nil, errors.Trace(io.EOF)
	}
	if data[0] == iERR {
		rows.mc = nil
		return nil, errors.Trace(rows.mc.handleErrorPacket(data))
	}

	// Preserve packet header for proxy usage
	return append(make([]byte, PacketHeaderLen, len(data)+PacketHeaderLen), data...), nil
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
