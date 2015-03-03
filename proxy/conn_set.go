package proxy

import (
	"fmt"
	. "github.com/wangjild/go-mysql-proxy/log"
	. "github.com/wangjild/go-mysql-proxy/mysql"
	"github.com/wangjild/go-mysql-proxy/sql"
	"strings"
)

func (c *Conn) handleSet(stmt *sql.Set, sql string) error {
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

func (c *Conn) handleSetAutoCommit(val sql.IExpr) error {

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

func (c *Conn) handleSetNames(val sql.IValExpr) error {
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

func (c *Conn) handleOtherSet(stmt sql.IStatement, sql string) error {
	return c.handleExec(stmt, sql, false)
}
