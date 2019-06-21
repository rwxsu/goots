package game

// World represents a game world, displayed in character list. Every world has
// its own port.
type World struct {
	Name  string
	Port  uint16
	Light Light
}
