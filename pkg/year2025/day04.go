package year2025

import (
	"aocgen/pkg/common/tilemap"
	"fmt"
)

type Day04 struct {
	world *tilemap.TileMap
}

func (p *Day04) parseInput(lines []string) {
	p.world = &tilemap.TileMap{}
	p.world.SetColumns(float64(len(lines[0])))
	for _, line := range lines[:len(lines)-1] {
		for _, floorContents := range line {
			mapTile := &tilemap.RuneTile{Value: floorContents}
			p.world.AddTile(mapTile)
		}
	}
}

func (p Day04) PartA(lines []string) any {
	p.parseInput(lines)
	// print(p.world.String())
	moveableJunk := 0

	for y := 0.0; y <= p.world.MaxY(); y++ {
		for x := 0.0; x <= p.world.MaxX(); x++ {
			thisTile := p.world.GetTileAt(x, y)
			if thisTile.GetValue() == '@' {
				neighbors := p.world.CountAround(func(t tilemap.Tile) bool { return t.GetValue() != '.' }, x, y)
				p.world.GetTileAt(x, y).SetValue(rune(fmt.Sprintf("%d", neighbors)[0]))
				if neighbors < 4 {
					moveableJunk++
				}
			}
		}
	}
	// print(p.world.String())
	return moveableJunk
}

func getMoveableJunk(tm *tilemap.TileMap) []complex128 {
	moveableJunk := []complex128{}
	for y := 0.0; y <= tm.MaxY(); y++ {
		for x := 0.0; x <= tm.MaxX(); x++ {
			thisTile := tm.GetTileAt(x, y)
			if thisTile.GetValue() != '.' {
				neighbors := tm.CountAround(func(t tilemap.Tile) bool { return t.GetValue() != '.' }, x, y)
				tm.GetTileAt(x, y).SetValue(rune(fmt.Sprintf("%d", neighbors)[0]))
				if neighbors < 4 {
					moveableJunk = append(moveableJunk, thisTile.GetPosition())
				}
			}
		}
	}
	return moveableJunk
}

func (p Day04) PartB(lines []string) any {
	p.parseInput(lines)
	// print(p.world.String())
	movedJunk := 0
	for moveableJunk := getMoveableJunk(p.world); len(moveableJunk) > 0; moveableJunk = getMoveableJunk(p.world) {
		fmt.Printf("Moving %d items\n", len(moveableJunk))
		movedJunk += len(moveableJunk)
		for _, coord := range moveableJunk {
			p.world.GetTileAt(tilemap.SplitCoord(coord)).SetValue('.')
		}
		// print(p.world.String())
	}

	// print(p.world.String())
	return movedJunk
}
