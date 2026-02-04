package year2025

import (
	"aocgen/pkg/common"
	"math"
)

type Day03 struct {
	batteryBanks [][]int
}

func (p *Day03) parseInput(lines []string) {
	for _, line := range lines[:len(lines)-1] {
		bank := []int{}
		for i := 0; i < len(line); i++ {
			bank = append(bank, common.Atoi(line[i:i+1]))
		}
		p.batteryBanks = append(p.batteryBanks, bank)
	}
}

func (p Day03) getJolts(bankIndex, digits int) int {
	bank := p.batteryBanks[bankIndex]
	// FS Go, why can't i use a slice here
	// tens := max(bank[0:len(bank)-1]...)
	totalJolts := 0
	cursor := 0
	for i := digits; i >= 0; i-- {
		jolts := 0
		jPos := 0
		for pos, val := range bank[cursor : len(bank)-i+1] {
			if val > jolts {
				jolts = val
				jPos = pos + 1
			}
		}
		totalJolts += jolts * int(math.Pow10(i-1))
		cursor += jPos
	}

	return totalJolts
}

func (p Day03) PartA(lines []string) any {
	p.parseInput(lines)
	maxJolt := 0
	joltageDigits := 2

	for i := 0; i < len(p.batteryBanks); i++ {
		maxJolt += p.getJolts(i, joltageDigits)
	}

	if maxJolt == 357 || maxJolt == 17196 {
		println("Correct")
	}
	return maxJolt
}

func (p Day03) PartB(lines []string) any {
	p.parseInput(lines)
	maxJolt := 0
	joltageDigits := 12

	for i := 0; i < len(p.batteryBanks); i++ {
		bankJolts := p.getJolts(i, joltageDigits)
		if bankJolts/100000000000 > 10 {
			println("Bignum")
			println(i)
		}
		maxJolt += bankJolts
	}

	if maxJolt == 3121910778619 || maxJolt == 171039099596062 {
		println("Correct")
	}
	return maxJolt
}
