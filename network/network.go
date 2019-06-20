package network

import (
	"net"
)

func RecvMessage(c *net.Conn) *Message {
	msg := NewMessage()
	(*c).Read(msg.Buffer[0:2]) // incoming message length
	if msg.Length() == 0 {
		return nil
	}
	bytes := make([]uint8, msg.Length())
	(*c).Read(bytes)
	msg.Buffer = append(msg.Buffer, bytes...)
	return msg
}

func SendMessage(c *net.Conn, msg *Message) {
	(*c).Write(msg.Buffer)
}
