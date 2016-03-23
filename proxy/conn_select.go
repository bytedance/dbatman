package proxy

import (
	"bytes"
	"fmt"
	"github.com/bytedance/dbatman/database/mysql"
	"github.com/bytedance/dbatman/database/sql/driver"
	"github.com/bytedance/dbatman/parser"
)

func (c *Session) buildSimpleSelectResult(value interface{}, name []byte, asName []byte) (*Resultset, error) {
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

func (c *Session) handleFieldList(data []byte) error {
	index := bytes.IndexByte(data, 0x00)
	table := string(data[0:index])
	wildcard := string(data[index+1:])

	if c.schema == nil {
		return NewDefaultError(mysql.ER_NO_DB_ERROR)
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

func (c *Session) writeFieldList(status uint16, fs []*Field) error {
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

func (session *Session) handleSelect(stmt parser.IStatement, sqlstmt string) error {

	if err := session.checkDB(); err != nil {
		return err
	}

	isread := false
	if s, ok := stmt.(parser.ISelect); ok {
		isread = !s.IsLocked()
	} else if _, sok := stmt.(parser.IShow); sok {
		isread = true
	}

	db, err := session.cluster.DB(isread)

	// TODO here if db is nil, then we should return a error?
	if err != nil {
		return err
	} else if db == nil {
		// r := c.newEmptyResultset(stmt)
		// return c.writeResultset(c.status, r)
		return fmt.Errorf("no available backend db")
	}

	var rs *driver.rows
	rs, err = db.Query(sqlstmt)

	// TODO here should handler error
	if err != nil {
		return err
	}

	defer rs.Close()

	if err := session.fc.WritePacket(res.DumpColumns()...); err != nil {
		return err
	}

	var payload driver.RawPayload
	for {
		payload, err := rs.NextRowPayload()
		if err != nil {
			if merr, ok := err.(*MySQLError); ok {
				session.fc.WriteError(merr)
			}
			return err
		}

		if err := session.fc.WritePacket(payload); err != nil {
			return err
		}
	}

	return nil
}

// TODO
func (c *Session) doQuery(sqlstmt string) error {
	return nil
}
