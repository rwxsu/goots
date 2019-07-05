package game

import (
	"testing"
)

func TestRemoveCreature(t *testing.T) {
	creature := Creature{ID: 1234}

	var tile Tile
	tile.AddCreature(&creature)

	if !tile.RemoveCreature(&creature) {
		t.Error("Failed to remove creature from tile.")
	}
}
