package proxy

import (
	"bytes"
	"fmt"
	. "github.com/wangjild/go-mysql-proxy/mysql"
	"github.com/wangjild/go-mysql-proxy/sql"
)

func (c *Conn) buildSimpleSelectResult(value interface{}, name []byte, asName []byte) (*Resultset, error) {
	field := &Field{}

	field.Name = name

	if asName != nil {
		field.Name = asName
	}

	field.OrgName = name

	formatField(field, value)

	r := &Resultset{Fields: []*Field{field}}
	row, err := formatValue(value)
	if err != nil {
		return nil, err
	}
	r.RowDatas = append(r.RowDatas, PutLengthEncodedString(row))

	return r, nil
}

func (c *Conn) handleFieldList(data []byte) error {
	index := bytes.IndexByte(data, 0x00)
	table := string(data[0:index])
	wildcard := string(data[index+1:])

	if c.schema == nil {
		return NewDefaultError(ER_NO_DB_ERROR)
	}

	co, err := c.schema.node.getMasterConn()
	if err != nil {
		return err
	}
	defer co.Close()

	if err = co.UseDB(c.schema.db); err != nil {
		return err
	}

	if fs, err := co.FieldList(table, wildcard); err != nil {
		return err
	} else {
		return c.writeFieldList(c.status, fs)
	}
}

func (c *Conn) writeFieldList(status uint16, fs []*Field) error {
	c.affectedRows = int64(-1)

	data := make([]byte, 4, 1024)

	for _, v := range fs {
		data = data[0:4]
		data = append(data, v.Dump()...)
		if err := c.writePacket(data); err != nil {
			return err
		}
	}

	if err := c.writeEOF(status); err != nil {
		return err
	}
	return nil
}

func (c *Conn) handleSelect(stmt sql.IStatement, sqlstmt string) error {

	if err := c.checkDB(); err != nil {
		return err
	}

	isread := false
	if s, ok := stmt.(sql.ISelect); ok {
		isread = !s.IsLocked()
	} else if _, sok := stmt.(sql.IShow); sok {
		isread = true
	}

	conn, err := c.getConn(c.schema.node, isread)

	if err != nil {
		return err
	} else if conn == nil {
		// r := c.newEmptyResultset(stmt)
		// return c.writeResultset(c.status, r)
		return fmt.Errorf("no available connection")
	}

	var res *Result
	res, err = conn.Execute(sqlstmt)

	c.closeDBConn(conn, false)

	if err == nil {
		err = c.mergeSelectResult(res)
	}

	return err
}
