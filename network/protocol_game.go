package network

import (
	"fmt"
	"net"

	"github.com/maksumic/goot/game"
)

func ParseCommand(c net.Conn, m *game.Map, p *game.Player, msg *Message, code uint8) {
	switch code {
	case 0x65:
		SendMoveCreature(c, m, p, game.North, code)
	case 0x66:
		SendMoveCreature(c, m, p, game.East, code)
	case 0x67:
		SendMoveCreature(c, m, p, game.South, code)
	case 0x68:
		SendMoveCreature(c, m, p, game.West, code)
	case 0x6f:
		SendTurnCreature(c, p, game.North)
	case 0x70:
		SendTurnCreature(c, p, game.East)
	case 0x71:
		SendTurnCreature(c, p, game.South)
	case 0x72:
		SendTurnCreature(c, p, game.West)
	case 0xa0:
		p.Tactic.FightMode = msg.ReadUint8()
		p.Tactic.ChaseOpponent = msg.ReadUint8()
		p.Tactic.AttackPlayers = msg.ReadUint8()
	default:
		SendSnapback(c, p)
	}
}

func SendSnapback(c net.Conn, p *game.Player) {
	msg := NewMessage()
	msg.WriteUint8(0xb5)
	msg.WriteUint8(p.Direction())
	SendMessage(c, msg)
	SendCancelMessage(c, "Sorry, not possible.")
}

func SendCancelMessage(c net.Conn, str string) {
	msg := NewMessage()
	AddPlayerMessage(msg, str, game.PlayerMessageTypeCancel)
	SendMessage(c, msg)
}

func SendMoveCreature(c net.Conn, m *game.Map, p *game.Player, direction, code uint8) {
	var offset game.Offset
	var width, height uint16
	from := p.Position()
	to := p.Position()
	switch direction {
	case game.North:
		offset.X = -8
		offset.Y = -6
		width = 18
		height = 1
		to.Y--
	case game.South:
		offset.X = -8
		offset.Y = 7
		width = 18
		height = 1
		to.Y++
	case game.East:
		offset.X = 9
		offset.Y = -6
		width = 1
		height = 14
		to.X++
	case game.West:
		offset.X = -8
		offset.Y = -6
		width = 1
		height = 14
		to.X--
	}
	if !m.MoveCreature(p, to, direction) {
		SendSnapback(c, p)
		return
	}
	msg := NewMessage()
	msg.WriteUint8(0x6d)
	AddPosition(msg, from)
	msg.WriteUint8(0x01) // oldStackPos
	AddPosition(msg, to)
	msg.WriteUint8(code)
	AddMapArea(msg, m, to, offset, width, height)
	SendMessage(c, msg)
}

func SendTurnCreature(c net.Conn, p *game.Player, direction uint8) {
	p.SetDirection(direction)
	msg := NewMessage()
	msg.WriteUint8(0x6b)
	AddPosition(msg, p.Position())
	msg.WriteUint8(1)
	msg.WriteUint16(0x63)
	msg.WriteUint32(p.ID())
	msg.WriteUint8(p.Direction())
	SendMessage(c, msg)
}

func SendAddCreature(c net.Conn, m *game.Map, p *game.Player) {
	res := NewMessage()
	res.WriteUint8(0x0a)
	res.WriteUint32(p.ID()) // ID
	res.WriteUint16(0x32)   // ?
	// can report bugs?
	if p.Access > game.Regular {
		res.WriteUint8(0x01)
	} else {
		res.WriteUint8(0x00)
	}
	if p.Access >= game.Gamemaster {
		res.WriteUint8(0x0b)
		for i := 0; i < 32; i++ {
			res.WriteUint8(0xff)
		}
	}
	tile := m.Tile(p.Position())
	tile.AddCreature(p)
	res.WriteUint8(0x64)
	AddPosition(res, p.Position())
	AddMapArea(res, m, p.Position(), game.Offset{X: -8, Y: -6, Z: 0}, 18, 14)
	AddMagicEffect(res, p.Position(), 0x0a)
	AddInventory(res, p)
	AddStats(res, p)
	AddSkills(res, p)
	AddWorldLight(res, p.World)
	AddCreatureLight(res, p)
	AddPlayerMessage(res, fmt.Sprintf("Welcome, %s.", p.Name()), game.PlayerMessageTypeInfo)
	AddPlayerMessage(res, "TODO: Last Login String 01-01-1970", game.PlayerMessageTypeInfo)
	AddCreatureLight(res, p)
	AddIcons(res, p)
	SendMessage(c, res)
}

func AddCreatureLight(msg *Message, c game.Creature) {
	msg.WriteUint8(0x8d)
	msg.WriteUint32(c.ID())
	msg.WriteUint8(c.Light().Level)
	msg.WriteUint8(c.Light().Color)
}

func AddWorldLight(msg *Message, w game.World) {
	msg.WriteUint8(0x82)
	msg.WriteUint8(w.Light.Level)
	msg.WriteUint8(w.Light.Color)
}

func AddIcons(msg *Message, p *game.Player) {
	msg.WriteUint8(0xa2)
	msg.WriteUint8(p.Icons)
}

func AddSkills(msg *Message, p *game.Player) {
	msg.WriteUint8(0xa1)
	msg.WriteUint8(p.Fist.Level)
	msg.WriteUint8(p.Fist.Percent)
	msg.WriteUint8(p.Club.Level)
	msg.WriteUint8(p.Club.Percent)
	msg.WriteUint8(p.Sword.Level)
	msg.WriteUint8(p.Sword.Percent)
	msg.WriteUint8(p.Axe.Level)
	msg.WriteUint8(p.Axe.Percent)
	msg.WriteUint8(p.Distance.Level)
	msg.WriteUint8(p.Distance.Percent)
	msg.WriteUint8(p.Shielding.Level)
	msg.WriteUint8(p.Shielding.Percent)
	msg.WriteUint8(p.Fishing.Level)
	msg.WriteUint8(p.Fishing.Percent)
}

func AddStats(msg *Message, p *game.Player) {
	msg.WriteUint8(0xa0) // send player stats
	msg.WriteUint16(p.HealthNow())
	msg.WriteUint16(p.HealthMax())
	msg.WriteUint16(p.Cap)
	msg.WriteUint32(p.Combat.Experience)
	msg.WriteUint8(p.Combat.Level)
	msg.WriteUint8(p.Combat.Percent)
	msg.WriteUint16(p.ManaNow)
	msg.WriteUint16(p.ManaMax)
	msg.WriteUint8(p.Magic.Level)
	msg.WriteUint8(p.Magic.Percent)
}

func AddInventory(msg *Message, p *game.Player) {
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

// AddMapArea adds the area starting at position+offset until width and height
// is reached. Returns the number of tiles (counting nil)
func AddMapArea(msg *Message, m *game.Map, pos game.Position, offset game.Offset, width, height uint16) int {
	pos.Offset(offset)
	count := 0
	skip := (uint8)(0)
	if pos.Z < 8 {
		for z := (int8)(7); z > -1; z-- {
			for x := (uint16)(0); x < width; x++ {
				for y := (uint16)(0); y < height; y++ {
					tile := m.Tile(game.Position{X: pos.X + x, Y: pos.Y + y, Z: (uint8)(z)})
					if tile != nil {
						if skip > 0 {
							msg.WriteUint8(skip - 1)
							msg.WriteUint8(0xff)
							skip = 0
						}
						AddTile(msg, tile)
					}
					skip++
					if skip == 0xff {
						msg.WriteUint8(0xff)
						msg.WriteUint8(0xff)
						skip = 0
					}
					count++
				}
			}
		}
	} else { // TODO: underground

	}
	// Remainder
	if skip > 0 {
		msg.WriteUint8(skip - 1)
		msg.WriteUint8(0xff)
		skip = 0
	}
	return count
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

func AddCreature(msg *Message, c game.Creature) {
	msg.WriteUint16(0x61) // unknown creature
	msg.WriteUint32(0x00) // something about caching known creatures
	msg.WriteUint32(c.ID())
	msg.WriteString(c.Name())
	msg.WriteUint8(game.HealthPercent(c))
	msg.WriteUint8(c.Direction())
	msg.WriteUint8(c.Outfit().Type)
	msg.WriteUint8(c.Outfit().Head)
	msg.WriteUint8(c.Outfit().Body)
	msg.WriteUint8(c.Outfit().Legs)
	msg.WriteUint8(c.Outfit().Feet)
	msg.WriteUint8(c.Light().Level)
	msg.WriteUint8(c.Light().Color)
	msg.WriteUint16(c.Speed())
	msg.WriteUint8(c.Skull())
	msg.WriteUint8(c.Party())
}

// AddTile adds all the tile items (including ground) and creatures WITHOUT the end of tile
// delimeter (0xSKIPCOUNT-0xff)
func AddTile(msg *Message, tile *game.Tile) {
	for _, i := range tile.Items {
		msg.WriteUint16(i.ID)
	}
	for _, c := range tile.Creatures {
		AddCreature(msg, c)
	}
}

func AddPlayerMessage(msg *Message, str string, kind uint8) {
	msg.WriteUint8(0xb4)
	msg.WriteUint8(kind)
	msg.WriteString(str)
}
