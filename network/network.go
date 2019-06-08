package network

import (
	"net"
)

func RecvPacket(c *net.Conn) *Packet {
	p := NewPacket()
	(*c).Read(p.Buffer[0:2]) // incoming packet length
	bytes := make([]uint8, p.Length())
	(*c).Read(bytes)
	p.Buffer = append(p.Buffer, bytes...)
	return p
}

func SendPacket(c *net.Conn, p *Packet) {
	(*c).Write(p.Buffer)
}
