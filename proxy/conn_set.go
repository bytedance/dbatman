package proxy

import (
	"fmt"
	. "github.com/bytedance/dbatman/database/sql/driver/mysql"
	. "github.com/bytedance/dbatman/log"
	"github.com/bytedance/dbatman/sql"
	"strings"
)

func (c *frontConn) handleSet(stmt *sql.Set, sql string) error {
	if len(stmt.VarList) < 1 {
		return fmt.Errorf("must set one item at least")
	}

	var err error
	for _, v := range stmt.VarList {
		if strings.ToUpper(v.Name) == "AUTOCOMMIT" {
			AppLog.Debug("handle autocommit")
			err = c.handleSetAutoCommit(v.Value)
		}
	}

	if err != nil {
		return err
	}
	return c.handleOtherSet(stmt, sql)
}

func (c *frontConn) handleSetAutoCommit(val sql.IExpr) error {

	var stmt *sql.Predicate
	var ok bool
	if stmt, ok = val.(*sql.Predicate); !ok {
		return fmt.Errorf("set autocommit is not support for complicate expressions")
	}

	switch value := stmt.Expr.(type) {
	case sql.NumVal:
		if i, err := value.ParseInt(); err != nil {
			return err
		} else if i == 1 {
			c.status |= SERVER_STATUS_AUTOCOMMIT
			AppLog.Debug("autocommit is set")
		} else if i == 0 {
			c.status &= ^SERVER_STATUS_AUTOCOMMIT
			AppLog.Debug("auto commit is unset")
		} else {
			return fmt.Errorf("Variable 'autocommit' can't be set to the value of '%s'", i)
		}
	case sql.StrVal:
		if s := value.Trim(); s == "" {
			return fmt.Errorf("Variable 'autocommit' can't be set to the value of ''")
		} else if us := strings.ToUpper(s); us == `ON` {
			c.status |= SERVER_STATUS_AUTOCOMMIT
			AppLog.Debug("auto commit is set")
		} else if us == `OFF` {
			c.status &= ^SERVER_STATUS_AUTOCOMMIT
			AppLog.Debug("auto commit is unset")
		} else {
			return fmt.Errorf("Variable 'autocommit' can't be set to the value of '%s'", us)
		}
	default:
		return fmt.Errorf("set autocommit error, value type is %T", val)
	}

	return nil
}

func (c *frontConn) handleSetNames(val sql.IValExpr) error {
	value, ok := val.(sql.StrVal)
	if !ok {
		return fmt.Errorf("set names charset error")
	}

	charset := strings.ToLower(string(value))
	cid, ok := CharsetIds[charset]
	if !ok {
		return fmt.Errorf("invalid charset %s", charset)
	}

	c.charset = charset
	c.collation = cid

	return c.writeOK(nil)
}

func (c *frontConn) handleOtherSet(stmt sql.IStatement, sql string) error {
	return c.handleExec(stmt, sql, false)
}
