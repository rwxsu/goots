package game

import (
	"testing"
)

type TestMap struct {
	m *Map
	c *Creature
}

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
	tile := m.GetTile(Position{X: spos.X*32 + 15, Y: spos.Y*32 + 15, Z: spos.Z})
	if len(tile.Creatures) == 0 {
		t.Error("creature not found at sector center")
	}
}

func TestMoveCreature(t *testing.T) {
	m := make(Map)
	c := Creature{ID: 1}
	spos := SectorPosition{X: 1, Y: 1, Z: 1}
	center := Position{spos.X*32 + 15, spos.Y*32 + 15, spos.Z}
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
