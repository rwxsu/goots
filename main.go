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

	const sectors = "data/map/sectors/*"
	filenames, _ := filepath.Glob(sectors)

	for _, fn := range filenames {
		m.LoadSector(fn)
	}

	l, err := net.Listen("tcp", ":7171")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	acceptConnections(l, m)
}

func acceptConnections(listener net.Listener, m game.Map) {
	for {
		c, err := listener.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		go handleConnection(c, m)
	}
}

func handleConnection(c net.Conn, m game.Map) {
	// Placeholder player
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
		Light: game.Light{Level: 0x7, Color: 0xd7},
		World: game.World{Light: game.Light{Level: 0x00, Color: 0xd7}},
		Speed: 70,
	}
connectionLoop:
	for {
		req := network.RecvMessage(c)
		if req == nil {
			return
		}
		code := req.ReadUint8()
		switch code {
		case 0x01: // request character list
			req.SkipBytes(2) // os := req.ReadUint16()
			if req.ReadUint16() != 740 {
				network.SendInvalidClientVersion(c)
				break connectionLoop
			}
			network.SendCharacterList(c)
			break
		case 0x0a: // request character login
			req.SkipBytes(2) // os := req.ReadUint16()
			if req.ReadUint16() != 740 {
				network.SendInvalidClientVersion(c)
				break connectionLoop
			}
			network.SendAddCreature(c, &player, &m)
		case 0x14: // logout
			break connectionLoop
		default:
			network.ParseCommand(c, req, &player, &m, code)
		}
	}
	c.Close()
}
