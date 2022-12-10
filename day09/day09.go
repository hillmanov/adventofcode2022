package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var f embed.FS

type Move struct {
	Direction string
	Steps     int
}

type Position struct {
	X, Y int
}

func (p *Position) Move(d Delta) {
	p.X -= d.X
	p.Y -= d.Y
}

type Delta struct {
	X, Y int
}

func (p Position) DeltaDirection(o Position) Delta {
	return Delta{
		X: getDirection(p.X - o.X),
		Y: getDirection(p.Y - o.Y),
	}
}

func (p Position) Dist(o Position) float64 {
	return math.Sqrt(math.Pow(float64(p.X-o.X), 2) + math.Pow(float64(p.Y-o.Y), 2))
}

func Part1() any {
	movements := getInput()

	knotCount := 2
	knotPositions := simulateRope(movements, knotCount)

	return len(knotPositions[knotCount-1])
}

func Part2() any {
	movements := getInput()

	knotCount := 10
	knotPositions := simulateRope(movements, knotCount)

	return len(knotPositions[knotCount-1])
}

func simulateRope(moves []Move, knotCount int) []map[Position]bool {
	HEAD := 0

	// Create the knots and the positions
	knots := make([]*Position, knotCount)
	knotPositions := make([]map[Position]bool, knotCount)
	for i := 0; i < knotCount; i++ {
		knots[i] = &Position{0, 0}
		knotPositions[i] = map[Position]bool{*knots[i]: true}
	}

	for _, move := range moves {
		for i := 0; i < move.Steps; i++ {
			switch move.Direction {
			case "U":
				knots[HEAD].Move(Delta{0, 1})
			case "D":
				knots[HEAD].Move(Delta{0, -1})
			case "L":
				knots[HEAD].Move(Delta{1, 0})
			case "R":
				knots[HEAD].Move(Delta{-1, 0})
			}
			knotPositions[HEAD][*knots[HEAD]] = true

			for knot := 0; knot < knotCount-1; knot++ {
				currentKnotPosition := knots[knot]
				nextKnotPosition := knots[knot+1]

				// Now check to see where the tail is relative to the head position and move accordingly, but only if we are far enough away.
				if nextKnotPosition.Dist(*currentKnotPosition) >= 2 {
					nextKnotPosition.Move(nextKnotPosition.DeltaDirection(*currentKnotPosition))
				}
				knotPositions[knot+1][*nextKnotPosition] = true
			}

		}
	}

	return knotPositions
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 09: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 09: Part 2: = %+v\n", part2Solution)
}

func getInput() []Move {
	lines, _ := utils.ReadLines(f, "input.txt")
	movements := []Move{}

	for _, line := range lines {
		movements = append(movements, Move{
			Direction: string(line[0]),
			Steps:     utils.ParseInt(strings.Trim(string(line[1:]), " ")),
		})
	}

	return movements
}

// Util
func getDirection(v int) int {
	if v == 0 {
		return 0
	}
	if v >= 1 {
		return 1
	}
	return -1
}
