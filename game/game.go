package game

// Light has the same structure for both creatures and world
type Light struct {
	Level uint8
	Color uint8
}

// World represents a game world, displayed in character list. Every world has
// its own port.
type World struct {
	Name  string
	Port  uint16
	Light Light
}
