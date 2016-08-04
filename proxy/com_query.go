package proxy

import (
	"errors"
	"fmt"
	"time"

	. "github.com/bytedance/dbatman/database/mysql"
	"github.com/bytedance/dbatman/hack"
	"github.com/bytedance/dbatman/parser"
	"github.com/ngaut/log"
)

//we just go the microsecond timestamp
func getTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (c *Session) comQuery(sqlstmt string) error {
	//calculate the ip
	// var excess int64

	// fp := query.Fingerprint(sqlstmt)
	// now := getTimestamp()
	// c.server.mu.Lock()
	// server := c.server
	// ip := "127.0.0.1"

	// if user, ok := server.users[c.user.Username]; ok {
	// 	if ipinfo, ok := user.iplist[ip]; ok {

	// 		if lr, ok := ipinfo.printfinger[fp]; ok {
	// 			//print finger exist calcute the qps
	// 			tmp := lr
	// 			log.Debug(tmp)

	// 		} else {
	// 			lr := &LimitReqNode{}
	// 			lr.query = fp
	// 			lr.count = 1
	// 			lr.excess = 0
	// 			lr.last = getTimestamp()
	// 			ipinfo.printfinger[fp] = lr
	// 		}
	// 	} else {
	// 		ipinfo := &Ip{}
	// 		ipinfo.ip = ip
	// 		ipinfo.printfinger = make(map[string]*LimitReqNode)

	// 	}
	// } else {
	// 	user := &User{}
	// 	user.user = c.user.Username
	// 	user.iplist = make(map[string]*Ip)

	// }
	// if lr, ok := c.server.fingerprints[fp]; ok {
	// 	//how many microsecond elapsed since last query
	// 	ms := now - lr.last
	// 	//Default, we have 1 r/s
	// 	excess = lr.excess - 1000*(ms/1000) + 1000

	// 	//If we need caculate every second speed,
	// 	//Shouldn't reset to zero;
	// 	if excess < 0 {
	// 		excess = 0
	// 	}
	// 	//the race out the max Burst?
	// 	if excess > 1000 {
	// 		//Just close the client or
	// 		return fmt.Errorf(`the query excess(%d) over the reqBurst(%d), sql: %s "`, excess, 1000, sqlstmt)

	// 		//TODO: more gracefully add a Timer and retry?
	// 	}
	// 	lr.excess = excess
	// 	lr.last = now
	// 	lr.count++

	// } else {
	// 	lr := &LimitReqNode{}
	// 	lr.excess = 0
	// 	lr.last = getTimestamp()
	// 	lr.query = fp

	// 	lr.count = 1
	// 	c.server.fingerprints[fp] = lr
	// }
	// c.server.mu.Unlock()

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
		return c.handleRollback()
	case parser.IShow:
		return c.handleShow(sqlstmt, v)
	case parser.IDDLStatement:
		return c.handleDDL(v, sqlstmt)
	case *parser.Do, *parser.Call, *parser.FlushTables:
		return c.handleExec(stmt, sqlstmt, false)
	case *parser.Use:
		if err := c.useDB(hack.String(stmt.(*parser.Use).DB)); err != nil {
			return c.handleMySQLError(err)
		} else {
			return c.fc.WriteOK(nil)
		}
	default:
		log.Warnf("statement %T[%s] not support now", stmt, sqlstmt)
		// err := log.Error("statement  not support now")
		// return nil
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

func (c *Session) collectfp(fp string) {
	now := getTimestamp()
	//*necessary to lock the server
	// c.server.mu.Lock()
	server := c.server
	ip := "127.0.0.1"

	if user, ok := server.users[c.user.Username]; ok {
		if ipinfo, ok := user.iplist[ip]; ok {
			ipinfo.mu.Unlock()

			if lr, ok := ipinfo.printfinger[fp]; ok {
				//print finger exist calcute the qps

				//calculate the interval
				interval := now - lr.last

				//if interval <1000ms ADD the reqst to the previous 1s period
				//if interval >1000ms begin a new period to record the count
				lr.count += 1
				if interval < 1000 {
					lr.period_count += 1
					//TODO process if qps > configuration
					//if lr.period_count > c.user.AuthIPs[ip]
					//return fmt.Errorf(`the query excess(%d) over the reqBurst(%d), sql: %s "`, lr.period_count, 1000, sqlstmt)
				} else {
					// new period
					if interval > 2000 { // previous period doesn`t have any query
						lr.lastqps = 0
					} else {
						lr.lastqps = lr.count
					}

					lr.last = now
					lr.period_count = 1
				}

			} else {
				lr := &LimitReqNode{}
				lr.query = fp
				lr.count = 1
				lr.last = getTimestamp()
				ipinfo.printfinger[fp] = lr
			}
			ipinfo.mu.Unlock()
		} else {
			ipinfo := &Ip{}
			ipinfo.ip = ip
			ipinfo.printfinger = make(map[string]*LimitReqNode)

		}
	} else {
		user := &User{}
		user.user = c.user.Username
		user.iplist = make(map[string]*Ip)

	}
}