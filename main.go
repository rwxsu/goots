package main

import (
	"net"
	"path/filepath"

	"github.com/rwxsu/goot/game"
	"github.com/rwxsu/goot/network"
)

func main() {
	const sectors = "data/map/sectors/*"
	filenames, _ := filepath.Glob(sectors)

	m := make(game.Map)
	player := game.NewPlayer(1, "admin", game.Position{X: 32000, Y: 32000, Z: 7})

	for _, fn := range filenames {
		m.LoadSector(fn)
	}

	cm := network.NewConnectionManager()

	l, _ := net.Listen("tcp", ":7171")
	defer l.Close()

	for {
		c, _ := l.Accept()
		cm.Add(&network.TibiaConnection{Connection: c})
		tc := cm.ByConnection(c)
		tc.Map = &m
		tc.Player = player
		func(c net.Conn) {
			for {
				if req := network.ReceiveMessage(tc.Connection); req != nil {
					code := req.ReadUint8()
					switch code {
					case 0x01:
						req.SkipBytes(2) // os := req.ReadUint16()
						if req.ReadUint16() != 740 {
							network.SendInvalidClientVersion(tc.Connection)
							return // breaks out of IIFE
						}
						network.SendCharacterList(tc.Connection)
						return // breaks out of IIFE
					case 0x0a:
						req.SkipBytes(2) // os := req.ReadUint16()
						if req.ReadUint16() != 740 {
							network.SendInvalidClientVersion(tc.Connection)
							return // breaks out of IIFE
						}
						network.SendAddCreature(tc)
					case 0x14: // logout
						tc.Map.Tile(tc.Player.Position()).RemoveCreature(tc.Player)
						return // breaks out of IIFE
					default:
						network.ParseCommand(tc, req, code)
						// continue inside IIFE
					}
				}
			}
		}(c)
		cm.Delete(tc)
		c.Close()
	}
}
