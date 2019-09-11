package main

import (
	"net"
	"path/filepath"

	"github.com/rwxsu/goot/game"
	"github.com/rwxsu/goot/network"
)

func main() {
	m := createMap("data/map/sectors/*")

	p := game.NewPlayer(1, "admin", game.Position{X: 32000, Y: 32000, Z: 7})

	l, _ := net.Listen("tcp", ":7171")
	defer l.Close()

	for {
		c, _ := l.Accept()
		run(c, m, p)
		c.Close()
	}
}

func createMap(sectors string) *game.Map {
	m := make(game.Map)
	filenames, _ := filepath.Glob(sectors)
	for _, fn := range filenames {
		m.LoadSector(fn)
	}
	return &m
}

func run(c net.Conn, m *game.Map, p *game.Player) {
	for {
		req := network.ReceiveMessage(c)
		if req != nil {
			code := req.ReadUint8()
			switch code {
			case 0x01: // character list
				req.SkipBytes(2)
				if req.ReadUint16() != 740 {
					network.SendInvalidClientVersion(c)
				} else {
					network.SendCharacterList(c)
				}
				return
			case 0x0a: // login with selected character
				req.SkipBytes(2)
				if req.ReadUint16() != 740 {
					network.SendInvalidClientVersion(c)
				} else {
					network.SendAddCreature(c, m, p)
				}
				return
			case 0x14: // login
				m.Tile(p.Position()).RemoveCreature(p)
				return
			default: // game commands
				network.ParseCommand(c, m, p, req, code)
			}
		}
	}
}
