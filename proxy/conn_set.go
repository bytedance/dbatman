package proxy

import (
	"fmt"
	. "github.com/bytedance/dbatman/database/mysql"
	"github.com/bytedance/dbatman/parser"
	"github.com/ngaut/log"
	"strings"
)

func (c *Session) handleSet(stmt *parser.Set, sql string) error {
	if len(stmt.VarList) < 1 {
		return fmt.Errorf("must set one item at least")
	}

	var err error
	for _, v := range stmt.VarList {
		if strings.ToUpper(v.Name) == "AUTOCOMMIT" {
			log.Debug("handle autocommit")
			err = c.handleSetAutoCommit(v.Value)
		}
	}

	if err != nil {
		return err
	}
	return c.handleOtherSet(stmt, sql)
}

func (c *Session) handleSetAutoCommit(val parser.IExpr) error {

	var stmt *parser.Predicate
	var ok bool
	if stmt, ok = val.(*parser.Predicate); !ok {
		return fmt.Errorf("set autocommit is not support for complicate expressions")
	}

	switch value := stmt.Expr.(type) {
	case parser.NumVal:
		if i, err := value.ParseInt(); err != nil {
			return err
		} else if i == 1 {
			c.status |= uint32(StatusInAutocommit)
			log.Debug("autocommit is set")
		} else if i == 0 {
			c.status &= ^uint32(StatusInAutocommit)
			log.Debug("auto commit is unset")
		} else {
			return fmt.Errorf("Variable 'autocommit' can't be set to the value of '%s'", i)
		}
	case parser.StrVal:
		if s := value.Trim(); s == "" {
			return fmt.Errorf("Variable 'autocommit' can't be set to the value of ''")
		} else if us := strings.ToUpper(s); us == `ON` {
			c.status |= uint32(StatusInAutocommit)
			log.Debug("auto commit is set")
		} else if us == `OFF` {
			c.status &= ^uint32(StatusInAutocommit)
			log.Debug("auto commit is unset")
		} else {
			return fmt.Errorf("Variable 'autocommit' can't be set to the value of '%s'", us)
		}
	default:
		return fmt.Errorf("set autocommit error, value type is %T", val)
	}

	return nil
}

func (c *Session) handleSetNames(val parser.IValExpr) error {
	value, ok := val.(parser.StrVal)
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

	return c.fc.WriteOK(nil)
}

func (c *Session) handleOtherSet(stmt parser.IStatement, sql string) error {
	return c.handleExec(stmt, sql, false)
}
