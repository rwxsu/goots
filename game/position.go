package game

import "fmt"

// Direction
const (
	North uint8 = iota
	East
	South
	West
)

// SectorPosition contains the position equal to the .sec filename
//
// To get the real in-game position, calculate:
//		x = SectorPosition.X * 32 + x offset
//		y = SectorPosition.Y * 32 + y offset
//		z = SectorPosition.Z
//
//		E.g.: 0999-0998-07.sec (02-04: Content={102})
//	 		x = 999 * 32 + 2
//	 		y = 998 * 32 + 4
//			z = 7
type SectorPosition Position

// Center returns a sector's real in-game center position
func Center(spos SectorPosition) Position {
	return Position{X: spos.X*32 + 15, Y: spos.Y*32 + 15, Z: spos.Z}
}

// Position is the real in-game position
type Position struct {
	X uint16
	Y uint16
	Z uint8
}

type Offset struct {
	X, Y, Z int8
}

// Offset the position by the given positive or negative offset
func (pos *Position) Offset(offset Offset) {
	pos.X += (uint16)(offset.X)
	pos.Y += (uint16)(offset.Y)
	pos.Z += (uint8)(offset.Z)
}

// Equals checks if two positions are the same.
func (pos Position) Equals(other Position) bool {
	return pos.X == other.X && pos.Y == other.Y && pos.Z == other.Z
}

func (pos Position) String() string {
	return fmt.Sprintf("X:%d, Y:%d, Z:%d", pos.X, pos.Y, pos.Z)
}
