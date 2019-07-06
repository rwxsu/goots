package game

// Access
const (
	Regular uint8 = iota + 1
	Tutor
	SeniorTutor
	Gamemaster
	God
)

// Player message type
const (
	PlayerMessageTypeInfo   uint8 = 0x15
	PlayerMessageTypeCancel uint8 = 0x17
)

// Slot
const (
	SlotHead uint8 = iota + 1
	SlotNecklace
	SlotBackpack
	SlotBody
	SlotShield // Right hand
	SlotWeapon // Left hand
	SlotLegs
	SlotRing
	SlotFeet
	SlotAmmo
)

// Slot Capacity
const (
	SlotEmpty    uint8 = 0x79
	SlotNotEmpty uint8 = 0x78
)

// Tactic contains the user selected options right of equipment slots in client
type Tactic struct {
	// Fight modes
	// 1 - Offensive fighting
	// 2 - Balanced fighting
	// 3 - Defensive fighting
	FightMode uint8

	// 0 - Stand while fighting
	// 1 - Chase opponent
	ChaseOpponent uint8

	// 0 - You cannot attack unmarked players
	// 1 - You can attack any player
	AttackPlayers uint8
}

type Skill struct {
	Experience uint32 // tries
	Level      uint8
	Percent    uint8
}

type Outfit struct {
	Type uint8
	Head uint8
	Body uint8
	Legs uint8
	Feet uint8
}

func NewPlayer(id uint32, name string, pos Position) *Creature {
	return &Creature{
		Access:    Regular,
		World:     World{Name: "test", Port: 7171, Light: Light{Color: 0x7, Level: 0xd7}},
		ID:        id,
		Name:      name,
		Position:  pos,
		Direction: South,
		Outfit:    Outfit{Type: 0x80, Head: 0x50, Body: 0x50, Legs: 0x50, Feet: 0x50},
		Speed:     60000,
		Skull:     0,
		Party:     0,
		Cap:       100,
		HealthNow: 50,
		HealthMax: 100,
		ManaNow:   50,
		ManaMax:   100,
		Combat:    Skill{Experience: 4200, Level: 8, Percent: 0},
		Magic:     Skill{Level: 10, Percent: 50},
		Fist:      Skill{Level: 10, Percent: 50},
		Club:      Skill{Level: 10, Percent: 50},
		Sword:     Skill{Level: 10, Percent: 50},
		Axe:       Skill{Level: 10, Percent: 50},
		Distance:  Skill{Level: 10, Percent: 50},
		Shielding: Skill{Level: 10, Percent: 50},
		Fishing:   Skill{Level: 10, Percent: 50},
		Light:     Light{Color: 0xd7, Level: 0xd7},
		Icons:     0,
		Tactic:    Tactic{FightMode: 0, ChaseOpponent: 0, AttackPlayers: 0},
	}
}

type Creature struct {
	Access    uint8
	World     World
	ID        uint32
	Name      string
	Position  Position
	Direction uint8
	Outfit    Outfit
	Speed     uint16
	Skull     uint8
	Party     uint8
	Cap       uint16
	HealthNow uint16
	HealthMax uint16
	ManaNow   uint16
	ManaMax   uint16
	Combat    Skill
	Magic     Skill
	Fist      Skill
	Club      Skill
	Sword     Skill
	Axe       Skill
	Distance  Skill
	Shielding Skill
	Fishing   Skill
	Light     Light
	Icons     uint8
	Tactic    Tactic
}

type Player struct {
	Creature
}
