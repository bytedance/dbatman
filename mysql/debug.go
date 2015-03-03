package mysql

import (
	"bytes"
	"fmt"
)

const (
	InitialHandshake = "Initial Handshake Packet"
)

func Dump(types string, action string, cId uint32, capability uint32, data []byte) {
	fmt.Println("---------------mysql packet dump----------------")
	fmt.Println("Action:", action)
	fmt.Println("ConnectionId:", cId)
	fmt.Println("PacketType:", types)

	switch types {
	case InitialHandshake:
		dumpInitialHandshake(data, capability)
	default:
		fmt.Println("Unsupport packet type")
	}
}

func dumpInitialHandshake(data []byte, capability uint32) {
	fmt.Println("\tProtocal Version:", data[0])

	strlen := bytes.IndexByte(data[1:], 0x00)
	fmt.Println("\tServer:", string(data[1:1+strlen]), "00")
	pos := 1 + strlen + 1

	//cipher := data[pos : pos+8]
	pos += 8

	// capability lower 2 byte
	// capa := binary.LittleEndian.Uint16(data[pos : pos+2])
	pos += 2

	if len(data) > pos {
		pos += 1 + 2 + 2 + 1 + 10
		//	cipher = append(cipher, data[pos:pos+12]...)
	}

	//fmt.Println("\tCipher:", cipher)
}
