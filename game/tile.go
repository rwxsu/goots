package game

func NewTile(pos Position) *Tile {
	return &Tile{
		Position:  pos,
		Players: make(map[uint32]*Player),
	}
}

type Tile struct {
	Position
	Items     []*Item
	Players map[uint32]*Player
}

func (t *Tile) AddItem(i *Item) {
	t.Items = append(t.Items, i)
}

func (t *Tile) AddPlayer(p *Player) {
	t.Players[p.ID()] = p
}

func (t *Tile) RemovePlayer(p *Player) {
	delete(t.Players, p.ID())
}
