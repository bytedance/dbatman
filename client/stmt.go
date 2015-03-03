package client

import (
	"encoding/binary"
	. "github.com/wangjild/go-mysql-proxy/mysql"
)

type Stmt struct {
	conn  *Conn
	id    uint32
	query string

	params    int
	ParamDefs [][]byte

	columns int
	ColDefs [][]byte

	flag byte
}

func (s *Stmt) ID() uint32 {
	return s.id
}

func (s *Stmt) ParamNum() int {
	return s.params
}

func (s *Stmt) ColumnNum() int {
	return s.columns
}

func (s *Stmt) Execute(data []byte) (*Result, error) {
	if err := s.write(data); err != nil {
		return nil, err
	}

	return s.conn.readResult(true)
}

func (s *Stmt) Close(closeConn bool) error {
	if err := s.conn.writeCommandUint32(COM_STMT_CLOSE, s.id); err != nil {
		return err
	}

	if closeConn {
		s.conn.Close()
		s.conn = nil
	}

	return nil
}

func (s *Stmt) write(param []byte) error {

	data := make([]byte, 4, 4+9+len(param))

	data = append(data, COM_STMT_EXECUTE)

	data = append(data, byte(s.id), byte(s.id>>8), byte(s.id>>16), byte(s.id>>24))

	//flag: CURSOR_TYPE_NO_CURSOR
	data = append(data, s.flag)

	data = append(data, 1, 0, 0, 0)

	data = append(data, param...)

	s.conn.pkg.Sequence = 0
	return s.conn.writePacket(data)
}

func (s *Stmt) SendLongData(pid uint16, payload []byte) error {

	data := make([]byte, 4, 4+7+len(payload))
	data = append(data, COM_STMT_SEND_LONG_DATA)
	data = append(data, byte(s.id), byte(s.id>>8), byte(s.id>>16), byte(s.id>>24))
	data = append(data, byte(pid), byte(pid>>8))

	data = append(data, payload...)

	s.conn.pkg.Sequence = 0
	return s.conn.writePacket(data)
}

func (s *Stmt) Reset() (*Result, error) {
	if err := s.conn.writeCommandUint32(COM_STMT_RESET, s.id); err != nil {
		return nil, err
	}

	s.flag = CURSOR_TYPE_NO_CURSOR
	return s.conn.readOK()
}

func (c *Conn) Prepare(query string) (*Stmt, error) {
	if err := c.writeCommandStr(COM_STMT_PREPARE, query); err != nil {
		return nil, err
	}

	data, err := c.readPacket()
	if err != nil {
		return nil, err
	}

	if data[0] == ERR_HEADER {
		return nil, c.handleErrorPacket(data)
	} else if data[0] != OK_HEADER {
		return nil, ErrMalformPacket
	}

	s := new(Stmt)
	s.conn = c

	pos := 1

	//for statement id
	s.id = binary.LittleEndian.Uint32(data[pos:])
	pos += 4

	//number columns
	s.columns = int(binary.LittleEndian.Uint16(data[pos:]))
	pos += 2

	//number params
	s.params = int(binary.LittleEndian.Uint16(data[pos:]))
	pos += 2

	//warnings
	//warnings = binary.LittleEndian.Uint16(data[pos:])

	if s.params > 0 {
		if ps, err := s.conn.readUntilEOF(s.params); err != nil {
			return nil, err
		} else {
			s.ParamDefs = ps
		}
	}

	if s.columns > 0 {
		if cs, err := s.conn.readUntilEOF(s.columns); err != nil {
			return nil, err
		} else {
			s.ColDefs = cs
		}
	}

	s.query = query
	return s, nil
}

func (s *Stmt) SetAttr(f byte) {
	s.flag = f
}
