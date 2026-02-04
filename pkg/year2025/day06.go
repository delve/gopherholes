package year2025

import (
	"aocgen/pkg/common"
	"fmt"
	"regexp"
	"strings"
)

type Day06 struct {
	problems     [][]string
	cephProblems []cephalopodProblem
}

type cephalopodProblem struct {
	numbers []int
	operand rune
}

func (p *Day06) parseInput(lines []string) {
	p.problems = [][]string{}
	tokenExtractor := func(c rune) bool {
		return c == ' '
	}
	for _, line := range lines[:len(lines)-1] {
		p.problems = append(p.problems, strings.FieldsFunc(line, tokenExtractor))
	}
	problemCount := len(p.problems[0])
	for idx, line := range p.problems {
		if len(line) != problemCount {
			panic(fmt.Sprintf("Inconsistent problem count line %d. Expected %d, found %d", idx, problemCount, len(line)))
		}
	}
}

func (p Day06) PartA(lines []string) any {
	p.parseInput(lines)

	mathSum := 0

	for i := 0; i < len(p.problems[0]); i++ {
		operand := p.problems[len(p.problems)-1][i]
		sum := 0
		if operand == "*" {
			sum = 1
		}
		for j := 0; j < len(p.problems)-1; j++ {
			switch operand {
			case "*":
				sum = sum * common.Atoi(p.problems[j][i])
			case "+":
				sum = sum + common.Atoi(p.problems[j][i])
			}
		}
		mathSum += sum
	}

	if len(lines[0]) <= 17 && mathSum == 4277556 {
		println("Correct")
	}
	return mathSum
}

func (p *Day06) parseCephalopodInput(lines []string) {
	lines = lines[:len(lines)-1]
	p.cephProblems = []cephalopodProblem{}
	digits := regexp.MustCompile(`[0-9]`)
	problem := cephalopodProblem{}
	for i := len(lines[0]) - 1; i >= 0; i-- {
		number := ""
		for j := 0; j < len(lines); j++ {
			val := string(lines[j][i])
			switch {
			case val == " ":
				continue
			case digits.MatchString(val):
				number += val
			case val == "+":
				problem.operand = '+'
			case val == "*":
				problem.operand = '*'
			}
		}
		problem.numbers = append(problem.numbers, common.Atoi(number))
		if problem.operand != 0 {
			// problem is complete, store it and advance the cursor
			p.cephProblems = append(p.cephProblems, problem)
			problem = cephalopodProblem{}
			i-- // skip the blank column between problems
		}
	}
}

func (p Day06) PartB(lines []string) any {
	p.parseCephalopodInput(lines)
	mathSum := 0
	for _, problem := range p.cephProblems {
		switch problem.operand {
		case '+':
			sum := 0
			for _, num := range problem.numbers {
				sum += num
			}
			mathSum += sum
		case '*':
			product := 1
			for _, num := range problem.numbers {
				product *= num
			}
			mathSum += product
		}
	}
	if len(lines[0]) <= 17 && mathSum == 3263827 {
		println("Correct")
	}
	return mathSum
}
