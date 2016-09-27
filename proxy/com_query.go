package proxy

import (
	"errors"
	"fmt"
	"time"

	. "github.com/bytedance/dbatman/database/mysql"
	"github.com/bytedance/dbatman/hack"
	"github.com/bytedance/dbatman/parser"
	"github.com/ngaut/log"
	"github.com/percona/go-mysql/query"
)

//we just go the microsecond timestamp
func getTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (c *Session) comQuery(sqlstmt string) error {

	//TODO accerlate the flow control module and the figerprint module
	// err := c.intercept(sqlstmt)
	// if err != nil {
	// return err
	// }
	// c.updatefp(sqlstmt)
	log.Infof("session %d: %s", c.sessionId, sqlstmt)
	stmt, err := parser.Parse(sqlstmt)
	if err != nil {
		log.Warningf(`parse sql "%s" error "%s"`, sqlstmt, err.Error())
		return c.handleMySQLError(
			NewDefaultError(ER_SYNTAX_ERROR, err.Error()))
	}
	switch v := stmt.(type) {
	case parser.ISelect:
		return c.handleQuery(v, sqlstmt)
	case *parser.Insert, *parser.Update, *parser.Delete, *parser.Replace:
		return c.handleExec(stmt, sqlstmt, false)
	case *parser.Set:
		return c.handleSet(v, sqlstmt)
	case *parser.Begin, *parser.StartTrans:
		return c.handleBegin()
	case *parser.Commit:
		return c.handleCommit()
	case *parser.Rollback:
		// log.Debug(hack.String(stmt.(*parser.Rollback).Point))
		if len(stmt.(*parser.Rollback).Point) > 0 {
			return c.handleExec(stmt, sqlstmt, false)
		}
		return c.handleRollback()
	case parser.IShow:
		return c.handleShow(sqlstmt, v)
	case parser.IDDLStatement:
		return c.handleDDL(v, sqlstmt)
	case *parser.Do, *parser.Call, *parser.FlushTables:
		return c.handleExec(stmt, sqlstmt, false)
		//add the describe table module
	case *parser.DescribeTable, *parser.DescribeStmt:
		return c.handleQuery(v, sqlstmt)
	case *parser.Use:

		if err := c.useDB(hack.String(stmt.(*parser.Use).DB)); err != nil {
			return c.handleMySQLError(err)
		} else {
			return c.fc.WriteOK(nil)
		}
	case *parser.SavePoint:
		return c.handleExec(stmt, sqlstmt, false)
		// return c.handleQuery(v, sqlstmt)
	default:
		log.Warnf("session %d : statement %T[%s] not support now", c.sessionId, stmt, sqlstmt)
		err := errors.New("statement not support now")
		return c.handleMySQLError(
			NewDefaultError(ER_SYNTAX_ERROR, err.Error()))
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

	if err := session.checkDB(stmt); err != nil {
		return session.handleMySQLError(err)
	}

	return session.exec(sqlstmt, isread)
}

// handleDDL process DDL Statements where
func (session *Session) handleDDL(ddl parser.IDDLStatement, sqlstmt string) error {
	if err := session.checkDB(ddl); err != nil {
		return session.handleMySQLError(err)
	}

	// All DDL statement must use master conn
	return session.exec(sqlstmt, false)
}

// for a weak secure issue, we check the db in statement to protect wrong ops
func (session *Session) checkDB(stmt parser.IStatement) error {
	if hasSchemas, ok := stmt.(parser.IDDLSchemas); ok {
		// check schemas to ensure a weak secure issue
		schemas := hasSchemas.GetSchemas()
		for _, s := range schemas {
			if len(s) > 0 && s != session.cluster.DBName {
				log.Warn("wrong here", session.user.Username,
					session.fc.RemoteAddr().String(),
					session.cluster.DBName)
				NewDefaultError(
					ER_DBACCESS_DENIED_ERROR,
					session.user.Username,
					session.fc.RemoteAddr().String(),
					session.cluster.DBName)
			}
		}
	}

	return nil
}

func (session *Session) exec(sqlstmt string, isread bool) error {

	rs, err := session.Executor(isread).Exec(sqlstmt)
	if err != nil {
		return session.handleMySQLError(err)
	}

	return session.fc.WriteOK(rs)
}

func (c *Session) intercept(sqlstmt string) error {
	var excess int64
	now := getTimestamp()
	c.server.mu.Lock()
	if qpsOnServer := c.server.qpsOnServer; qpsOnServer != nil {

		//how many microsecond elapsed since last query
		ms := now - qpsOnServer.last
		//TODO modify here
		//Default, we have 1 r/s and  *1000 add the switch to ms 1000 means 1 req per ms
		excess = qpsOnServer.excess - (c.config.Global.ReqRate*1000*ms)/1000 + 1000
		// if excess < 0 {
		// excess = 0
		// }
		// log.Info("current qps excess is : ", excess)
		//If we need caculate every second speed,
		//Shouldn't reset to zero;

		//the race out the max Burst?
		if excess > c.config.Global.ReqBurst {
			//Just close the client or
			err := fmt.Errorf(`the query excess(%d) over the reqBurst(%d), sql: %s "`, excess, c.config.Global.ReqBurst, sqlstmt)
			log.Warn(err)
			//TODO: more gracefully add a Timer and retry?
		}
		qpsOnServer.excess = excess
		qpsOnServer.last = now
		qpsOnServer.count++
	} else {
		qpsOnServer := &LimitReqNode{}
		qpsOnServer.count = 1
		qpsOnServer.currentcount = 1
		qpsOnServer.last = now
		qpsOnServer.excess = 0
		c.server.qpsOnServer = qpsOnServer

	}

	c.server.mu.Unlock()
	return nil
}

//only collect the qps of the Fingerprint on server
func (c *Session) updatefp(sqlstmt string) {
	now := getTimestamp()
	//*necessary to lock the server

	fp := query.Fingerprint(sqlstmt)

	c.server.mu.Lock()

	if lr, ok := c.server.fingerprints[fp]; ok {
		//how many microsecond elapsed since last query
		interval := now - lr.start

		if interval < 1000 {
			lr.currentcount++
		} else {
			lr.lastcount = lr.currentcount
			lr.start = lr.start + interval/1000*1000
			lr.currentcount = 1
		}
		lr.count++ //total num of printfinger

	} else {
		lr := &LimitReqNode{}
		lr.start = now
		lr.lastcount = 0
		lr.currentcount = 1
		lr.count = 1
		c.server.fingerprints[fp] = lr
	}
	c.server.mu.Unlock()
}
