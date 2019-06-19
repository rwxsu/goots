package main

import (
	"fmt"
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
				player := &game.Creature{
					ID:        0x04030201,
					Access:    game.Tutor,
					Name:      "rwxsu",
					Cap:       50,
					Combat:    game.Skill{Level: 8, Percent: 0, Experience: 4200},
					HealthNow: 100,
					HealthMax: 200,
					ManaNow:   50,
					ManaMax:   100,
					Magic:     game.Skill{Level: 10, Percent: 50},
					Fist:      game.Skill{Level: 10, Percent: 50},
					Club:      game.Skill{Level: 10, Percent: 50},
					Sword:     game.Skill{Level: 10, Percent: 50},
					Axe:       game.Skill{Level: 10, Percent: 50},
					Distance:  game.Skill{Level: 10, Percent: 50},
					Shielding: game.Skill{Level: 10, Percent: 50},
					Fishing:   game.Skill{Level: 10, Percent: 50},
					Direction: game.South,
					Position:  game.Position{X: 32000, Y: 32000, Z: 7},
					Outfit: game.Outfit{
						Type: 0x80,
						Head: 0x50,
						Body: 0x50,
						Legs: 0x50,
						Feet: 0x50,
					},
					Skull: 3,
					Icons: 1,
					Light: game.Light{0x7, 0xd7},
					World: game.World{Light: game.Light{0x00, 0xd7}},
				}
				network.SendAddCreature(c, player, &m)
				return
			case 0x65:
				fmt.Println("parseMove:NORTH")
				break
			case 0x66:
				fmt.Println("parseMove:EAST")
				break
			case 0x67:
				fmt.Println("parseMove:SOUTH")
				break
			case 0x68:
				fmt.Println("parseMove:WEST")
				break
			}
		}(&c)
	}
}
