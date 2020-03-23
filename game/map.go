package game

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"time"

	"github.com/maksumic/goot/parser"
)

// Column consists of 32 tiles (stacked vertically)
type Column [32]*Tile

// Sector consists of 32*32=256 tiles (32 columns, stacked horizontally)
type Sector [32]*Column

// Map is a collection of sectors loaded from data/map/sectors/*.sec
type Map map[SectorPosition]*Sector

// Tile gets the SectorPosition by dividing the real x and y position by 32.
// Mod is used to get the remainder, which is the x and y offset.
func (m *Map) Tile(pos Position) *Tile {
	var spos SectorPosition
	spos.X = pos.X / 32
	spos.Y = pos.Y / 32
	spos.Z = pos.Z
	if (*m)[spos] == nil {
		return nil
	}
	return (*m)[spos][pos.X%32][pos.Y%32]
}

// LoadSector loads a single .sec file and stores the sector at the correct
// SectorPosition extracted from the filename in the map.
func (m *Map) LoadSector(filename string) {
	var p parser.Parser
	p.Filename = filename
	dirOffset := len(filepath.Dir(p.Filename)) + 1
	x, _ := strconv.Atoi(p.Filename[dirOffset+0 : dirOffset+4])
	y, _ := strconv.Atoi(p.Filename[dirOffset+5 : dirOffset+9])
	z, _ := strconv.Atoi(p.Filename[dirOffset+10 : dirOffset+12])
	fmt.Printf(":: Loading %04d-%04d-%02d.sec ", x, y, z)
	begin := time.Now()

	if fileBytes, err := ioutil.ReadFile(p.Filename); err == nil {
		p.Buffer = bytes.NewBuffer(fileBytes)
	} else {
		panic(err)
	}

	spos := SectorPosition{X: (uint16)(x), Y: (uint16)(y), Z: (uint8)(z)}
	(*m)[spos] = new(Sector)
	for offsetX := (uint16)(0); offsetX < 32; offsetX++ {
		(*m)[spos][offsetX] = new(Column)
		for offsetY := (uint16)(0); offsetY < 32; offsetY++ {
			tile := NewTile(Position{X: spos.X*32 + offsetX, Y: spos.Y*32 + offsetY, Z: spos.Z})
			p.NextToken() // skip offsetX
			p.NextToken() // skip offsetY
			itemids := p.NextToken()
			switch itemids := itemids.(type) {
			case []int:
				for _, id := range itemids {
					tile.AddItem(&Item{ID: (uint16)(id)})
				}
			default:
				panic("in LoadSector: could not get item ids")
			}
			if len(tile.Items) > 0 {
				(*m)[spos][offsetX][offsetY] = tile
			} else {
				(*m)[spos][offsetX][offsetY] = nil
			}
		}
	}
	fmt.Printf("[%v]\n", time.Since(begin))
}

// MovePlayer on map
func (m *Map) MovePlayer(p *Player, pos Position, direction uint8) bool {
	from := m.Tile(p.Position())
	to := m.Tile(pos)
	if from == nil || to == nil {
		return false
	}
	from.RemovePlayer(p)
	to.AddPlayer(p)
	p.SetPosition(pos)
	p.SetDirection(direction)
	return true
}
