package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var f embed.FS

const priority = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Rucksack struct {
	Contents []string
}

func (r Rucksack) Compartment1() []string {
	return r.Contents[0 : len(r.Contents)/2]
}

func (r Rucksack) Compartment2() []string {
	return r.Contents[len(r.Contents)/2:]
}

func Part1() any {
	rucksacks := getInput()

	sum := 0
	for _, r := range rucksacks {
		for _, c := range utils.Intersection(r.Compartment1(), r.Compartment2()) {
			sum += strings.Index(priority, c)
		}
	}
	return sum
}

func Part2() any {
	rucksacks := getInput()

	sum := 0
	for i := 0; i < len(rucksacks); i += 3 {
		group := rucksacks[i : i+3]
		badge := utils.Intersection(group[0].Contents, utils.Intersection(group[1].Contents, group[2].Contents))
		sum += strings.Index(priority, badge[0])
	}
	return sum
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 03: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 03: Part 2: = %+v\n", part2Solution)
}

func getInput() []Rucksack {
	lines, _ := utils.ReadLines(f, "input.txt")
	rucksacks := []Rucksack{}

	for _, line := range lines {
		rucksacks = append(rucksacks, Rucksack{
			Contents: strings.Split(line, ""),
		})
	}

	return rucksacks
}
