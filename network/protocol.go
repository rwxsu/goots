package network

import (
	"fmt"
	"net"
	"time"

	"github.com/rwxsu/goot/game"
)

func AddCreatureLight(msg *Message, c *game.Creature) {
	msg.WriteUint8(0x8d)
	msg.WriteUint32(c.ID)
	msg.WriteUint8(c.Light.Level)
	msg.WriteUint8(c.Light.Color)
}

func AddWorldLight(msg *Message, w *game.World) {
	msg.WriteUint8(0x82)
	msg.WriteUint8(w.Light.Level) // 0xfa
	msg.WriteUint8(w.Light.Color) // 0xd7
}

func AddIcons(msg *Message, c *game.Creature) {
	msg.WriteUint8(0xa2)
	msg.WriteUint8(c.Icons)
}

func AddSkills(msg *Message, c *game.Creature) {
	msg.WriteUint8(0xa1)
	msg.WriteUint8(c.Fist.Level)
	msg.WriteUint8(c.Fist.Percent)
	msg.WriteUint8(c.Club.Level)
	msg.WriteUint8(c.Club.Percent)
	msg.WriteUint8(c.Sword.Level)
	msg.WriteUint8(c.Sword.Percent)
	msg.WriteUint8(c.Axe.Level)
	msg.WriteUint8(c.Axe.Percent)
	msg.WriteUint8(c.Distance.Level)
	msg.WriteUint8(c.Distance.Percent)
	msg.WriteUint8(c.Shielding.Level)
	msg.WriteUint8(c.Shielding.Percent)
	msg.WriteUint8(c.Fishing.Level)
	msg.WriteUint8(c.Fishing.Percent)
}

func AddStats(msg *Message, c *game.Creature) {
	msg.WriteUint8(0xa0) // send player stats
	msg.WriteUint16(c.HealthNow)
	msg.WriteUint16(c.HealthMax)
	msg.WriteUint16(c.Cap)
	msg.WriteUint32(c.Combat.Experience)
	msg.WriteUint8(c.Combat.Level)
	msg.WriteUint8(c.Combat.Percent)
	msg.WriteUint16(c.ManaNow)
	msg.WriteUint16(c.ManaMax)
	msg.WriteUint8(c.Magic.Level)
	msg.WriteUint8(c.Magic.Percent)
}

func AddInventory(msg *Message, c *game.Creature) {
	msg.WriteUint8(game.SlotEmpty)
	msg.WriteUint8(game.SlotHead)

	msg.WriteUint8(game.SlotEmpty)
	msg.WriteUint8(game.SlotNecklace)

	msg.WriteUint8(game.SlotNotEmpty)
	msg.WriteUint8(game.SlotBackpack)
	msg.WriteUint16(0x7c4) // backpack

	msg.WriteUint8(game.SlotNotEmpty)
	msg.WriteUint8(game.SlotBody)
	msg.WriteUint16(0x9a8) // magic plate armor

	msg.WriteUint8(game.SlotEmpty)
	msg.WriteUint8(game.SlotShield)

	msg.WriteUint8(game.SlotNotEmpty)
	msg.WriteUint8(game.SlotWeapon)
	msg.WriteUint16(0x997) // crossbow

	msg.WriteUint8(game.SlotEmpty)
	msg.WriteUint8(game.SlotLegs)

	msg.WriteUint8(game.SlotEmpty)
	msg.WriteUint8(game.SlotFeet)

	msg.WriteUint8(game.SlotEmpty)
	msg.WriteUint8(game.SlotRing)

	msg.WriteUint8(game.SlotNotEmpty)
	msg.WriteUint8(game.SlotAmmo)
	msg.WriteUint16(0x9ef) // bolts
	msg.WriteUint8(33)     // count
}

func AddMapDescription(msg *Message, pos game.Position, m *game.Map) {
	fmt.Printf(":: AddMapDescription ")
	begin := time.Now()
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
					if tile != nil {
						if skip > 0 {
							msg.WriteUint8(skip)
							msg.WriteUint8(0xff)
							skip = 0
						}
						AddTile(msg, tile)
					} else {
						skip++
						if skip == 0xff {
							msg.WriteUint8(skip)
							msg.WriteUint8(0xff)
							skip = 0
						}
					}
				}
			}
		}
	} else { // underground

	}

	// Remainder
	msg.WriteUint8(skip)
	msg.WriteUint8(0xff)

	fmt.Printf("[%v]\n", time.Since(begin))
}

func AddPosition(msg *Message, pos game.Position) {
	msg.WriteUint16(pos.X)
	msg.WriteUint16(pos.Y)
	msg.WriteUint8(pos.Z)
}

func AddMagicEffect(msg *Message, pos game.Position, kind uint8) {
	msg.WriteUint8(0x83)
	AddPosition(msg, pos)
	msg.WriteUint8(kind)
}

func AddCreature(msg *Message, c *game.Creature) {
	msg.WriteUint16(0x61) // unknown creature
	msg.WriteUint32(0x00) // something about caching known creatures
	msg.WriteUint32(c.ID)
	msg.WriteString(c.Name)
	msg.WriteUint8((uint8)(c.HealthNow*100/c.HealthMax) + 1)
	msg.WriteUint8(c.Direction) // look dir
	msg.WriteUint8(c.Outfit.Type)
	msg.WriteUint8(c.Outfit.Head)
	msg.WriteUint8(c.Outfit.Body)
	msg.WriteUint8(c.Outfit.Legs)
	msg.WriteUint8(c.Outfit.Feet)
	msg.WriteUint8(c.Light.Level)
	msg.WriteUint8(c.Light.Color)
	msg.WriteUint16(c.Speed)
	msg.WriteUint8(c.Skull)
	msg.WriteUint8(c.Party)
}

func AddTile(msg *Message, tile *game.Tile) {
	for _, i := range tile.Items {
		msg.WriteUint16(i.ID)
	}
	for _, c := range tile.Creatures {
		AddCreature(msg, c)
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

func SendAddCreature(c *net.Conn, character *game.Creature, m *game.Map) {
	res := NewMessage()
	res.WriteUint8(0x0a)
	res.WriteUint32(character.ID) // ID
	res.WriteUint16(0x32)         // ?

	// can report bugs?
	if character.Access > game.Regular {
		res.WriteUint8(0x01)
	} else {
		res.WriteUint8(0x00)
	}
	if character.Access >= game.Gamemaster {
		res.WriteUint8(0x0b)
		for i := 0; i < 32; i++ {
			res.WriteUint8(0xff)
		}
	}

	tile := m.GetTile(character.Position)
	tile.AddCreature(character)
	AddMapDescription(res, character.Position, m)
	AddMagicEffect(res, character.Position, 0x0a)
	AddInventory(res, character)
	AddStats(res, character)
	AddSkills(res, character)
	AddWorldLight(res, &character.World)
	AddCreatureLight(res, character)
	AddInfoString(res, fmt.Sprintf("Welcome, %s.", character.Name))
	AddInfoString(res, "TODO: Last Login String 01-01-1970")
	AddCreatureLight(res, character)
	AddIcons(res, character)
	SendMessage(c, res)
	res.HexDump("response")
}

func AddInfoString(msg *Message, str string) {
	msg.WriteUint8(0xb4)
	msg.WriteUint8(0x15) // type info
	msg.WriteString(str)
}
