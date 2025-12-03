package year2025

import (
	"aocgen/pkg/common"
	"fmt"
	"strings"
)

type Day02 struct {
	ranges [][]string
}

func (p *Day02) parseInput(lines []string) {
	ranges := [][]string{}
	for _, idRange := range strings.Split(lines[0], ",") {
		ranges = append(ranges, strings.Split(idRange, "-"))
	}
	p.ranges = ranges
}

func (p Day02) PartA(lines []string) any {
	p.parseInput(lines)

	sumInvalids := 0

	for _, idRange := range p.ranges {
		idStart := common.Atoi(idRange[0])
		idStop := common.Atoi(idRange[1])

		for i := idStart; i <= idStop; i++ {
			strId := fmt.Sprintf("%d", i)
			first := strId[0 : len(strId)/2]
			last := strId[len(strId)/2:]
			if first == last {
				sumInvalids += i
			}
		}
	}

	return sumInvalids
}

func (p Day02) PartB(lines []string) any {
	p.parseInput(lines)

	sumInvalids := 0

	for _, idRange := range p.ranges {
		idStart := common.Atoi(idRange[0])
		idStop := common.Atoi(idRange[1])
	idLoop:
		for i := idStart; i <= idStop; i++ {
			strId := fmt.Sprintf("%d", i)
		repeatLoop:
			for rptLen := 1; rptLen <= len(strId)/2; rptLen++ {
				if len(strId)%rptLen != 0 {
					//early break, it's not even division so it cant be a repeat of this length
					continue
				}
				for x := rptLen; x < len(strId); x = x + rptLen {
					if strId[0:rptLen] != strId[x:x+rptLen] {
						// doesn't match this segment, not a repeat, breakout to the next length
						continue repeatLoop
					}
				}
				// all segments match, it's invalid, add it and breakout to the next id
				sumInvalids += i
				continue idLoop
			}
		}
	}

	return sumInvalids
}
