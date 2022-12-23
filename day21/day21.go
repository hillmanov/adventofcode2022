package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
	"math"
	"regexp"
	"strings"
)

type MonkeyType string

const (
	NUMBER = MonkeyType("NUMBER")
	MATH   = MonkeyType("MATH")
)

type Monkey struct {
	Name string
	Type MonkeyType

	// For NUMBER monkeys
	Number float64

	// For MATH monkeys
	Left      string
	Right     string
	Operation string
}

func (m Monkey) GetValue(monkeysByName map[string]*Monkey) float64 {
	if m.Type == NUMBER {
		return m.Number
	}

	left := monkeysByName[m.Left]
	right := monkeysByName[m.Right]

	switch m.Operation {
	case "+":
		return left.GetValue(monkeysByName) + right.GetValue(monkeysByName)
	case "-":
		return left.GetValue(monkeysByName) - right.GetValue(monkeysByName)
	case "*":
		return left.GetValue(monkeysByName) * right.GetValue(monkeysByName)
	case "/":
		return left.GetValue(monkeysByName) / right.GetValue(monkeysByName)
	case "=":
		return left.GetValue(monkeysByName) - right.GetValue(monkeysByName) // If this returns 0, then they are equal
	}

	return 0
}

//go:embed input.txt
var f embed.FS

func Part1() any {
	monkeysByName := getInput()
	return int(monkeysByName["root"].GetValue(monkeysByName))
}

func Part2() any {
	monkeysByName := getInput()

	root := monkeysByName["root"]
	human := monkeysByName["humn"]
	root.Operation = "="

	minMax := []float64{float64(math.MinInt), float64(math.MaxInt)}
	human.Number = (minMax[0] + minMax[1]) / 2
	for root.GetValue(monkeysByName) != 0 {
		if root.GetValue(monkeysByName) > 0 {
			minMax[0] = human.Number
		} else {
			minMax[1] = human.Number
		}
		human.Number = (minMax[0] + minMax[1]) / 2
	}
	return int(human.Number)
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 21: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 21: Part 2: = %+v\n", part2Solution)
}

func getInput() map[string]*Monkey {
	lines, _ := utils.ReadLines(f, "input.txt")
	isNumber := regexp.MustCompile(`^\d+$`)

	monkeysByName := map[string]*Monkey{}

	for _, line := range lines {
		monkey := Monkey{}

		parts := strings.Split(line, ":")

		monkey.Name = strings.Trim(parts[0], " ")

		if isNumber.MatchString(strings.Trim(parts[1], " ")) {
			monkey.Type = NUMBER
			monkey.Number = float64(utils.ParseInt(parts[1]))
		} else {
			monkey.Type = MATH
			mathParts := strings.Split(strings.Trim(parts[1], " "), " ")
			monkey.Left = strings.Trim(mathParts[0], " ")
			monkey.Operation = strings.Trim(mathParts[1], " ")
			monkey.Right = strings.Trim(mathParts[2], " ")
		}
		monkeysByName[monkey.Name] = &monkey
	}

	return monkeysByName
}
