package main

import (
	"embed"
	"fmt"
)

//go:embed input.txt
var f embed.FS

func Part1() any {
	return nil
}

func Part2() any {
	return nil
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

  fmt.Printf("Day 23: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 23: Part 2: = %+v\n", part2Solution)
}
