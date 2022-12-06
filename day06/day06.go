package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
	"strings"
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

func findNUnique(str string, n int) int {
	for i := 0; i < len(str)-n; i++ {
		if len(utils.UniqueOf(strings.Split(str[i:i+n], ""))) == n {
			return i + n
		}
	}
	return -1
}

func getInput() string {
	contents, _ := utils.ReadContents(f, "input.txt")
	return contents
}
