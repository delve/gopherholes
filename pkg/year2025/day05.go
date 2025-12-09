package year2025

import (
	"aocgen/pkg/common"
	"fmt"
	"math"
	"sort"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
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

func (p *Day05) parseFresh(lines []string) int {
	fresh := mapset.NewThreadUnsafeSet[int]()

	for _, line := range lines[:len(lines)-1] {
		if len(line) == 0 {
			break
		}
		in := strings.Split(line, "-")
		freshSet := [2]int{common.Atoi(in[0]), common.Atoi(in[1])}
		for i := freshSet[0]; i <= freshSet[1]; i++ {
			fresh.Add(i)
		}

	}

	return fresh.Cardinality()
}

func (p Day05) PartB(lines []string) any {
	p.parseInput(lines)
	sortedList := make([][2]int, len(p.freshList))
	copy(sortedList, p.freshList)
	sort.Slice(sortedList, func(i, j int) bool {
		return sortedList[i][0] < sortedList[j][0]
	})

	for i := 0; i < len(sortedList)-1; i++ {
		lowerBoundCheck := 0
		if i > 0 {
			lowerBoundCheck = sortedList[i][0] - sortedList[i-1][1]
		}
		upperBoundCheck := sortedList[i+1][0] - sortedList[i][1]
		fmt.Printf("len: %d\ni: %d\nlower: %d\nupper: %d\n----------\n", len(sortedList), i, lowerBoundCheck, upperBoundCheck)

		if i > 0 && sortedList[i][0] <= sortedList[i-1][1]+1 {
			// lower overlap
			sortedList[i][0] = sortedList[i-1][0]
			sortedList = append(sortedList[:i-1], sortedList[i:]...)
			// restest against next lower index
			i--
			continue
		}
		if sortedList[i][1] >= sortedList[i+1][0]-1 {
			// upper overlap
			sortedList[i][1] = sortedList[i+1][1]
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

	if len(lines) < 13 {
		if freshcount == 14 {
			println("correct")
		} else {
			println("incorrect")
		}

	} else {
		if freshcount <= 336231364595248 {
			println("too low")
		} else if freshcount >= 355473061408314 {
			println("too high")
		} else {
			println("might be correct?")
		}
	}

	return freshcount
}
