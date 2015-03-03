package proxy

import (
	"github.com/wangjild/go-mysql-proxy/hack"
	. "github.com/wangjild/go-mysql-proxy/mysql"
	"github.com/wangjild/go-mysql-proxy/sql"
)

func (c *Conn) handleShow(strsql string, stmt sql.IShow) error {
	var err error

	switch stmt.(type) {
	case *sql.ShowDatabases:
		err = c.handleShowDatabases()
	default:
		err = c.handleSelect(stmt, strsql)
	}

	return err

}

func (c *Conn) handleShowDatabases() error {
	dbs := make([]interface{}, 0, len(c.server.schemas))
	for key := range c.server.schemas {
		dbs = append(dbs, key)
	}

	if r, err := c.buildSimpleShowResultset(dbs, "Database"); err != nil {
		return err
	} else {
		return c.writeResultset(c.status, r)
	}
}

func (c *Conn) buildSimpleShowResultset(values []interface{}, name string) (*Resultset, error) {

	r := new(Resultset)

	field := &Field{}

	field.Name = hack.Slice(name)
	field.Charset = 33
	field.Type = MYSQL_TYPE_VAR_STRING

	r.Fields = []*Field{field}

	var row []byte
	var err error

	for _, value := range values {
		row, err = formatValue(value)
		if err != nil {
			return nil, err
		}
		r.RowDatas = append(r.RowDatas,
			PutLengthEncodedString(row))
	}

	return r, nil
}
