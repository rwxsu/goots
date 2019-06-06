package packet

import (
	"github.com/rwxsu/goot/netmsg"
)

const (
	InfoFreePremium         = 0x00
	InfoOsWindows           = 0x02
	RequestCharacterList    = 0x01
	RequestCharacterLogin   = 0x0A
	ResponseDisconnect      = 0x0A
	ResponseMessageOfTheDay = 0x14
	ResponseCharacterList   = 0x64
)

func SendDisconnect(msg *netmsg.NetMsg, str string) {
	msg.ResetWriter()
	msg.WriteUint8(ResponseDisconnect)
	msg.WriteString(str)
	msg.Send()
}

func SendCharacterList(msg *netmsg.NetMsg) {
	msg.ResetWriter()

	msg.WriteUint8(ResponseMessageOfTheDay)
	msg.WriteString("Welcome to GoOT!")

	msg.WriteUint8(ResponseCharacterList)
	msg.WriteUint8(0x01) // Character count

	msg.WriteString("rwxsu")
	msg.WriteString("world")
	msg.WriteUint8(127)   // IP 127
	msg.WriteUint8(0)     // 0
	msg.WriteUint8(0)     // 0
	msg.WriteUint8(1)     // 1
	msg.WriteUint16(7171) // Port to send login packet to (usually game port 7172)

	msg.WriteUint16(InfoFreePremium)
	msg.Send()
}
