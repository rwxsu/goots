package main

import (
	"fmt"
	"log"
	"net"

	"github.com/rwxsu/goot/game"
	"github.com/rwxsu/goot/network"
)

func main() {
	l, err := net.Listen("tcp", ":7171")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func(c *net.Conn) {
			req := network.RecvPacket(c)
			req.HexDump("request")
			code := req.ReadUint8()
			req.SkipBytes(2) // os := req.ReadUint16()
			if req.ReadUint16() != 740 {
				res := network.NewPacket()
				res.WriteUint8(0x0a)
				res.WriteString("Only protocol 7.40 allowed!")
				network.SendPacket(c, res)
				res.HexDump("response")
				return
			}
			switch code {
			case 0x01: // request character list
				req.ReadUint32() // Tibia.spr version
				req.ReadUint32() // Tibia.dat version
				req.ReadUint32() // Tibia.pic version
				req.ReadUint32() // acc
				req.ReadString() // pwd
				characters := make([]game.Character, 2)
				characters[0].Name = "admin"
				characters[0].World.Name = "test"
				characters[0].World.Port = 7171
				characters[1].Name = "rwxsu"
				characters[1].World.Name = "test"
				characters[1].World.Port = 7171
				res := network.NewPacket()
				res.WriteUint8(0x14) // MOTD
				res.WriteString("Welcome to GoOT.")
				res.WriteUint8(0x64) // character list
				res.WriteUint8((uint8)(len(characters)))
				for i := 0; i < len(characters); i++ {
					res.WriteString(characters[i].Name)
					res.WriteString(characters[i].World.Name)
					res.WriteUint8(127)
					res.WriteUint8(0)
					res.WriteUint8(0)
					res.WriteUint8(1)
					res.WriteUint16(characters[i].World.Port)
				}
				res.WriteUint16(0) // premium days
				network.SendPacket(c, res)
				res.HexDump("response")
				return
			case 0x0a: // request character login
				req.SkipBytes(1) // ?
				req.ReadUint32() // acc
				name := req.ReadString()
				req.ReadString() // pwd
				res := network.NewPacket()
				res.WriteUint8(0x15) // FYI
				res.WriteString(fmt.Sprintf("Welcome, %s!", name))
				network.SendPacket(c, res)
				res.HexDump("response")
				return
			}
		}(&c)
	}
}
