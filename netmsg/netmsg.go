package netmsg

import (
	"bufio"
	"encoding/binary"
	"net"
)

func New(c *net.Conn) *NetMsg {
	return &NetMsg{
		conn:   c,
		reader: bufio.NewReader(*c),
		writer: bufio.NewWriter(*c),
	}
}

type NetMsg struct {
	conn   *net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
	buf    []byte
}

func (msg *NetMsg) SkipBytes(n int) {
	msg.reader.Discard(n)
}

func (msg *NetMsg) ReadUint8() byte {
	b, _ := msg.reader.ReadByte()
	return b
}

func (msg *NetMsg) ReadUint16() uint16 {
	b := make([]byte, 2)
	b[0] = msg.ReadUint8()
	b[1] = msg.ReadUint8()
	return binary.LittleEndian.Uint16(b)
}

func (msg *NetMsg) ReadUint32() uint32 {
	b := make([]byte, 4)
	b[0] = msg.ReadUint8()
	b[1] = msg.ReadUint8()
	b[2] = msg.ReadUint8()
	b[3] = msg.ReadUint8()
	return binary.LittleEndian.Uint32(b)
}

func (msg *NetMsg) ReadString() string {
	var s string
	length := (int)(msg.ReadUint16())
	for i := 0; i < length; i++ {
		s += (string)(msg.ReadUint8())
	}
	return s
}

func (msg *NetMsg) ResetReader() {
	msg.reader.Reset(*msg.conn)
}

func (msg *NetMsg) WriteUint8(b byte) {
	msg.buf = append(msg.buf, b)
}

func (msg *NetMsg) WriteUint16(i uint16) {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, i)
	msg.WriteUint8(b[0])
	msg.WriteUint8(b[1])
}

func (msg *NetMsg) WriteUint32(i uint32) {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, i)
	msg.WriteUint8(b[0])
	msg.WriteUint8(b[1])
	msg.WriteUint8(b[2])
	msg.WriteUint8(b[3])
}

func (msg *NetMsg) WriteString(s string) {
	msg.WriteUint16((uint16)(len(s)))
	for i := 0; i < len(s); i++ {
		msg.WriteUint8((byte)(s[i]))
	}
}

func (msg *NetMsg) ResetWriter() {
	msg.writer.Reset(*msg.conn)
	msg.buf = make([]byte, 2) // Reset packet size
}

func (msg *NetMsg) Send() {
	packetSize := (uint16)(len(msg.buf) - 2)
	binary.LittleEndian.PutUint16(msg.buf[0:2], packetSize)
	var i uint16
	for i = 0; i < packetSize+2; i++ {
		msg.writer.WriteByte(msg.buf[i])
	}
	msg.writer.Flush()
}
