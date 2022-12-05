package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
)

type Assignment struct {
	Start int
	End   int
}

func (a Assignment) Contains(b Assignment) bool {
	return a.Start <= b.Start && a.End >= b.End
}

func (a Assignment) Overlaps(b Assignment) bool {
	return a.Start >= b.Start && a.Start <= b.End ||
		a.End >= b.End && a.End <= b.End
}

//go:embed input.txt
var f embed.FS

func Part1() any {
	assignmentPairs := getInput()

	count := 0
	for _, assignmentPair := range assignmentPairs {
		if assignmentPair[0].Contains(assignmentPair[1]) || assignmentPair[1].Contains(assignmentPair[0]) {
			count++
		}
	}

	return count
}

func Part2() any {
	assignmentPairs := getInput()

	count := 0
	for _, assignmentPair := range assignmentPairs {
		if assignmentPair[0].Overlaps(assignmentPair[1]) || assignmentPair[1].Overlaps(assignmentPair[0]) {
			count++
		}
	}

	return count
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 04: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 04: Part 2: = %+v\n", part2Solution)
}

func getInput() [][2]Assignment {
	lines, _ := utils.ReadLines(f, "input.txt")

	assignments := [][2]Assignment{}
	for _, line := range lines {
		first := Assignment{}
		second := Assignment{}

		fmt.Sscanf(line, "%d-%d,%d-%d", &first.Start, &first.End, &second.Start, &second.End)
		assignments = append(assignments, [2]Assignment{first, second})
	}

	return assignments
}
