package game

// Creature is a common interface shared between player and monster
type Creature interface {

	// ID has to be unique
	SetID(uint32)
	ID() uint32

	SetName(string)
	Name() string

	SetPosition(Position)
	Position() Position

	SetDirection(uint8)
	Direction() uint8

	SetOutfit(Outfit)
	Outfit() Outfit

	SetHealthNow(uint16)
	SetHealthMax(uint16)
	HealthNow() uint16
	HealthMax() uint16

	SetSpeed(uint16)
	Speed() uint16

	SetSkull(uint8)
	Skull() uint8

	SetParty(uint8)
	Party() uint8

	SetLight(Light)
	Light() Light
}

func HealthPercent(c Creature) uint8 {
	return (uint8)((c.HealthNow() * 100 / c.HealthMax()))
}
