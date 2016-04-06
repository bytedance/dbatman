package proxy

import (
	"fmt"
	"github.com/bytedance/dbatman/database/mysql"
	"github.com/bytedance/dbatman/database/sql"
	"github.com/bytedance/dbatman/hack"
	"github.com/bytedance/dbatman/parser"
)

func (c *Session) comQuery(sqlstmt string) (err error) {

	var stmt parser.IStatement
	stmt, err = parser.Parse(sqlstmt)
	if err != nil {
		return fmt.Errorf(`parse sql "%s" error "%s"`, sqlstmt, err.Error())
	}

	switch v := stmt.(type) {
	case parser.ISelect:
		return c.handleQuery(v, sqlstmt)
	case *parser.Insert:
		return c.handleExec(stmt, sqlstmt, false)
	case *parser.Update:
		return c.handleExec(stmt, sqlstmt, false)
	case *parser.Delete:
		return c.handleExec(stmt, sqlstmt, false)
	case *parser.Replace:
		return c.handleExec(stmt, sqlstmt, false)
	case *parser.Set:
		return c.handleSet(v, sqlstmt)
	case *parser.Begin:
		// return c.handleBegin()
		return nil
	case *parser.Commit:
		// return c.handleCommit()
		return nil
	case *parser.Rollback:
		// return c.handleRollback()
		return nil
	case parser.IShow:
		return c.handleShow(sqlstmt, v)
	case parser.IDDLStatement:
		return c.handleExec(stmt, sqlstmt, false)
	case *parser.Do:
		return c.handleExec(stmt, sqlstmt, false)
	case *parser.Call:
		return c.handleExec(stmt, sqlstmt, false)
	case *parser.Use:
		if err := c.useDB(hack.String(stmt.(*parser.Use).DB)); err != nil {
			return err
		} else {
			return c.fc.WriteOK(nil)
		}

	default:
		return fmt.Errorf("statement %T[%s] not support now", stmt, sqlstmt)
	}

	return nil
}

/*
func (c *Session) getConn(n *Node, isSelect bool) (co *backend.SqlConn, err error) {
	if !c.needBeginTx() {
		if isSelect {
			co, err = n.getSelectConn()
		} else {
			co, err = n.getMasterConn()
		}
		if err != nil {
			return
		}
	} else {
		var ok bool
		c.Lock()
		co, ok = c.txConns[n]
		c.Unlock()

		if !ok {
			if co, err = n.getMasterConn(); err != nil {
				return
			}

			if err = co.SetAutocommit(c.IsAutoCommit()); err != nil {
				return
			}

			if err = co.Begin(); err != nil {
				return
			}

			c.Lock()
			c.txConns[n] = co
			c.Unlock()
		}
	}

	//todo, set conn charset, etc...
	if err = co.UseDB(c.schema.db); err != nil {
		return
	}

	if err = co.SetCharset(c.charset); err != nil {
		return
	}

	return
}

func (c *Session) closeDBConn(co *backend.SqlConn, rollback bool) {
	// since we have DDL, and when server is not in autoCommit,
	// we do not release the connection and will reuse it later
	if c.isInTransaction() || !c.isAutoCommit() {
		return
	}

	if rollback {
		co.Rollback()
	}

	co.Close()
}
*/

func makeBindVars(args []interface{}) map[string]interface{} {
	bindVars := make(map[string]interface{}, len(args))

	for i, v := range args {
		bindVars[fmt.Sprintf("v%d", i+1)] = v
	}

	return bindVars
}

func (session *Session) handleExec(stmt parser.IStatement, sqlstmt string, isread bool) error {

	if err := session.checkDB(); err != nil {
		return err
	}

	db, err := session.cluster.DB(isread)

	if err != nil {
		return err
	}

	defer db.Close()

	var rs sql.Result
	rs, err = db.Exec(sqlstmt)

	if err == nil {
		if mysql_rs, ok := rs.(*mysql.MySQLResult); ok {
			err = session.fc.WriteOK(mysql_rs)
		}
	}

	return err
}
