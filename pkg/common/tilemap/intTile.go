package tilemap

import "fmt"

type IntTile struct {
	// zero based XY coords
	position complex128
	// Store random data about a tile. Must be convertible to rune for printing
	Value int
}

func (t *IntTile) SetPosition(x, y float64) {
	t.position = complex(x, y)
}

func (t IntTile) GetPosition() complex128 {
	return t.position
}

func (t IntTile) X() float64 {
	return real(t.position)
}

func (t IntTile) Y() float64 {
	return imag(t.position)
}

func (t IntTile) GetValue() any {
	return t.Value
}

func (t *IntTile) SetValue(v any) {
	t.Value = v.(int)
}

func (t IntTile) String() string {
	return fmt.Sprintf("%d", t.Value)
}

func (t IntTile) Rune() rune {
	return rune(t.Value)
}
