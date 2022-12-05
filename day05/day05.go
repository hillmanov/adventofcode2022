package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed input.txt
var f embed.FS

type Crate = string

type Move struct {
	Amount int
	From   int
	To     int
}

func Part1() any {
	stacks := readStacks()
	moves := readMoves()

	for _, move := range moves {
		stacks[move.To] = append(stacks[move.To], utils.Reverse(stacks[move.From][len(stacks[move.From])-move.Amount:])...)
		stacks[move.From] = stacks[move.From][:len(stacks[move.From])-move.Amount]
	}

	topCrates := ""
	for i := 1; i <= len(stacks); i++ {
		topCrates += stacks[i][len(stacks[i])-1]
	}

	return topCrates
}

func Part2() any {
	stacks := readStacks()
	moves := readMoves()

	for _, move := range moves {
		stacks[move.To] = append(stacks[move.To], stacks[move.From][len(stacks[move.From])-move.Amount:]...)
		stacks[move.From] = stacks[move.From][:len(stacks[move.From])-move.Amount]
	}

	topCrates := ""
	for i := 1; i <= len(stacks); i++ {
		topCrates += stacks[i][len(stacks[i])-1]
	}

	return topCrates
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 05: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 05: Part 2: = %+v\n", part2Solution)
}

func readStacks() map[int][]Crate {
	isNumber := regexp.MustCompile(`^\d+$`)
	isLetter := regexp.MustCompile(`^[A-Z]$`)

	lines, _ := utils.ReadLines(f, "input.txt")

	raw := [][]string{}
	for _, line := range lines {
		if line == "" {
			break
		}
		raw = append(raw, strings.Split(line, ""))
	}

	var currentStack int
	stacks := map[int][]Crate{}
	for i := 0; i < len(raw[0]); i++ {
		for j := len(raw) - 1; j >= 0; j-- {
			switch {
			case isNumber.MatchString(raw[j][i]):
				currentStack = utils.ParseInt(raw[j][i])
				stacks[currentStack] = []Crate{}
			case isLetter.MatchString(raw[j][i]):
				stacks[currentStack] = append(stacks[currentStack], raw[j][i])
			}

		}
	}
	return stacks
}

func readMoves() []Move {
	lines, _ := utils.ReadLines(f, "input.txt")

	moves := []Move{}
	for _, line := range lines {
		if !strings.HasPrefix(line, "move") {
			continue
		}
		move := Move{}
		fmt.Sscanf(line, "move %d from %d to %d", &move.Amount, &move.From, &move.To)
		moves = append(moves, move)
	}

	return moves
}
