package network

import (
	"fmt"
	"net"

	"github.com/rwxsu/goot/game"
)

func AddMapDescription(msg *Message, pos game.Position, m *game.Map) {
	msg.WriteUint8(0x64) // send map description
	AddPosition(msg, pos)

	// offset
	pos.X = pos.X - 8
	pos.Y = pos.Y - 6

	skip := uint8(0)

	if pos.Z < 8 {
		for z := (int8)(7); z > -1; z-- {
			for x := (uint16)(0); x < 18; x++ {
				for y := (uint16)(0); y < 14; y++ {
					tile := m.GetTile(game.Position{pos.X + x, pos.Y + y, (uint8)(z)})
					if tile == nil {
						skip++
						if skip == 0xff {
							msg.WriteUint8(skip)
							msg.WriteUint8(0xff)
							skip = 0
						}
					} else {
						if skip > 0 {
							msg.WriteUint8(skip)
							msg.WriteUint8(0xff)
							skip = 0
						}
						AddTile(msg, tile)
					}
				}
			}
		}
	} else { // underground

	}

	// Remainder
	msg.WriteUint8(skip)
	msg.WriteUint8(0xff)
}

func AddPosition(msg *Message, pos game.Position) {
	msg.WriteUint16(pos.X)
	msg.WriteUint16(pos.Y)
	msg.WriteUint8(pos.Z)
}

func AddTile(msg *Message, tile *game.Tile) {
	for _, i := range tile.Items {
		msg.WriteUint16(i.ID)
	}
	for _, c := range tile.Creatures {
		msg.WriteUint16(0x61)       // send add creature
		msg.WriteUint32(0x00)       // something about caching known creatures
		msg.WriteUint32(c.ID)       // ID
		msg.WriteString(c.Name)     // player name
		msg.WriteUint8(0x63)        // send look dir
		msg.WriteUint8(c.Direction) // look dir
		msg.WriteUint8(c.Outfit.Type)
		msg.WriteUint8(c.Outfit.Head)
		msg.WriteUint8(c.Outfit.Body)
		msg.WriteUint8(c.Outfit.Legs)
		msg.WriteUint8(c.Outfit.Feet)
		msg.WriteUint8(0) // light brightness?
		msg.WriteUint8(0) // light color?
		msg.WriteUint16(c.Speed)
		msg.WriteUint8(c.Skull)
		msg.WriteUint8(c.Party)
	}
	msg.WriteUint16(0xff00)
}

func ParseCharacterList(c *net.Conn) {
	characters := make([]game.Creature, 2)
	characters[0].Name = "admin"
	characters[0].World.Name = "test"
	characters[0].World.Port = 7171
	characters[1].Name = "rwxsu"
	characters[1].World.Name = "test"
	characters[1].World.Port = 7171
	res := NewMessage()
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
	SendMessage(c, res)
	res.HexDump("response")
}

func ParseLogin(c *net.Conn, m *game.Map) {
	res := NewMessage()
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

	tile := m.GetTile(player.Position)
	tile.AddCreature(player)
	AddMapDescription(res, player.Position, m)

	res.WriteUint8(0x83) // send magic effect
	AddPosition(res, player.Position)
	res.WriteUint8(0x0a) // type teleport

	res.WriteUint8(game.SlotEmpty)
	res.WriteUint8(game.SlotHead)

	res.WriteUint8(game.SlotEmpty)
	res.WriteUint8(game.SlotNecklace)

	res.WriteUint8(game.SlotNotEmpty)
	res.WriteUint8(game.SlotBackpack)
	res.WriteUint16(0x7c4) // backpack

	res.WriteUint8(game.SlotNotEmpty)
	res.WriteUint8(game.SlotBody)
	res.WriteUint16(0x9a8) // magic plate armor

	res.WriteUint8(game.SlotEmpty)
	res.WriteUint8(game.SlotShield)

	res.WriteUint8(game.SlotNotEmpty)
	res.WriteUint8(game.SlotWeapon)
	res.WriteUint16(0x997) // crossbow

	res.WriteUint8(game.SlotEmpty)
	res.WriteUint8(game.SlotLegs)

	res.WriteUint8(game.SlotEmpty)
	res.WriteUint8(game.SlotFeet)

	res.WriteUint8(game.SlotEmpty)
	res.WriteUint8(game.SlotRing)

	res.WriteUint8(game.SlotNotEmpty)
	res.WriteUint8(game.SlotAmmo)
	res.WriteUint16(0x9ef) // bolts
	res.WriteUint8(33)     // count

	res.WriteUint8(0xa0) // send player stats
	res.WriteUint16(player.HealthNow)
	res.WriteUint16(player.HealthMax)
	res.WriteUint16(player.Cap)
	res.WriteUint32(player.Combat.Experience)
	res.WriteUint8(player.Combat.Level)
	res.WriteUint8(player.Combat.Percent)
	res.WriteUint16(player.ManaNow)
	res.WriteUint16(player.ManaMax)
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

	SendMessage(c, res)
	res.HexDump("response")
}
