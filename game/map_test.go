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
	spos := SectorPosition{X: 1000, Y: 1000, Z: 7}
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

func NewTestMap(t *testing.T) *TestMap {
	// Setup a 32x32, single-level map.
	spos := SectorPosition{X: 1, Y: 1, Z: 1}
	m := make(Map)
	m.InitializeSector(spos, 104)
	// Add a creature to the center of the map.
	centerPos := Position{X: (spos.X * 32) + 15, Y: (spos.Y * 32) + 15, Z: 1}
	creature := Creature{ID: 1234, Position: centerPos}
	tile := m.GetTile(centerPos)
	if tile == nil {
		t.Errorf("No tile found at: %s", centerPos.String())
	}
	tile.AddCreature(&creature)
	return &TestMap{&m, &creature}
}

func TestMoveCreatureNorth(t *testing.T) {
	testMap := NewTestMap(t)
	endPos := Position{X: testMap.c.Position.X, Y: testMap.c.Position.Y - 1, Z: testMap.c.Position.Z}
	if !testMap.m.MoveCreature(testMap.c, endPos, North) {
		t.Error("Failed to move creature.")
	}
}

func TestMoveCreatureEast(t *testing.T) {
	testMap := NewTestMap(t)
	endPos := Position{X: testMap.c.Position.X + 1, Y: testMap.c.Position.Y, Z: testMap.c.Position.Z}
	if !testMap.m.MoveCreature(testMap.c, endPos, East) {
		t.Error("Failed to move creature.")
	}
}

func TestMoveCreatureSouth(t *testing.T) {
	testMap := NewTestMap(t)
	endPos := Position{X: testMap.c.Position.X, Y: testMap.c.Position.Y + 1, Z: testMap.c.Position.Z}
	if !testMap.m.MoveCreature(testMap.c, endPos, South) {
		t.Error("Failed to move creature.")
	}
}

func TestMoveCreatureWest(t *testing.T) {
	testMap := NewTestMap(t)
	endPos := Position{X: testMap.c.Position.X - 1, Y: testMap.c.Position.Y, Z: testMap.c.Position.Z}
	if !testMap.m.MoveCreature(testMap.c, endPos, West) {
		t.Error("Failed to move creature.")
	}
}
