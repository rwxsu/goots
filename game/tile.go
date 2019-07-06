package game

func NewTile(pos Position) *Tile {
	return &Tile{
		Position:  pos,
		Creatures: make(map[uint32]Creature),
	}
}

type Tile struct {
	Position
	Items     []*Item
	Creatures map[uint32]Creature
}

func (t *Tile) AddItem(i *Item) {
	t.Items = append(t.Items, i)
}

func (t *Tile) FindCreature(c Creature) bool {
	return t.Creatures[c.ID()] != nil
}

func (t *Tile) AddCreature(c Creature) {
	t.Creatures[c.ID()] = c
}

func (t *Tile) RemoveCreature(c Creature) {
	delete(t.Creatures, c.ID())
}
