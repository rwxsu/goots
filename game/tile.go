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
