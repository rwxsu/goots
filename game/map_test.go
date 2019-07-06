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
			tile := m.GetTile(Position{spos.X*32 + offsetX, spos.Y*32 + offsetY, spos.Z})
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
	c := Creature{ID: 1}
	spos := SectorPosition{X: 1, Y: 1, Z: 1}
	m.InitializeSector(spos, 104)
	m.AddCreatureToSectorCenter(spos, &c)
	tile := m.GetTile(Center(spos))
	if len(tile.Creatures) == 0 {
		t.Error("creature not found at sector center")
	}
}

func TestMoveCreature(t *testing.T) {
	m := make(Map)
	c := Creature{ID: 1}
	spos := SectorPosition{X: 1, Y: 1, Z: 1}
	center := Center(spos)
	m.InitializeSector(spos, 104)
	m.AddCreatureToSectorCenter(spos, &c)
	if !m.MoveCreature(&c, Position{c.X, c.Y - 1, c.Z}, North) {
		t.Error("could not move creature to the north")
	}
	if !m.MoveCreature(&c, Position{c.X + 1, c.Y, c.Z}, East) {
		t.Error("could not move creature to the east")
	}
	if !m.MoveCreature(&c, Position{c.X, c.Y + 1, c.Z}, South) {
		t.Error("could not move creature to the south")
	}
	if !m.MoveCreature(&c, Position{c.X - 1, c.Y, c.Z}, West) {
		t.Error("could not move creature to the west")
	}
	if c.Position != center {
		t.Error("creature did not returned to center")
	}
}
