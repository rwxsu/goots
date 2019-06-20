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

	player := game.Creature{
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

	l, err := net.Listen("tcp", ":7171")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		log.Println(err)
		return
	}

	// break: creates new listener
	// return: continues on same listener
	for {
		func(c *net.Conn) {
			// fmt.Println("\n[waiting]")
			req := network.RecvMessage(c)
			if req == nil {
				return
			}
			code := req.ReadUint8()
			req.HexDump(fmt.Sprintf("request code=0x%02x", code))
			switch code {
			case 0x01: // request character list
				if !network.ValidateClientVersion(c, req) {
					break
				}
				network.SendCharacterList(c)
				break
			case 0x0a: // request character login
				if !network.ValidateClientVersion(c, req) {
					return
				}
				network.SendAddCreature(c, &player, &m)
				return
			case 0x14: // logout
				break
			default:
				network.ParsePacket(c, &player, &m, code)
				return
			}
			(*c).Close()
			(*c), err = l.Accept()
			if err != nil {
				log.Println(err)
				return
			}
		}(&c)
	}
}
