package proxy

import (
	"encoding/binary"
	"fmt"
	"github.com/wangjild/go-mysql-proxy/client"
	. "github.com/wangjild/go-mysql-proxy/log"
	. "github.com/wangjild/go-mysql-proxy/mysql"
	"github.com/wangjild/go-mysql-proxy/sql"
	"strconv"
)

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

	s sql.IStatement

	sqlstmt string

	cstmt *client.Stmt
}

func (s *Stmt) ClearParams() {
	s.args = make([]interface{}, s.params)
}

func (s *Stmt) Close() {
	s.cstmt.Close(true)
}

func (c *Conn) handleComStmtPrepare(sqlstmt string) error {
	if c.schema == nil {
		return NewDefaultError(ER_NO_DB_ERROR)
	}

	s := new(Stmt)

	var err error
	s.s, err = sql.Parse(sqlstmt)
	if err != nil {
		return fmt.Errorf(`prepare parse sql "%s" error`, sqlstmt)
	}

	s.sqlstmt = sqlstmt

	var co *client.SqlConn
	co, err = c.schema.node.getMasterConn()
	// TODO tablename for select
	if err != nil {
		return fmt.Errorf("prepare error %s", err)
	}

	if err = co.UseDB(c.schema.db); err != nil {
		co.Close()
		return fmt.Errorf("parepre error %s", err)
	}

	if t, err := co.Prepare(sqlstmt); err != nil {
		co.Close()
		return fmt.Errorf("parepre error %s", err)
	} else {
		s.params = t.ParamNum()
		s.types = make([]byte, 0, s.params*2)
		s.columns = t.ColumnNum()
		s.bid = t.ID()
		s.cstmt = t
	}

	s.id = c.stmtId
	c.stmtId++

	if err = c.writePrepare(s); err != nil {
		return err
	}

	s.ClearParams()

	c.stmts[s.id] = s

	return nil
}

func (c *Conn) writePrepare(s *Stmt) error {
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

func (c *Conn) handleComStmtExecute(data []byte) error {
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

func (c *Conn) handleComStmtSendLongData(data []byte) error {
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

func (c *Conn) handleComStmtReset(data []byte) error {
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

func (c *Conn) handleComStmtClose(data []byte) error {
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
func (c *Conn) handleStmtExec(prepared *Stmt, data []byte, resultSet bool) error {

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
