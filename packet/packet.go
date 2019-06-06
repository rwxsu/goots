package packet

import (
	"github.com/rwxsu/goot/netmsg"
)

const (
	INFO_OS_WINDOWS = 0x02
)

const (
	REQUEST_CHARACTER_LIST  = 0x01
	REQUEST_CHARACTER_LOGIN = 0x0A
)

const (
	RESPONSE_DISCONNECT         = 0x0A
	RESPONSE_MESSAGE_OF_THE_DAY = 0x14
	RESPONSE_CHARACTER_LIST     = 0x64
)

func SendDisconnect(msg *netmsg.NetMsg, str string) {
	msg.ResetWriter()
	msg.WriteUint8(RESPONSE_DISCONNECT)
	msg.WriteString(str)
	msg.Send()
}

// TODO: func SendCharacterList(msg, characters)
func SendCharacterList(msg *netmsg.NetMsg) {
	msg.ResetWriter()

	motd := "Welcome to GoOT!"
	world := "GoOT"
	charName := "rwxsu"

	msg.WriteUint8(RESPONSE_MESSAGE_OF_THE_DAY)
	msg.WriteString(motd)

	msg.WriteUint8(RESPONSE_CHARACTER_LIST)
	msg.WriteUint8(0x01) // Character count

	msg.WriteString(charName)
	msg.WriteString(world)
	msg.WriteUint8(127)   // IP 127
	msg.WriteUint8(0)     // 0
	msg.WriteUint8(0)     // 0
	msg.WriteUint8(1)     // 1
	msg.WriteUint16(7171) // Port to send login packet to (usually game port 7172)

	msg.WriteUint16(0x00) // Premium days
	msg.Send()
}
