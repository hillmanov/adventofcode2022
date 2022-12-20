package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
	"math/bits"
	"strconv"
	"strings"
)

//go:embed input.txt
var f embed.FS

var mask = b("1111111")
var clearMask = b("0000000")

type Rock []uint

type Chamber []uint

func (c Chamber) Window(endOfWindow, size int) Chamber {
	return c[utils.Max(endOfWindow-size, 0):endOfWindow]
}

func (c Chamber) Dump() {
	for row := range c {
		row := fmt.Sprintf("|%07b|", c[row])
		row = strings.ReplaceAll(row, "0", ".")
		row = strings.ReplaceAll(row, "1", "#")
		fmt.Printf("%s\n", row)
	}
}

func (c Chamber) floorOrPotentialNewFloor() int {
	for row := range c {
		if c[row] == b("1111111") {
			return row
		}
	}
	return -1
}

func (c Chamber) floorOrHighestRock() int {
	for row := range c {
		if c[row] != 0 {
			return row
		}
	}
	return -1
}

var rockShapes = []Rock{
	{
		b("0011110"), // Horizontal line
	},
	{
		b("0001000"), // Plus
		b("0011100"),
		b("0001000"),
	},
	{
		b("0000100"), // L (backwards)
		b("0000100"),
		b("0011100"),
	},
	{
		b("0010000"), // Vertical line
		b("0010000"),
		b("0010000"),
		b("0010000"),
	},
	{
		b("0011000"), // Square
		b("0011000"),
	},
}

func (s Rock) String() string {
	var str string
	for row := range s {
		str += fmt.Sprintf("%07b\n", s[row])
	}
	return str
}

func (s Rock) shift(direction string) Rock {
	shifted := utils.CopyOf(s)

	switch direction {
	case ">":
		for row := range s {
			if bits.OnesCount(s[row]>>1) != bits.OnesCount(s[row]) {
				return s
			}
		}

		for row := range s {
			shifted[row] >>= 1
		}
	case "<":
		for row := range s {
			if bits.OnesCount(s[row]<<1&mask) != bits.OnesCount(s[row]) {
				return s
			}
		}

		for row := range s {
			shifted[row] <<= 1
		}
	}

	return shifted
}

func Part1() any {
	jetPattern := getInput()

	chamber := Chamber(make([]uint, 2022*3)) // Number of rocks * number of shapes. If we pad the slice a lot, we save a LOT of time.
	chamber[len(chamber)-1] = b("1111111")   // Create a floor

	chamber = run(chamber, jetPattern, 2022)
	return len(chamber) - chamber.floorOrHighestRock() - 1
}

func Part2() any {
	return nil
}

func run(chamber Chamber, jetPattern []string, iterations int) Chamber {
	floorOrHighestRock := chamber.floorOrHighestRock()

	jetIndex := 0
	for rocksFallen := 0; rocksFallen < iterations; rocksFallen++ {
		rock := rockShapes[rocksFallen%len(rockShapes)]

		// We start our shape 3 rows above the highest rock/floor of chamber
		bottomOfRock := floorOrHighestRock - 3

		for {
			jetDirection := jetPattern[jetIndex%len(jetPattern)]
			jetIndex++

			// Try to move left or right
			if !detectCollision(chamber.Window(bottomOfRock, len(rock)), rock.shift(jetDirection)) {
				rock = rock.shift(jetDirection)
			}

			// Try to move down
			if detectCollision(chamber.Window(bottomOfRock+1, len(rock)), rock) {
				for row := range rock {
					chamber[bottomOfRock-len(rock)+row] |= rock[row]
					floorOrHighestRock = utils.Min(floorOrHighestRock, bottomOfRock-len(rock))
				}
				break
			}

			bottomOfRock++
		}
	}

	return chamber
}

func detectCollision(chamberWindow Chamber, s Rock) bool {
	for row := range chamberWindow {
		if bits.OnesCount(chamberWindow[row]^s[row]) < bits.OnesCount(chamberWindow[row])+bits.OnesCount(s[row]) {
			return true
		}
	}
	return false
}

func b(binary string) uint {
	i, _ := strconv.ParseInt(binary, 2, 64)
	return uint(i)
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 17: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 17: Part 2: = %+v\n", part2Solution)
}

func getInput() []string {
	contents, _ := utils.ReadContents(f, "input.txt")
	return strings.Split(strings.Trim(contents, "\n"), "")
}

// Check to see if there is a repeating pattern in the array
func hasRepeatingPattern(chamber Chamber) bool {
	chamber = chamber[chamber.floorOrHighestRock() : len(chamber)-1] // Start at the highest rock and go up before the floor

	if len(chamber)%2 != 0 {
		return false
	}

	for i := 0; i < len(chamber)/2; i++ {
		if chamber[i] != chamber[i+len(chamber)/2] {
			return false
		}
	}

	return true

}
