package game

import (
	"testing"
)

func TestInitializeSector(t *testing.T) {
	m := make(Map)
	spos := SectorPosition{X: 1, Y: 1, Z: 1}
	m.InitializeSector(spos, 104)
	count := 0
	for offsetX := (uint16)(0); offsetX < 32; offsetX++ {
		for offsetY := (uint16)(0); offsetY < 32; offsetY++ {
			tile := m.Tile(Position{spos.X*32 + offsetX, spos.Y*32 + offsetY, spos.Z})
			if tile == nil {
				count++
			}
		}
	}
	if count > 0 {
		t.Errorf("%d tiles not initialized", count)
	}
}

func TestAddCreatureToSectorCenter(t *testing.T) {
	m := make(Map)
	var p Player
	p.SetID(1)
	spos := SectorPosition{X: 1, Y: 1, Z: 1}
	m.InitializeSector(spos, 104)
	m.AddCreatureToSectorCenter(spos, &p)
	tile := m.Tile(Center(spos))
	if len(tile.Creatures) == 0 {
		t.Error("creature not found at sector center")
	}
}

func TestMoveCreature(t *testing.T) {
	m := make(Map)
	var p Player
	p.SetID(1)
	spos := SectorPosition{X: 1, Y: 1, Z: 1}
	m.InitializeSector(spos, 104)
	m.AddCreatureToSectorCenter(spos, &p)
	if !m.MoveCreature(&p, Position{p.Position().X, p.Position().Y - 1, p.Position().Z}, North) {
		t.Error("could not move creature to the north")
	}
	if !m.MoveCreature(&p, Position{p.Position().X + 1, p.Position().Y, p.Position().Z}, East) {
		t.Error("could not move creature to the east")
	}
	if !m.MoveCreature(&p, Position{p.Position().X, p.Position().Y + 1, p.Position().Z}, South) {
		t.Error("could not move creature to the south")
	}
	if !m.MoveCreature(&p, Position{p.Position().X - 1, p.Position().Y, p.Position().Z}, West) {
		t.Error("could not move creature to the west")
	}
	if !p.Position().Equals(Center(spos)) {
		t.Errorf("creature did not returned to center: expected {%s} got {%s}", Center(spos).String(), p.Position().String())
	}
}
