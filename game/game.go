package game

// Access
const (
	Regular uint8 = iota + 1
	Tutor
	SeniorTutor
	Gamemaster
	God
)

// Direction
const (
	North uint8 = iota
	East
	South
	West
)

// Slot
const (
	SlotHead uint8 = iota + 1
	SlotNecklace
	SlotBackpack
	SlotBody
	SlotShield
	SlotWeapon
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

type Position struct {
	X uint16
	Y uint16
	Z uint8
}

type Outfit struct {
	Type uint8
	Head uint8
	Body uint8
	Legs uint8
	Feet uint8
}

type Item struct {
	ID uint16
}

type World struct {
	Name string
	Port uint16
}

type Skill struct {
	Experience uint32 // tries
	Level      uint8
	Percent    uint8
}

type Creature struct {
	Position  Position
	Access    uint8
	World     World
	ID        uint32
	Name      string
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
}
