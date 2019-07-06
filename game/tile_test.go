package game

import "testing"

func TestRemoveCreature(t *testing.T) {
	m := make(Map)
	spos := SectorPosition{X: 1, Y: 1, Z: 1}
	p := NewPlayer(1, "test", Center(spos))
	m.InitializeSector(spos, 104)
	m.AddCreatureToSectorCenter(spos, p)
	tile := m.Tile(Center(spos))
	if !tile.FindCreature(p) {
		t.Error("could not find creature on tile")
	}
	tile.RemoveCreature(p)
	if tile.FindCreature(p) {
		t.Error("could not remove creature from tile")
	}
}
