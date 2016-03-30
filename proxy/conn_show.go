package proxy

import (
	. "github.com/bytedance/dbatman/database/mysql"
	"github.com/bytedance/dbatman/hack"
	"github.com/bytedance/dbatman/parser"
)

func (c *Session) handleShow(strsql string, stmt parser.IShow) error {
	var err error

	switch stmt.(type) {
	case *parser.ShowDatabases:
		err = c.handleShowDatabases()
	default:
		err = c.handleQuery(stmt, strsql)
	}

	return err

}

func (session *Session) handleShowDatabases() error {
	dbs := make([]interface{}, 0, 1)
	dbs[0] = session.user.DBName

	if r, err := session.buildSimpleShowResultset(dbs, "Database"); err != nil {
		return err
	} else {
		return c.WriteResult(session.status, r)
	}
}

func (c *Session) buildSimpleShowResultset(values []interface{}, name string) (*MySQLResult, error) {

	r := new(Result)

	field := &Field{}

	field.Name = hack.Slice(name)
	field.Charset = 33
	field.Type = mysql.MYSQL_TYPE_VAR_STRING

	r.Fields = []*Field{field}

	var row []byte
	var err error

	for _, value := range values {
		row, err = formatValue(value)
		if err != nil {
			return nil, err
		}
		r.RowDatas = append(r.RowDatas, PutLengthEncodedString(row))
	}

	return r, nil
}
