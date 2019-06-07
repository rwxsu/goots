package main

import (
	"fmt"
	"net"

	"github.com/rwxsu/goot/game"

	"github.com/rwxsu/goot/constant"
	"github.com/rwxsu/goot/netmsg"
	"github.com/rwxsu/goot/packet"
)

func main() {
	fmt.Printf(":: Loading game info ")
	info := game.Info{
		World: "world",
	}
	fmt.Println("[done]")

	l, err := net.Listen("tcp", ":7171")
	if err != nil {
		fmt.Println(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go func(c net.Conn) {
			msg := netmsg.New(&c)
			req := OnRecvHeader(msg)
			switch req {
			case constant.RequestCharacterLogin:
				OnRequestCharacterLogin(msg, &info)
				break
			case constant.RequestCharacterList:
				OnRequestCharacterList(msg, &info)
				break
			}
		}(conn)
	}
}

func OnRecvHeader(msg *netmsg.NetMsg) uint8 {
	fmt.Println("\nheader.packet.len:", msg.ReadUint16())
	reqCode := msg.ReadUint8()
	msg.SkipBytes(2) // os := msg.ReadUint16()
	if msg.ReadUint16() != 740 {
		packet.SendMessage(msg, constant.MessageBoxSorry, "Only protocol 7.40 allowed!")
		return 0
	}
	return reqCode
}

func OnRequestCharacterList(msg *netmsg.NetMsg, info *game.Info) {
	fmt.Println("[OnRequestCharacterList]")
	msg.SkipBytes(12)
	// msg.ReadUint32() // Tibia.spr version
	// msg.ReadUint32() // Tibia.dat version
	// msg.ReadUint32() // Tibia.pic version
	acc := msg.ReadUint32()
	pwd := msg.ReadString()
	fmt.Println("login.acc:", acc)
	fmt.Println("login.pwd:", pwd)
	// TODO: Authenticate and retrieve characters
	// Dummy character list
	characters := make([]game.Character, 2)
	characters[0].Name = "rwxsu"
	characters[1].Name = "Test"
	packet.SendCharacterList(msg, info, characters)
}

func OnRequestCharacterLogin(msg *netmsg.NetMsg, info *game.Info) {
	fmt.Println("[OnRequestCharacterLogin]")
	msg.SkipBytes(1)
	acc := msg.ReadUint32()
	name := msg.ReadString()
	pwd := msg.ReadString()
	fmt.Printf("login.acc: %d\n", acc)
	fmt.Printf("login.pwd: %s\n", pwd)
	fmt.Printf("login.character.name: %s\n", name)
	// TODO: Authenticate
	character := game.Character{
		Name: name,
	}
	packet.SendCharacterLogin(msg, &character)
}
