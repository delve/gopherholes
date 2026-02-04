package year2025

import "aocgen/pkg/common/tilemap"

type Day07 struct {
	manifold *tilemap.TileMap
}

func (p *Day07) parseInput(lines []string) {
	p.manifold = &tilemap.TileMap{}
	p.manifold.SetColumns(float64(len(lines[0])))
	for _, line := range lines[:len(lines)-1] {
		for _, floorContents := range line {
			mapTile := &tilemap.RuneTile{Value: floorContents}
			p.manifold.AddTile(mapTile)
		}
	}
}

func (p Day07) PartA(lines []string) any {
	p.parseInput(lines)
	splits := 0
	for y := 1.0; y <= p.manifold.MaxY(); y++ {
		for x := 0.0; x <= p.manifold.MaxX(); x++ {
			thisTile := p.manifold.GetTileAt(x, y)
			valueAbove := p.manifold.GetTileAt(x, y-1).GetValue()
			if valueAbove == '|' || valueAbove == 'S' {
				if thisTile.GetValue() == '.' {
					thisTile.SetValue('|')
				} else if thisTile.GetValue() == '^' {
					splits++
					p.manifold.GetTileAt(x-1, y).SetValue('|')
					p.manifold.GetTileAt(x+1, y).SetValue('|')
				}
			}
		}
	}
	// println(p.manifold.String())

	if len(lines[0]) <= 18 && splits == 21 {
		println("Correct")
	}
	return splits
}

func (p Day07) PartB(lines []string) any {
	p.parseInput(lines)
	return "implement_me"
}
