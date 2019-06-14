package main

import (
	"fmt"
	"log"
	"net"

	"github.com/rwxsu/goot/source/game"
	"github.com/rwxsu/goot/source/network"
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
				// req.SkipBytes(1) // ?
				// req.ReadUint32() // acc
				// name := req.ReadString()
				// req.ReadString() // pwd

				res := network.NewPacket()

				player := &game.Character{
					Access:    game.Tutor,
					ID:        0x04030201,
					Name:      "rwxsu",
					Combat:    game.Skill{Level: 8, Percent: 0, Experience: 4200},
					Cap:       50,
					Health:    game.Health{Now: 100, Max: 200},
					Mana:      game.Mana{Now: 20, Max: 35},
					Magic:     game.Skill{Level: 10, Percent: 50},
					Fist:      game.Skill{Level: 10, Percent: 50},
					Club:      game.Skill{Level: 10, Percent: 50},
					Sword:     game.Skill{Level: 10, Percent: 50},
					Axe:       game.Skill{Level: 10, Percent: 50},
					Distance:  game.Skill{Level: 10, Percent: 50},
					Shielding: game.Skill{Level: 10, Percent: 50},
					Fishing:   game.Skill{Level: 10, Percent: 50},
					Direction: game.West,
					Position:  game.Position{X: 100, Y: 100, Z: 7},
					Outfit: game.Outfit{
						Type: 0x80,
						Head: 0x50,
						Body: 0x50,
						Legs: 0x50,
						Feet: 0x50,
					},
				}

				res.WriteUint8(0x0a)
				res.WriteUint32(player.ID) // ID
				res.WriteUint16(0x32)      // ?

				// can report bugs?
				if player.Access > game.Regular {
					res.WriteUint8(0x01)
				} else {
					res.WriteUint8(0x00)
				}

				// ?
				if player.Access >= game.Gamemaster {
					res.WriteUint8(0x0b)
					for i := 0; i < 32; i++ {
						res.WriteUint8(0xff)
					}
				}

				res.WriteUint8(0x64) // send map description

				res.WriteUint16(player.Position.X)
				res.WriteUint16(player.Position.Y)
				res.WriteUint8(player.Position.Z)

				// tile structure
				// ground
				// items
				// tiles to skip
				// 0xff

				// column 0 (out of screen) floor 7
				for i := 0; i < 14; i++ {
					res.WriteUint8(0x64) // void
					res.WriteUint8(0x00) // void
					res.WriteUint8(0x00)
					res.WriteUint8(0xff)
				}

				// column 1 floor 7
				for i := 0; i < 14; i++ {
					res.WriteUint8(0x65) // dirt
					res.WriteUint8(0x00) // dirt
					res.WriteUint8(0x00)
					res.WriteUint8(0xff)
				}

				// column 2 floor 7
				for i := 0; i < 14; i++ {
					res.WriteUint8(0x66) // grass
					res.WriteUint8(0x00) // grass
					res.WriteUint8(0x00)
					res.WriteUint8(0xff)
				}

				// column 3 floor 7
				for i := 0; i < 14; i++ {
					res.WriteUint8(0x9e) // snow
					res.WriteUint8(0x02) // snow
					res.WriteUint8(0x00)
					res.WriteUint8(0xff)
				}

				// column 4 floor 7
				for i := 0; i < 14; i++ {
					res.WriteUint8(0x64) // void
					res.WriteUint8(0x00) // void
					res.WriteUint8(0x00)
					res.WriteUint8(0xff)
				}

				// column 5 floor 7
				for i := 0; i < 14; i++ {
					res.WriteUint8(0x65) // dirt
					res.WriteUint8(0x00) // dirt
					res.WriteUint8(0x00)
					res.WriteUint8(0xff)
				}

				// column 6 floor 7
				for i := 0; i < 14; i++ {
					res.WriteUint8(0x66) // grass
					res.WriteUint8(0x00) // grass
					res.WriteUint8(0x00)
					res.WriteUint8(0xff)
				}

				// column 7 floor 7
				for i := 0; i < 14; i++ {
					res.WriteUint8(0x9e) // snow
					res.WriteUint8(0x02) // snow
					res.WriteUint8(0x00)
					res.WriteUint8(0xff)
				}

				// column 8 (before player) floor 7
				for i := 0; i < 6; i++ {
					res.WriteUint8(0x66) // grass
					res.WriteUint8(0x00) // grass
					res.WriteUint8(0x00)
					res.WriteUint8(0xff)
				}

				res.WriteUint8(0x95) // wooden floor
				res.WriteUint8(0x01) // wooden floor

				res.WriteUint16(0x61)      // send add creature
				res.WriteUint32(0x00)      // something about caching known creatures
				res.WriteUint32(player.ID) // ID
				res.WriteString("rwxsu")   // player name

				res.WriteUint8(0x63)             // send look dir
				res.WriteUint8(player.Direction) // look dir

				res.WriteUint8(player.Outfit.Type)
				res.WriteUint8(player.Outfit.Head)
				res.WriteUint8(player.Outfit.Body)
				res.WriteUint8(player.Outfit.Legs)
				res.WriteUint8(player.Outfit.Feet)
				res.WriteUint8(0x00)
				res.WriteUint8(0x00)
				res.WriteUint8(0xa2)
				res.WriteUint8(0x01)
				res.WriteUint8(0x00)
				res.WriteUint8(0x00)

				res.WriteUint8(0x00)
				res.WriteUint8(0xff)

				// column 8 (after player) floor 7
				for i := 0; i < 7; i++ {
					res.WriteUint8(0x66) // grass
					res.WriteUint8(0x00) // grass
					res.WriteUint8(0x00)
					res.WriteUint8(0xff)
				}

				// column 9 floor 7
				for i := 0; i < 14; i++ {
					res.WriteUint8(0x67) // dirt
					res.WriteUint8(0x00) // dirt
					res.WriteUint8(0x00)
					res.WriteUint8(0xff)
				}

				// column 10 floor 7
				for i := 0; i < 14; i++ {
					res.WriteUint8(0x9e) // snow
					res.WriteUint8(0x02) // snow
					res.WriteUint8(0x00)
					res.WriteUint8(0xff)
				}

				// column 11 floor 7
				for i := 0; i < 14; i++ {
					res.WriteUint8(0x64) // void
					res.WriteUint8(0x00) // void
					res.WriteUint8(0x00)
					res.WriteUint8(0xff)
				}

				// column 12 floor 7
				for i := 0; i < 14; i++ {
					res.WriteUint8(0x9e) // snow
					res.WriteUint8(0x02) // snow
					res.WriteUint8(0x00)
					res.WriteUint8(0xff)
				}

				// column 13 floor 7
				for i := 0; i < 14; i++ {
					res.WriteUint8(0x66) // grass
					res.WriteUint8(0x00) // snow
					res.WriteUint8(0x00)
					res.WriteUint8(0xff)
				}

				// column 14 floor 7
				for i := 0; i < 14; i++ {
					res.WriteUint8(0x9e) // snow
					res.WriteUint8(0x02) // snow
					res.WriteUint8(0x00)
					res.WriteUint8(0xff)
				}

				// column 15 floor 7
				for i := 0; i < 14; i++ {
					res.WriteUint8(0x67) // dirt
					res.WriteUint8(0x00) // dirt
					res.WriteUint8(0x00)
					res.WriteUint8(0xff)
				}

				// column 16 floor 7
				for i := 0; i < 14; i++ {
					res.WriteUint8(0x9e) // snow
					res.WriteUint8(0x02) // snow
					res.WriteUint8(0x00)
					res.WriteUint8(0xff)
				}

				// column 17 floor 7
				for i := 0; i < 14; i++ {
					res.WriteUint16(0x029e)
					res.WriteUint16(0x05a2) // statue
					res.WriteUint16(0xff00)
				}

				// skip 6th floor
				res.WriteUint8(0xfc) // hex(fc) = decimal(18) * decimal(14)
				res.WriteUint8(0xff)

				// skip 5th floor
				res.WriteUint8(0xfc)
				res.WriteUint8(0xff)

				// skip 4th floor
				res.WriteUint8(0xfc)
				res.WriteUint8(0xff)

				// skip 3rd floor
				res.WriteUint8(0xfc)
				res.WriteUint8(0xff)

				// skip 2th floor
				res.WriteUint8(0xfc)
				res.WriteUint8(0xff)

				// skip 1th floor
				res.WriteUint8(0xfc)
				res.WriteUint8(0xff)

				// skip 0th floor??
				res.WriteUint8(0xfc)
				res.WriteUint8(0xff)

				res.WriteUint8(0x83) // send magic effect
				res.WriteUint16(player.Position.X)
				res.WriteUint16(player.Position.Y)
				res.WriteUint8(player.Position.Z)
				res.WriteUint8(0x0a) // type teleport

				res.WriteUint8(game.Empty)
				res.WriteUint8(game.Head)

				res.WriteUint8(game.Empty)
				res.WriteUint8(game.Necklace)

				res.WriteUint8(game.NotEmpty)
				res.WriteUint8(game.Backpack)
				res.WriteUint16(0x7c4) // backpack

				res.WriteUint8(game.NotEmpty)
				res.WriteUint8(game.Body)
				res.WriteUint16(0x9a8) // magic plate armor

				res.WriteUint8(game.Empty)
				res.WriteUint8(game.Shield)

				res.WriteUint8(game.NotEmpty)
				res.WriteUint8(game.Weapon)
				res.WriteUint16(0x997) // crossbow

				res.WriteUint8(game.Empty)
				res.WriteUint8(game.Legs)

				res.WriteUint8(game.Empty)
				res.WriteUint8(game.Feet)

				res.WriteUint8(game.Empty)
				res.WriteUint8(game.Ring)

				res.WriteUint8(game.NotEmpty)
				res.WriteUint8(game.Ammo)
				res.WriteUint16(0x9ef) // bolts
				res.WriteUint8(33)     // count

				res.WriteUint8(0xa0) // send player stats
				res.WriteUint16(player.Health.Now)
				res.WriteUint16(player.Health.Max)
				res.WriteUint16(player.Cap)
				res.WriteUint32(player.Combat.Experience)
				res.WriteUint8(player.Combat.Level)
				res.WriteUint8(player.Combat.Percent)
				res.WriteUint16(player.Mana.Now)
				res.WriteUint16(player.Mana.Max)
				res.WriteUint8(player.Magic.Level)
				res.WriteUint8(player.Magic.Percent)

				res.WriteUint8(0xa1) // send player skills
				res.WriteUint8(player.Fist.Level)
				res.WriteUint8(player.Fist.Percent)
				res.WriteUint8(player.Club.Level)
				res.WriteUint8(player.Club.Percent)
				res.WriteUint8(player.Sword.Level)
				res.WriteUint8(player.Sword.Percent)
				res.WriteUint8(player.Axe.Level)
				res.WriteUint8(player.Axe.Percent)
				res.WriteUint8(player.Distance.Level)
				res.WriteUint8(player.Distance.Percent)
				res.WriteUint8(player.Shielding.Level)
				res.WriteUint8(player.Shielding.Percent)
				res.WriteUint8(player.Fishing.Level)
				res.WriteUint8(player.Fishing.Percent)

				res.WriteUint8(0x82) // send light
				res.WriteUint8(0xfa) // light level
				res.WriteUint8(0xd7) // light unknown

				res.WriteUint8(0x8d)

				res.WriteUint32(player.ID) // ID

				res.WriteUint8(0x00)
				res.WriteUint8(0x00)

				res.WriteUint8(0xb4)
				res.WriteUint8(0x15)

				res.WriteString(fmt.Sprintf("Welcome, %s.", player.Name))

				res.WriteUint8(0xb4)
				res.WriteUint8(0x15)

				res.WriteString("Last login string")

				res.WriteUint8(0x8d)

				res.WriteUint32(player.ID) // ID

				res.WriteUint8(0x0a)
				res.WriteUint8(0xd7)

				res.WriteUint8(0xa2) // send icons
				res.WriteUint8(0x00) // icons

				network.SendPacket(c, res)
				res.HexDump("response")
				return
			}
		}(&c)
	}
}
