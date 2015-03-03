package proxy

import (
	"fmt"
	"github.com/wangjild/go-mysql-proxy/client"
	"github.com/wangjild/go-mysql-proxy/hack"
	. "github.com/wangjild/go-mysql-proxy/mysql"
	"github.com/wangjild/go-mysql-proxy/sql"
)

func (c *Conn) handleQuery(sqlstmt string) (err error) {
	/*defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("execute %s error %v", sql, e)
			return
		}
	}()*/

	var stmt sql.IStatement
	stmt, err = sql.Parse(sqlstmt)
	if err != nil {
		return fmt.Errorf(`parse sql "%s" error "%s"`, sqlstmt, err.Error())
	}

	switch v := stmt.(type) {
	case sql.ISelect:
		return c.handleSelect(v, sqlstmt)
	case *sql.Insert:
		return c.handleExec(stmt, sqlstmt, false)
	case *sql.Update:
		return c.handleExec(stmt, sqlstmt, false)
	case *sql.Delete:
		return c.handleExec(stmt, sqlstmt, false)
	case *sql.Replace:
		return c.handleExec(stmt, sqlstmt, false)
	case *sql.Set:
		return c.handleSet(v, sqlstmt)
	case *sql.Begin:
		return c.handleBegin()
	case *sql.Commit:
		return c.handleCommit()
	case *sql.Rollback:
		return c.handleRollback()
	case sql.IShow:
		return c.handleShow(sqlstmt, v)
	case sql.IDDLStatement:
		return c.handleExec(stmt, sqlstmt, false)
	case *sql.Do:
		return c.handleExec(stmt, sqlstmt, false)
	case *sql.Call:
		return c.handleExec(stmt, sqlstmt, false)
	case *sql.Use:
		if err := c.useDB(hack.String(stmt.(*sql.Use).DB)); err != nil {
			return err
		} else {
			return c.writeOK(nil)
		}

	default:
		return fmt.Errorf("statement %T[%s] not support now", stmt, sqlstmt)
	}

	return nil
}

func (c *Conn) getConn(n *Node, isSelect bool) (co *client.SqlConn, err error) {
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

func (c *Conn) closeDBConn(co *client.SqlConn, rollback bool) {
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

func makeBindVars(args []interface{}) map[string]interface{} {
	bindVars := make(map[string]interface{}, len(args))

	for i, v := range args {
		bindVars[fmt.Sprintf("v%d", i+1)] = v
	}

	return bindVars
}

func (c *Conn) handleExec(stmt sql.IStatement, sqlstmt string, isread bool) error {

	if err := c.checkDB(); err != nil {
		return err
	}

	conn, err := c.getConn(c.schema.node, isread)
	if err != nil {
		return err
	} else if conn == nil {
		return fmt.Errorf("no available connection")
	}

	var rs *Result
	rs, err = conn.Execute(sqlstmt)

	c.closeDBConn(conn, err != nil)

	if err == nil {
		err = c.writeOK(rs)
	}

	return err
}

func (c *Conn) mergeSelectResult(rs *Result) error {
	r := rs.Resultset
	status := c.status | rs.Status
	return c.writeResultset(status, r)
}
