package mysql

import (
	"bufio"
	"io"
	. "github.com/wangjild/go-mysql-proxy/log"
	"net"
)

type PacketIO struct {
	reader io.Reader
	writer io.Writer

	Sequence uint8
}

func NewPacketIO(conn net.Conn) *PacketIO {
	p := new(PacketIO)

	p.reader = bufio.NewReader(conn)
	p.writer = conn

	p.Sequence = 0

	return p
}

func (p *PacketIO) ReadPacket() ([]byte, error) {

	var payload []byte
	for {

		var header [PacketHeadSize]byte
		if n, err := io.ReadFull(p.reader, header[:]); err != nil {
			AppLog.Warn("wrong packet format, head size is %d", n)
			return nil, ErrBadConn
		}

		length := int(uint32(header[0]) | uint32(header[1])<<8 | uint32(header[2])<<16)
		if length < 1 {
			AppLog.Warn("wrong packet length, size is %d", length)
			return nil, ErrBadPkgLen
		}

		if uint8(header[3]) != p.Sequence {
			if uint8(header[3]) > p.Sequence {
				return nil, ErrPktSyncMul
			} else {
				return nil, ErrPktSync
			}
		}

		p.Sequence++

		data := make([]byte, length, length)
		var err error
		if _, err = io.ReadFull(p.reader, data); err != nil {
			AppLog.Warn("read packet from conn error: %s", err.Error())
			return nil, ErrBadConn
		}

		lastPacket := (length < MaxPacketSize)

		if lastPacket && payload == nil {
			return data, nil
		}

		payload = append(payload, data...)

		if lastPacket {
			return payload, nil
		}

	}
}

//data already have header
func (p *PacketIO) WritePacket(data []byte) error {
	length := len(data) - 4

	for length >= MaxPayloadLen {

		data[0] = 0xff
		data[1] = 0xff
		data[2] = 0xff

		data[3] = p.Sequence

		if n, err := p.writer.Write(data[:4+MaxPayloadLen]); err != nil {
			return ErrBadConn
		} else if n != (4 + MaxPayloadLen) {
			return ErrBadConn
		} else {
			p.Sequence++
			length -= MaxPayloadLen
			data = data[MaxPayloadLen:]
		}
	}

	data[0] = byte(length)
	data[1] = byte(length >> 8)
	data[2] = byte(length >> 16)
	data[3] = p.Sequence

	if n, err := p.writer.Write(data); err != nil {
		return ErrBadConn
	} else if n != len(data) {
		return ErrBadConn
	} else {
		p.Sequence++
		return nil
	}
}
