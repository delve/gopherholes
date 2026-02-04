package year2025

import (
	"aocgen/pkg/common"
	"math"
	"sort"
	"strings"
)

type Day05 struct {
	freshList   [][2]int
	ingredients []int
	min         int
	max         int
}

func (p *Day05) parseInput(lines []string) {
	p.min = math.MaxInt
	p.max = math.MinInt
	for _, line := range lines[:len(lines)-1] {
		if len(line) == 0 {
			continue
		}
		in := strings.Split(line, "-")
		if len(in) == 2 {
			set := [2]int{common.Atoi(in[0]), common.Atoi(in[1])}
			p.freshList = append(p.freshList, set)
			if set[0] < p.min {
				p.min = set[0]
			}
			if set[1] > p.max {
				p.max = set[1]
			}
		} else {
			p.ingredients = append(p.ingredients, common.Atoi(line))
		}
	}
}

func (p Day05) PartA(lines []string) any {
	p.parseInput(lines)
	fresh := 0

ingredientloop:
	for _, ing := range p.ingredients {
		for _, set := range p.freshList {
			if ing >= set[0] && ing <= set[1] {
				fresh++
				continue ingredientloop
			}
		}
	}

	return fresh
}

func (p Day05) PartB(lines []string) any {
	p.parseInput(lines)
	sortedList := make([][2]int, len(p.freshList))
	copy(sortedList, p.freshList)
	sort.Slice(sortedList, func(i, j int) bool {
		return sortedList[i][0] < sortedList[j][0]
	})

	for i := 0; i < len(sortedList)-1; i++ {
		if sortedList[i][1] >= sortedList[i+1][0]-1 {
			// upper overlap
			sortedList[i][1] = max(sortedList[i+1][1], sortedList[i][1])
			sortedList = append(sortedList[:i+1], sortedList[i+2:]...)
			// restest against next higher index
			i--
			continue
		}
	}

	freshcount := 0
	for _, set := range sortedList {
		freshcount += set[1] - set[0] + 1
	}

	return freshcount
}
