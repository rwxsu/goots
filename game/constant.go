package game

// Account Type
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
	Head uint8 = iota + 1
	Necklace
	Backpack
	Body
	Shield
	Weapon
	Legs
	Ring
	Feet
	Ammo
)

// Slot capacity
const (
	NotEmpty uint8 = 0x78
	Empty    uint8 = 0x79
)
