package tilemap

import "fmt"

// type TileMapper interface {
// 	MaxX() float64
// 	MaxY() float64
// 	SetColumns(float64)
// 	SetArbitraryTile(Tile)
// 	AddTile(Tile)
// 	String() string
// 	GetTileAt(float64, float64) Tile
// 	CountAround(TileFilter, float64, float64) int
// }

type TileFilter func(Tile) bool

type Tile interface {
	SetPosition(float64, float64)
	// SetPosition(complex128) >:|
	GetPosition() complex128
	X() float64
	Y() float64
	GetValue() any
	SetValue(any)
	String() string
	Rune() rune
}

type TileMap struct {
	tileData map[complex128]Tile
	// Maximum X and Y index value, 0 based
	boundary complex128
	// Position of last tile added, for linebreak behavior
	lastTile complex128
}

func SplitCoord(c complex128) (float64, float64) {
	return real(c), imag(c)
}

func (tm *TileMap) init() {
	tm.tileData = map[complex128]Tile{}
	tm.lastTile = complex(-1, 0)
}

func (tm *TileMap) setBoundaries(x, y float64) {
	tm.boundary = complex(x, y)
}

// Returns the largest 0 based X (column) index
func (tm TileMap) MaxX() float64 {
	return real(tm.boundary)
}

// Returns the largest 0 based Y (row) index
func (tm TileMap) MaxY() float64 {
	return imag(tm.boundary)
}

// Set the number of columns (1 based)
func (tm *TileMap) SetColumns(maxX float64) {
	tm.setBoundaries(maxX-1, tm.MaxY())
}

// Set an arbitrary tile. t.position is used to update map extents and tile data
func (tm *TileMap) SetArbitraryTile(t Tile) {
	if tm.tileData == nil {
		tm.init()
	}
	if t.X() > tm.MaxX() {
		tm.setBoundaries(t.X(), tm.MaxY())
	}
	if t.Y() > tm.MaxY() {
		tm.setBoundaries(tm.MaxX(), t.Y())
	}

	tm.tileData[t.GetPosition()] = t
}

func (tm *TileMap) getNextTilePosition() (float64, float64) {
	next := tm.lastTile
	if real(next) == tm.MaxX() {
		next = complex(0, imag(next)+1)
		tm.boundary += 1i
	} else {
		next += 1.0
	}
	tm.lastTile = next
	return SplitCoord(next)
}

// Add a tile to the map. Tiles are added left to right until the maximum Y value is reached, then moves down to the next row.
func (tm *TileMap) AddTile(t Tile) {
	if tm.tileData == nil {
		tm.init()
	}
	t.SetPosition(tm.getNextTilePosition())
	tm.SetArbitraryTile(t)
}

// Returns a string representation of the map state
func (tm TileMap) String() string {
	rowHeader := "%2d:"
	output := ""
	for rowNum := 0.0; rowNum <= imag(tm.boundary); rowNum++ {
		output += fmt.Sprintf(rowHeader, int(rowNum))
		// causes read of undefined map element if all tiles aren't set
		for colNum := 0.0; colNum <= real(tm.boundary); colNum++ {
			out := tm.tileData[complex(colNum, rowNum)]
			output += fmt.Sprintf("%v", out.String())
		}
		output += "\n"
	}
	return output
}

func (tm TileMap) boundCheck(candidate complex128) bool {
	if real(candidate) < 0 || imag(candidate) < 0 {
		return false
	}
	if real(candidate) > tm.MaxX() || imag(candidate) > tm.MaxY() {
		return false
	}
	return true
}

// Retrieve the tile struct at a given 0 based coordinate
func (tm TileMap) GetTileAt(x, y float64) Tile {
	return tm.tileData[complex(x, y)]
}

// Counts neighbors around a 0 based coordinate that match the provided Tile interface
func (tm TileMap) CountAround(comparator TileFilter, x, y float64) int {
	found := 0
	center := complex(x, y)
	target := center - 1 - 1i
	if tm.boundCheck(target) &&
		comparator(tm.tileData[target]) {
		found++
	}

	target = center - 1i
	if tm.boundCheck(target) &&
		comparator(tm.tileData[target]) {
		found++
	}

	target = center + 1 - 1i
	if tm.boundCheck(target) &&
		comparator(tm.tileData[target]) {
		found++
	}

	target = center - 1
	if tm.boundCheck(target) &&
		comparator(tm.tileData[target]) {
		found++
	}

	target = center + 1
	if tm.boundCheck(target) &&
		comparator(tm.tileData[target]) {
		found++
	}

	target = center - 1 + 1i
	if tm.boundCheck(target) &&
		comparator(tm.tileData[target]) {
		found++
	}

	target = center + 1i
	if tm.boundCheck(target) &&
		comparator(tm.tileData[target]) {
		found++
	}

	target = center + 1 + 1i
	if tm.boundCheck(target) &&
		comparator(tm.tileData[target]) {
		found++
	}
	return found
}
