package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
	"sort"
)

//go:embed input.txt
var f embed.FS

type elf struct {
	food []int
}

func (e elf) totalCalories() int {
	return utils.SumOf(e.food)
}

func Part1() any {
	elves := readInput()
	maxCalories := 0
	for _, e := range elves {
		maxCalories = utils.Max(maxCalories, e.totalCalories())
	}
	return maxCalories
}

func Part2() any {
	elves := readInput()

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].totalCalories() > elves[j].totalCalories()
	})

	return utils.SumOf([]int{elves[0].totalCalories(), elves[1].totalCalories(), elves[2].totalCalories()})
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 01: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 01: Part 2: = %+v\n", part2Solution)
}

func readInput() []elf {
	lines, _ := utils.ReadLines(f, "input.txt")

	elves := []elf{}
	e := elf{}
	for _, line := range lines {
		if line == "" {
			elves = append(elves, e)
			e = elf{}
			continue
		}
		e.food = append(e.food, utils.ParseInt(line))
	}
	elves = append(elves, e)
	return elves
}
