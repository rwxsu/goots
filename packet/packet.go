package packet

import (
	"fmt"

	"github.com/rwxsu/goot/constant"
	"github.com/rwxsu/goot/game"
	"github.com/rwxsu/goot/netmsg"
)

func SendMessage(msg *netmsg.NetMsg, kind byte, str string) {
	msg.ResetWriter()
	msg.WriteUint8(kind)
	msg.WriteString(str)
	msg.Send()
}

func SendCharacterList(msg *netmsg.NetMsg, info *game.Info, characters []game.Character) {
	msg.ResetWriter()

	msg.WriteUint8(constant.ResponseMOTD)
	msg.WriteString("Welcome to GoOT!")

	msg.WriteUint8(constant.ResponseCharList)
	msg.WriteUint8((byte)(len(characters))) // Character count

	for i := 0; i < len(characters); i++ {
		msg.WriteString(characters[i].Name)
		msg.WriteString(info.World)
		msg.WriteUint8(127) // IP 127.0.0.1
		msg.WriteUint8(0)
		msg.WriteUint8(0)
		msg.WriteUint8(1)
		msg.WriteUint16(7171) // Port to send login packet to (usually game port 7172)
	}

	msg.WriteUint16(constant.FreePremium)
	msg.Send()
}

func SendCharacterLogin(msg *netmsg.NetMsg, character *game.Character) {
	SendMessage(msg, constant.MessageBoxInfo, fmt.Sprintf("Welcome, %s!", character.Name))
}
