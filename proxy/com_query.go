package proxy

import (
	"fmt"
	. "github.com/bytedance/dbatman/database/mysql"
	"github.com/bytedance/dbatman/database/sql"
	"github.com/bytedance/dbatman/errors"
	"github.com/bytedance/dbatman/hack"
	"github.com/bytedance/dbatman/parser"
	"github.com/ngaut/log"
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
		return c.handleDDL(v, sqlstmt)
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

func makeBindVars(args []interface{}) map[string]interface{} {
	bindVars := make(map[string]interface{}, len(args))

	for i, v := range args {
		bindVars[fmt.Sprintf("v%d", i+1)] = v
	}

	return bindVars
}

func (session *Session) handleExec(stmt parser.IStatement, sqlstmt string, isread bool) error {

	log.Debug("handle exec", sqlstmt)

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
		if mysql_rs, ok := rs.(*MySQLResult); ok {
			err = session.fc.WriteOK(mysql_rs)
		}
	}

	return err
}

// handleDDL process DDL Statements where
func (session *Session) handleDDL(ddl parser.IDDLStatement, sqlstmt string) error {
	if hasSchemas, ok := ddl.(parser.IDDLSchemas); ok {
		// check schemas to ensure a weak secure issue
		schemas := hasSchemas.GetSchemas()
		for _, s := range schemas {
			if len(s) > 0 && s != session.cluster.DBName {
				return session.handleMySQLError(
					NewDefaultError(
						ER_DBACCESS_DENIED_ERROR,
						session.user.Username,
						session.fc.RemoteAddr().String(),
						session.cluster.DBName))
			}
		}
	}

	// All DDL statement must use master conn
	return errors.Trace(session.exec(sqlstmt, false))
}

func (session *Session) exec(sqlstmt string, isread bool) error {

	db, err := session.cluster.DB(isread)
	if err != nil {
		return errors.Trace(err)
	}

	defer db.Close()

	var rs sql.Result
	rs, err = db.Exec(sqlstmt)

	if err != nil {
		return errors.Trace(session.handleMySQLError(err))
	}

	return errors.Trace(session.fc.WriteOK(rs))
}
