package main

import (
	"log"
	"net"
	"path/filepath"

	"github.com/maksumic/goot/game"
	"github.com/maksumic/goot/network"
)

func main() {
	m := make(game.Map)
	filenames, _ := filepath.Glob("data/map/sectors/*")
	for _, filename := range filenames {
		m.LoadSector(filename)
	}

	l, err := net.Listen("tcp", ":7171")
	if err != nil {
		panic(err)
	}
	
	p := game.NewPlayer(1, "admin", game.Position{X: 32000, Y: 32000, Z: 7})

	c, err := l.Accept()
	if err != nil {
		log.Println(err)
	}

	for {
		req := network.ReceiveMessage(c)
		if req != nil {
			code := req.ReadUint8()
			switch code {
			case 0x01:
				req.SkipBytes(2)
				if req.ReadUint16() != 740 {
					network.SendInvalidClientVersion(c)
				} else {
					network.SendCharacterList(c)
				}
				c.Close()
				c, err = l.Accept()
				if err != nil {
					log.Println(err)
				}
			case 0x0a:
				req.SkipBytes(2)
				if req.ReadUint16() != 740 {
					network.SendInvalidClientVersion(c)
				} else {
					network.SendAddCreature(c, &m, p)
				}
			case 0x14:
				m.Tile(p.Position()).RemoveCreature(p)
				c.Close()
				c, err = l.Accept()
				if err != nil {
					log.Println(err)
				}
			default:
				network.ParseCommand(c, &m, p, req, code)
			}
		}
	}
}
