package proxy

import (
	"errors"
	. "github.com/bytedance/dbatman/database/mysql"
	"github.com/bytedance/dbatman/database/sql/driver"
	"io"
)

// SimpleRows implements sql.Row
type SimpleRows struct {
	Columns []*MySQLField
	Rows    []driver.RawPacket
	rowsi   int
}

//	Next() bool
//	NextRowPacket() (driver.RawPacket, error)
//	ColumnPackets() ([]driver.RawPacket, error)
//	Scan(dest ...interface{}) error
//	Close() error
//	Err() error

func (rs *SimpleRows) Next() bool {
	return rs.rowsi < len(rs.Rows)
}

func (rs *SimpleRows) NextRowPacket() (driver.RawPacket, error) {
	if rs.rowsi >= len(rs.Rows) {
		return nil, io.EOF
	}

	ret := make([]byte, PacketHeaderLen, len(rs.Rows[rs.rowsi])+PacketHeaderLen)
	ret = append(ret, rs.Rows[rs.rowsi]...)
	rs.rowsi += 1
	return ret, nil
}

func (rs *SimpleRows) ColumnPackets() ([]driver.RawPacket, error) {
	pkgs := make([]driver.RawPacket, len(rs.Columns))

	for i, column := range rs.Columns {
		pkgs[i] = driver.RawPacket(column.Dump())
	}

	return pkgs, nil
}

func (rs *SimpleRows) Scan(dest ...interface{}) error {
	return errors.New("SimpleRows does not support scan operations")
}

func (rs *SimpleRows) Close() error {
	return nil
}

func (rs *SimpleRows) Err() error {
	return nil
}
