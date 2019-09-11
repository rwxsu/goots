package network

import (
	"net"
)

// ReceiveMessage reads the incoming message length (first two bytes), followed by
// the message.
func ReceiveMessage(c net.Conn) *Message {
	msg := NewMessage()
	c.Read(msg.Buffer[0:2]) // read incoming message length
	if msg.Length() == 0 {
		return nil
	}
	bytes := make([]uint8, msg.Length())
	c.Read(bytes) // read rest of message
	msg.Buffer = append(msg.Buffer, bytes...)
	return msg
}

func SendMessage(c net.Conn, msg *Message) {
	c.Write(msg.Buffer)
}
