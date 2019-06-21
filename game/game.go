package game

// Direction
const (
	North uint8 = iota
	East
	South
	West
)

// Position is the real in-game position
type Position struct {
	X uint16
	Y uint16
	Z uint8
}

// Light has the same structure for both creatures and world
type Light struct {
	Level uint8
	Color uint8
}
