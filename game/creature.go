package game

// Access
const (
	Regular uint8 = iota + 1
	Tutor
	SeniorTutor
	Gamemaster
	God
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
	Light     Light
	Icons     uint8
}
