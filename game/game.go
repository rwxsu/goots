package game

// Direction
const (
	North uint8 = iota
	East
	South
	West
)

type Position struct {
	X uint16
	Y uint16
	Z uint8
}

type Light struct {
	Level uint8
	Color uint8
}
