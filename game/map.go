package game

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"time"

	"github.com/rwxsu/goot/parser"
)

type SectorPosition Position

type Column [32]*Tile
type Sector [32]*Column
type Map map[SectorPosition]*Sector

func (m *Map) SetTile(tile *Tile) {
	var spos SectorPosition
	spos.X = tile.X / 32
	spos.Y = tile.Y / 32
	spos.Z = tile.Z
	if (*m)[spos] == nil {
		return
	}
	(*m)[spos][tile.X%32][tile.Y%32] = tile
}

func (m *Map) GetTile(pos Position) *Tile {
	var spos SectorPosition
	spos.X = pos.X / 32
	spos.Y = pos.Y / 32
	spos.Z = pos.Z
	if (*m)[spos] == nil {
		return nil
	}
	return (*m)[spos][pos.X%32][pos.Y%32]
}

func (m *Map) GetSpectators(pos Position) []*Creature {
	pos.X = pos.X - 8
	pos.Y = pos.Y - 6
	var spectators []*Creature
	for x := (uint16)(0); x < 18; x++ {
		for y := (uint16)(0); y < 14; y++ {
			tile := m.GetTile(pos)
			if tile != nil {
				spectators = append(spectators, tile.Creatures...)
			}
		}
	}
	return spectators
}

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
			var tile Tile
			tile.SetPosition(spos.X*32+offsetX, spos.Y*32+offsetY, spos.Z)
			p.NextToken() // skip offsetX
			p.NextToken() // skip offsetY
			itemids := p.NextToken()
			switch itemids := itemids.(type) {
			case []int:
				for _, id := range itemids {
					tile.AddItem(&Item{ID: (uint16)(id)})
				}
				break
			default:
				panic("in LoadSector: could not get item ids")
			}
			(*m)[spos][offsetX][offsetY] = &tile
		}
	}

	fmt.Printf("[%v]\n", time.Since(begin))
}
