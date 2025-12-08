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
	print(p.world.String())
	movableJunk := 0

	for y := 0.0; y < p.world.MaxY(); y++ {
		for x := 0.0; x < p.world.MaxX(); x++ {
			if p.world.TileAt(x, y).GetValue() == '@' {
				neighbors := p.world.CountAround(p.world.TileAt(x, y), x, y)
				p.world.TileAt(x, y).SetValue(rune(fmt.Sprintf("%d", neighbors)[0]))
				if neighbors < 4 {
					movableJunk++
				}
			}
		}
	}
	print(p.world.String())
	return movableJunk
}

func (p Day04) PartB(lines []string) any {
	return "implement_me"
}
