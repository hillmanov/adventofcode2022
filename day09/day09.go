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

type Motion struct {
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

	ropeLength := 2
	ropePositions := simulateRope(movements, ropeLength)

	return len(ropePositions[ropeLength-1])
}

func Part2() any {
	movements := getInput()

	ropeLength := 10
	ropePositions := simulateRope(movements, ropeLength)

	return len(ropePositions[ropeLength-1])
}

func simulateRope(movements []Motion, ropeLength int) []map[Position]bool {
	HEAD := 0

	// Create the rope and the positions
	rope := make([]*Position, ropeLength)
	ropePositions := make([]map[Position]bool, ropeLength)
	for i := 0; i < ropeLength; i++ {
		rope[i] = &Position{0, 0}
		ropePositions[i] = map[Position]bool{*rope[i]: true}
	}

	for _, movement := range movements {
		for i := 0; i < movement.Steps; i++ {
			for segment := 0; segment < ropeLength-1; segment++ {

				// Head and tail are relative to each other in the rope segment
				headPosition := rope[segment]
				tailPosition := rope[segment+1]

				if segment == HEAD { // Movements apply to the head only
					switch movement.Direction {
					case "U":
						headPosition.Move(Delta{0, 1})
					case "D":
						headPosition.Move(Delta{0, -1})
					case "L":
						headPosition.Move(Delta{1, 0})
					case "R":
						headPosition.Move(Delta{-1, 0})
					}
					ropePositions[segment][*headPosition] = true
				}

				// Now check to see where the tail is relative to the head position and move accordingly, but only if we are far enough away.
				if tailPosition.Dist(*headPosition) >= 2 {
					tailPosition.Move(tailPosition.DeltaDirection(*headPosition))
				}
				ropePositions[segment+1][*tailPosition] = true
			}

		}
	}

	return ropePositions
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 09: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 09: Part 2: = %+v\n", part2Solution)
}

func getInput() []Motion {
	lines, _ := utils.ReadLines(f, "input.txt")
	movements := []Motion{}

	for _, line := range lines {
		movements = append(movements, Motion{
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
