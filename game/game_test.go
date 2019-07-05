package game

import (
	"testing"
)

func TestPositionOffset(t *testing.T) {
	tables := []struct {
		startPos Position
		offset   Offset
		endPos   Position
	}{
		{Position{X: 0, Y: 0, Z: 0}, Offset{X: 1, Y: 2, Z: 3}, Position{X: 1, Y: 2, Z: 3}},
		{Position{X: 1, Y: 2, Z: 3}, Offset{X: -1, Y: -2, Z: -3}, Position{X: 0, Y: 0, Z: 0}},
		{Position{X: 0, Y: 0, Z: 0}, Offset{X: -1, Y: -1, Z: -1}, Position{X: 65535, Y: 65535, Z: 255}},
	}

	for _, table := range tables {
		table.startPos.Offset(table.offset)
		if !table.startPos.Equals(table.endPos) {
			t.Errorf("Offset position was incorrect: %s. Expected: %s", table.startPos.String(), table.endPos.String())
		}
	}
}
