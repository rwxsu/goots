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

func NewPlayer(ID uint32, name string, pos Position) *Player {
	return &Player{
		Access:    Regular,
		World:     World{Name: "test", Port: 7171, Light: Light{Color: 0x7, Level: 0xd7}},
		id:        ID,
		name:      name,
		position:  pos,
		direction: South,
		outfit:    Outfit{Type: 0x80, Head: 0x50, Body: 0x50, Legs: 0x50, Feet: 0x50},
		speed:     600,
		skull:     0,
		party:     0,
		Cap:       100,
		healthNow: 50,
		healthMax: 100,
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
		light:     Light{Color: 0xd7, Level: 0xd7},
		Icons:     0,
		Tactic:    Tactic{FightMode: 0, ChaseOpponent: 0, AttackPlayers: 0},
	}
}

type Player struct {
	Access    uint8
	World     World
	id        uint32
	name      string
	position  Position
	direction uint8
	outfit    Outfit
	speed     uint16
	skull     uint8
	party     uint8
	Cap       uint16
	healthNow uint16
	healthMax uint16
	ManaNow   uint16
	ManaMax   uint16
	Combat    Skill // Regular level
	Magic     Skill
	Fist      Skill
	Club      Skill
	Sword     Skill
	Axe       Skill
	Distance  Skill
	Shielding Skill
	Fishing   Skill
	light     Light
	Icons     uint8
	Tactic    Tactic
}

func (p *Player) SetID(id uint32) {
	p.id = id
}

func (p *Player) ID() uint32 {
	return p.id
}

func (p *Player) SetName(name string) {
	p.name = name
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) SetPosition(pos Position) {
	p.position = pos
}

func (p *Player) Position() Position {
	return p.position
}

func (p *Player) SetDirection(dir uint8) {
	p.direction = dir
}

func (p *Player) Direction() uint8 {
	return p.direction
}

func (p *Player) SetOutfit(o Outfit) {
	p.outfit = o
}

func (p *Player) Outfit() Outfit {
	return p.outfit
}

func (p *Player) SetHealthNow(hp uint16) {
	p.healthNow = hp
}

func (p *Player) SetHealthMax(hp uint16) {
	p.healthMax = hp
}

func (p *Player) HealthNow() uint16 {
	return p.healthNow
}

func (p *Player) HealthMax() uint16 {
	return p.healthMax
}

func (p *Player) SetSpeed(speed uint16) {
	p.speed = speed
}

func (p *Player) Speed() uint16 {
	return p.speed
}

func (p *Player) SetSkull(skull uint8) {
	p.skull = skull
}

func (p *Player) Skull() uint8 {
	return p.skull
}

func (p *Player) SetParty(party uint8) {
	p.party = party
}

func (p *Player) Party() uint8 {
	return p.party
}

func (p *Player) SetLight(l Light) {
	p.light = l
}

func (p *Player) Light() Light {
	return p.light
}

func (p *Player) HealthPercent() uint8 {
	return (uint8)((p.HealthNow() * 100 / p.HealthMax()))
}
