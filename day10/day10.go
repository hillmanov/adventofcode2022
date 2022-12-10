package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var f embed.FS

type Instruction struct {
	OpCode string
	Arg    int
}

type registerValue = int

func Part1() any {
	instructions := getInput()
	cycleValues := runProgram(instructions)

	signalStrength := 0
	for _, index := range []int{20, 60, 100, 140, 180, 220} {
		signalStrength += ((index) * cycleValues[index-1])
	}

	return signalStrength
}

func Part2() any {
	instructions := getInput()
	cycleValues := runProgram(instructions)

	s := strings.Builder{}
	cycleIndex := 0
	for row := 0; row < 6; row++ {
		for col := 0; col < 40; col++ {
			if inRange(cycleValues[cycleIndex], col) {
				s.WriteString("#")
			} else {
				s.WriteString(".")
			}
			cycleIndex++
		}
		s.WriteString("\n")
	}

	return s.String()
}

func runProgram(instructions []Instruction) []registerValue {
	cycleValues := []registerValue{}
	x := 1
	for _, instruction := range instructions {
		switch instruction.OpCode {
		case "noop":
			cycleValues = append(cycleValues, x)
		case "addx":
			cycleValues = append(cycleValues, x)
			cycleValues = append(cycleValues, x)
			x += instruction.Arg
		}
	}

	return cycleValues
}

func inRange(middle int, check int) bool {
	return middle-1 == check || middle == check || middle+1 == check
}

func getInput() []Instruction {
	lines, _ := utils.ReadLines(f, "input.txt")

	instructions := []Instruction{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		instruction := Instruction{
			OpCode: parts[0],
		}

		if instruction.OpCode == "addx" {
			instruction.Arg = utils.ParseInt(parts[1])
		}
		instructions = append(instructions, instruction)
	}

	return instructions
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 10: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 10: Part 2: = \n%+v\n", part2Solution)
}
