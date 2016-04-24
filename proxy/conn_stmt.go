package proxy

import (
	. "github.com/bytedance/dbatman/database/mysql"
	"github.com/bytedance/dbatman/database/sql"
	"github.com/bytedance/dbatman/parser"
	"github.com/ngaut/log"
)

/*

var paramFieldData []byte
var columnFieldData []byte

func init() {
	var p = &Field{Name: []byte("?")}
	var c = &Field{}

	paramFieldData = p.Dump()
	columnFieldData = c.Dump()
}

type Stmt struct {
	id  uint32
	bid uint32

	params  int
	types   []byte
	columns int

	args []interface{}

	s parser.IStatement

	sqlstmt string
}

func (s *Stmt) ClearParams() {
	s.args = make([]interface{}, s.params)
}

func (s *Stmt) Close() {
	s.cstmt.Close(true)
}
*/

func (c *Session) handleComStmtPrepare(sqlstmt string) error {
	stmt, err := parser.Parse(sqlstmt)
	if err != nil {
		log.Warningf(`parse sql "%s" error "%s"`, sqlstmt, err.Error())
		return c.handleMySQLError(
			NewDefaultError(ER_SYNTAX_ERROR, err.Error()))
	}

	// Only a few statements supported by prepare statements
	// http://dev.mysql.com/worklog/task/?id=2871
	switch v := stmt.(type) {
	case parser.ISelect, *parser.Insert, *parser.Update, *parser.Delete, *parser.Replace:
		return c.prepare(v, sqlstmt)
	case parser.IDDLStatement:
		// return c.prepareDDL(v, sqlstmt)
		return nil
	default:
		log.Warnf("statement %T[%s] not support prepare ops", stmt, sqlstmt)
		return c.handleMySQLError(
			NewDefaultError(ER_UNSUPPORTED_PS))
	}
}

func (session *Session) prepare(istmt parser.IStatement, sqlstmt string) error {
	if err := session.checkDB(istmt); err != nil {
		log.Debugf("check db error: %s", err.Error())
		return err
	}

	isread := false

	if s, ok := istmt.(parser.ISelect); ok {
		isread = !s.IsLocked()
	}

	if session.isInTransaction() || !session.isAutoCommit() {
		isread = false
	}

	stmt, err := session.Executor(isread).Prepare(sqlstmt)
	// TODO here should handler error
	if err != nil {
		return session.handleMySQLError(err)
	}

	return session.writePrepareResult(stmt)
}

func (session *Session) writePrepareResult(stmt *sql.Stmt) error {

	colen := len(stmt.Columns)
	paramlen := len(stmt.Params)

	// Prepare Header
	header := make([]byte, PacketHeaderLen, 12+PacketHeaderLen)

	// OK Status
	header = append(header, 0)
	header = append(header, byte(stmt.ID), byte(stmt.ID>>8), byte(stmt.ID>>16), byte(stmt.ID>>24))

	header = append(header, byte(colen), byte(colen>>8))
	header = append(header, byte(paramlen), byte(paramlen>>8))

	// reserved 00
	header = append(header, 0)

	// warning count 00
	// TODO
	header = append(header, 0, 0)

	if err := session.fc.WritePacket(header); err != nil {
		return session.handleMySQLError(err)
	}

	if paramlen > 0 {
		for _, p := range stmt.Params {
			if err := session.fc.WritePacket(p); err != nil {
				return session.handleMySQLError(err)
			}
		}

		if err := session.fc.WriteEOF(); err != nil {
			return session.handleMySQLError(err)
		}

	}

	if colen > 0 {
		for _, c := range stmt.Columns {
			if err := session.fc.WritePacket(c); err != nil {
				return session.handleMySQLError(err)
			}
		}

		if err := session.fc.WriteEOF(); err != nil {
			return session.handleMySQLError(err)
		}
	}

	return nil
}

/*
func (c *Session) writePrepare(s *Stmt) error {
	data := make([]byte, 4, 128)

	//status ok
	data = append(data, 0)
	//stmt id
	data = append(data, Uint32ToBytes(s.id)...)
	//number columns
	data = append(data, Uint16ToBytes(uint16(s.columns))...)
	//number params
	data = append(data, Uint16ToBytes(uint16(s.params))...)
	//filter [00]
	data = append(data, 0)
	//warning count
	data = append(data, 0, 0)

	if err := c.writePacket(data); err != nil {
		return err
	}

	if s.params > 0 {
		for i := 0; i < s.params; i++ {
			data = data[0:4]
			data = append(data, []byte(s.cstmt.ParamDefs[i])...)

			if err := c.writePacket(data); err != nil {
				return err
			}
		}

		if err := c.writeEOF(c.status); err != nil {
			return err
		}
	}

	if s.columns > 0 {
		for i := 0; i < s.columns; i++ {
			data = data[0:4]
			data = append(data, []byte(s.cstmt.ColDefs[i])...)

			if err := c.writePacket(data); err != nil {
				return err
			}
		}

		if err := c.writeEOF(c.status); err != nil {
			return err
		}

	}
	return nil
}

func (c *Session) handleComStmtExecute(data []byte) error {
	if len(data) < 9 {
		AppLog.Warn("ErrMalFormPacket: length %d", len(data))
		return ErrMalformPacket
	}

	pos := 0
	id := binary.LittleEndian.Uint32(data[0:4])
	pos += 4

	s, ok := c.stmts[id]
	if !ok {
		return NewDefaultError(ER_UNKNOWN_STMT_HANDLER,
			strconv.FormatUint(uint64(id), 10), "stmt_execute")
	}

	flag := data[pos]
	pos++

	//now we only support CURSOR_TYPE_NO_CURSOR flag
	if flag != 0 {
		return NewError(ER_UNKNOWN_ERROR, fmt.Sprintf("unsupported flag %d", flag))
	}

	s.cstmt.SetAttr(flag)

	//skip iteration-count, always 1
	pos += 4

	st, isread := s.s.(sql.ISelect)
	if isread {
		isread = (!st.IsLocked())
	}
	err := c.handleStmtExec(s, data[pos:], isread)

	s.ClearParams()

	return err
}

func (c *Session) handleComStmtSendLongData(data []byte) error {
	if len(data) < 6 {
		AppLog.Warn("ErrMalFormPacket")
		return ErrMalformPacket
	}

	id := binary.LittleEndian.Uint32(data[0:4])

	s, ok := c.stmts[id]
	if !ok {
		return NewDefaultError(ER_UNKNOWN_STMT_HANDLER,
			strconv.FormatUint(uint64(id), 10), "stmt_send_longdata")
	}

	paramId := binary.LittleEndian.Uint16(data[4:6])
	if paramId >= uint16(s.params) {
		return NewDefaultError(ER_WRONG_ARGUMENTS, "stmt_send_longdata")
	}

	s.cstmt.SendLongData(paramId, data[6:])
	return nil
}

func (c *Session) handleComStmtReset(data []byte) error {
	if len(data) < 4 {
		AppLog.Warn("ErrMalFormPacket")
		return ErrMalformPacket
	}

	id := binary.LittleEndian.Uint32(data[0:4])

	s, ok := c.stmts[id]
	if !ok {
		return NewDefaultError(ER_UNKNOWN_STMT_HANDLER,
			strconv.FormatUint(uint64(id), 10), "stmt_reset")
	}

	if r, err := s.cstmt.Reset(); err != nil {
		return err
	} else {
		s.ClearParams()
		return c.writeOK(r)
	}
}

func (c *Session) handleComStmtClose(data []byte) error {
	if len(data) < 4 {
		return nil
	}

	id := binary.LittleEndian.Uint32(data[0:4])

	if cstmt, ok := c.stmts[id]; ok {
		cstmt.Close()
	}

	delete(c.stmts, id)

	return nil
}

//
func (c *Session) handleStmtExec(prepared *Stmt, data []byte, resultSet bool) error {

	res, err := prepared.cstmt.Execute(data)
	if err != nil {
		return err
	}

	if resultSet {
		err = c.mergeSelectResult(res)
	} else {
		err = c.writeOK(res)
	}
	return err
}
*/
