package proxy

import (
	"bytes"

	"github.com/bytedance/dbatman/database/mysql"
	"github.com/bytedance/dbatman/hack"
	"github.com/bytedance/dbatman/parser"
	"github.com/ngaut/log"
)

func (session *Session) handleShow(sqlstmt string, stmt parser.IShow) error {
	var err error

	switch stmt.(type) {
	case *parser.ShowDatabases:
		err = session.handleShowDatabases()
	default:
		err = session.handleQuery(stmt, sqlstmt)
	}

	if err != nil {
		return session.handleMySQLError(err)
	}

	return nil
}

func (session *Session) handleFieldList(data []byte) error {
	index := bytes.IndexByte(data, 0x00)
	table := string(data[0:index])
	wildcard := string(data[index+1:])

	rs, err := session.bc.master.FieldList(table, wildcard)
	// TODO here should handler error
	if err != nil {
		return session.handleMySQLError(err)
	}

	defer rs.Close()

	return session.writeFieldList(rs)
}

func (session *Session) writeFieldList(rs mysql.Rows) error {

	cols, err := rs.ColumnPackets()

	if err != nil {
		return session.handleMySQLError(err)
	}

	// Write Columns Packet
	for _, col := range cols {
		if err := session.fc.WritePacket(col); err != nil {
			log.Debugf("write columns packet error %v", err)
			return err
		}
	}

	// TODO Write a ok packet
	if err = session.fc.WriteEOF(); err != nil {
		return err
	}

	return nil
}

func (session *Session) handleShowDatabases() error {
	dbs := make([]interface{}, 0, 1)
	dbs = append(dbs, session.user.DBName)

	if r, err := session.buildSimpleShowResultset(dbs, "Database"); err != nil {
		return err
	} else {
		return session.writeRows(r)
	}
}

func (session *Session) buildSimpleShowResultset(values []interface{}, name string) (mysql.Rows, error) {

	r := new(SimpleRows)

	r.Cols = []*mysql.MySQLField{
		&mysql.MySQLField{
			Name:      hack.Slice(name),
			Charset:   uint16(session.fc.Collation()),
			FieldType: mysql.FieldTypeVarString,
		},
	}

	var row []byte
	var err error

	for _, value := range values {
		row, err = formatValue(value)
		if err != nil {
			return nil, err
		}

		r.Rows = append(r.Rows, mysql.AppendLengthEncodedString(make([]byte, 0, len(row)+9), row))
	}

	return r, nil
}
