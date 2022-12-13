package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
	"math/bits"
)

//go:embed input.txt
var f embed.FS

func Part1() any {
	signal := getInput()
	return findNUnique(signal, 4)
}

func Part2() any {
	signal := getInput()
	return findNUnique(signal, 14)
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 06: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 06: Part 2: = %+v\n", part2Solution)
}

func findNUnique(str string, windowSize int) int {
	set := uint(0)

	for i := 0; i < len(str)-windowSize; i++ {
		set = set ^ (1 << (str[i] - 'a'))
		if i >= windowSize {
			set = set ^ (1 << (str[i-windowSize] - 'a'))
		}
		if bits.OnesCount(set) == windowSize {
			return i + 1
		}
	}
	return -1
}

func getInput() string {
	contents, _ := utils.ReadContents(f, "input.txt")
	return contents
}
