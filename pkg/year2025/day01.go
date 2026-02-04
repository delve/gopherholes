package year2025

import (
	"aocgen/pkg/common"
	"aocgen/pkg/common/dials"
	"math"
)

type Day01 struct {
	sequence []int
}

func (p *Day01) parseInput(lines []string) {
	for _, line := range lines[:len(lines)-1] {
		dir := line[0]
		step := common.Atoi(line[1:])
		if dir == 'L' {
			step = step * -1
		}
		p.sequence = append(p.sequence, step)
	}
}

func (p Day01) PartA(lines []string) any {
	p.parseInput(lines)

	zeroCount := 0
	d := dials.New()
	d.Set(50)
	for _, step := range p.sequence {
		d.Step(step)
		if d.Position() == 0 {
			zeroCount++
		}
	}

	return zeroCount
}

func (p Day01) PartB(lines []string) any {
	p.parseInput(lines)
	zeroCount := 0
	d := dials.New()
	d.Set(50)

	for _, step := range p.sequence {
		orgPos := d.Position()
		// add the looparounds
		zeroCount += int(math.Abs(float64(step / 100)))
		step = step % 100
		d.Step(step)
		if d.Position() == 0 { // landed on 0
			zeroCount++
		} else if orgPos != 0 { // didn't land on 0 but might have passed it without looping
			if step < 0 {
				if (orgPos + step) < 0 {
					zeroCount++
				}
			} else {
				if orgPos+step > 100 {
					zeroCount++
				}
			}
		}
	}

	// for _, step := range p.sequence {
	// 	orgPos := d.Position()
	// 	// add the looparounds
	// 	if step > 0 {
	// 		zeroCount += (step + orgPos) / 100
	// 	} else {
	// 		if step+orgPos < 0 {
	// 			// zeroCount++
	// 			zeroCount += (step + orgPos) / -100
	// 		}
	// 	}
	// 	d.Step(step)
	// 	if d.Position() == 0 { // landed on 0
	// 		zeroCount++
	// 	}
	// }

	return zeroCount
}
