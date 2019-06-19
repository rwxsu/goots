package main

import (
	"log"
	"net"
	"path/filepath"

	"github.com/rwxsu/goot/game"
	"github.com/rwxsu/goot/network"
)

func main() {
	m := make(game.Map)

	filenames, err := filepath.Glob("data/map/sectors/*")
	if err != nil {
		panic(err)
	}
	for _, fn := range filenames {
		m.LoadSector(fn)
	}

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
			req := network.RecvMessage(c)
			req.HexDump("request")
			code := req.ReadUint8()
			req.SkipBytes(2) // os := req.ReadUint16()
			if req.ReadUint16() != 740 {
				res := network.NewMessage()
				res.WriteUint8(0x0a)
				res.WriteString("Only protocol 7.40 allowed!")
				network.SendMessage(c, res)
				res.HexDump("response")
				return
			}
			switch code {
			case 0x01: // request character list
				network.ParseCharacterList(c)
				return
			case 0x0a: // request character login
				network.ParseLogin(c, &m)
				return
			}
		}(&c)
	}
}
