package tilemap

type RuneTile struct {
	// zero based XY coords
	position complex128
	// Store random data about a tile. Must be convertible to rune for printing
	Value rune
}

func (t *RuneTile) SetPosition(x, y float64) {
	t.position = complex(x, y)
}

func (t RuneTile) GetPosition() complex128 {
	return t.position
}

func (t RuneTile) X() float64 {
	return real(t.position)
}

func (t RuneTile) Y() float64 {
	return imag(t.position)
}

func (t RuneTile) GetValue() any {
	return t.Value
}

func (t *RuneTile) SetValue(v any) {
	t.Value = v.(rune)
}

func (t RuneTile) String() string {
	return string(t.Value)
}

func (t RuneTile) Rune() rune {
	return t.Value
}
