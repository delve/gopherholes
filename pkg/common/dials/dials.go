package dials

// Dial represents a circular dial with 100 positions (0–99).
type Dial struct {
	pos int
}

// New returns a new Dial starting at position 0.
func New() *Dial {
	return &Dial{pos: 0}
}

// Position returns the current position of the dial.
func (d *Dial) Position() int {
	return d.pos
}

// Set sets dial to a specific position
func (d *Dial) Set(p int) {
	d.pos = ((p % 100) + 100) % 100 // ensure range 0–99, handle negatives
}

// Step moves dial by n steps
func (d *Dial) Step(n int) {
	d.pos = ((d.pos+n)%100 + 100) % 100
}

// Right is a semantic wrapper for Step(+x).
func (d *Dial) Right(step int) {
	if step < 0 {
		step = step * -1
	}
	d.Step(step)
}

// Left is a semantic wrapper for Step(+x).
func (d *Dial) Left(step int) {
	if step > 0 {
		step = step * -1
	}
	d.Step(step)
}
