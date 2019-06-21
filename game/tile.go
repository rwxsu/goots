package game

const (
	// EOT End Of Tile
	EOT uint8 = 0xff
)

type Tile struct {
	Position
	Items     []*Item
	Creatures []*Creature
}

func (t *Tile) SetPosition(x, y uint16, z uint8) {
	t.X = x
	t.Y = y
	t.Z = z
}

func (t *Tile) AddItem(i *Item) {
	t.Items = append(t.Items, i)
}

func (t *Tile) AddCreature(c *Creature) {
	t.Creatures = append(t.Creatures, c)
}

// RemoveCreature compares the ID of the creatures on the tile with the given
// creature's ID
func (t *Tile) RemoveCreature(c *Creature) bool {
	for i, tc := range t.Creatures {
		if c.ID == tc.ID {
			// Swap ith elementh with last elementh
			t.Creatures[i] = t.Creatures[len(t.Creatures)-1]
			// Decrease slice length by 1
			t.Creatures = t.Creatures[:len(t.Creatures)-1]
			return true
		}
	}
	return false
}
