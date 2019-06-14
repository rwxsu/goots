package game

type Health struct {
	Now uint16
	Max uint16
}

type Mana struct {
	Now uint16
	Max uint16
}

type Character struct {
	ID        uint32
	Access    uint8
	Name      string
	Cap       uint16
	Health    Health
	Mana      Mana
	Combat    Skill
	Magic     Skill
	Fist      Skill
	Club      Skill
	Sword     Skill
	Axe       Skill
	Distance  Skill
	Shielding Skill
	Fishing   Skill
	Direction uint8
	Position  Position
	Outfit    Outfit
	World     World
}
