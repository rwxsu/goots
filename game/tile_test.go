package game

import "testing"

func TestRemoveCreature(t *testing.T) {
	m := make(Map)
	spos := SectorPosition{X: 1, Y: 1, Z: 1}
	p0 := NewPlayer(0, "test", Center(spos))
	p1 := NewPlayer(1, "test", Center(spos)) // creature too remove
	p2 := NewPlayer(2, "test", Center(spos))
	m.InitializeSector(spos, 104)
	m.AddCreatureToSectorCenter(spos, p0)
	m.AddCreatureToSectorCenter(spos, p1)
	m.AddCreatureToSectorCenter(spos, p2)
	tile := m.Tile(Center(spos))
	if !tile.FindCreature(p1) {
		t.Error("could not find creature on tile")
	}
	tile.RemoveCreature(p1)
	if tile.FindCreature(p1) {
		t.Error("could not remove creature from tile")
	}
	if !tile.FindCreature(p0) {
		t.Error("removed wrong creature from tile")
	}
	if !tile.FindCreature(p2) {
		t.Error("removed wrong creature from tile")
	}
	if len(tile.Creatures) == 0 {
		t.Error("removed all creatures from tile")
	}
}
