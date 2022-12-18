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

var mask = getInt("1111111")
var clearMask = getInt("0000000")

type Shape []uint

type Chamber []uint

func (c Chamber) floorOrHighestRock() int {
	var floorOrHighestRock int
	for row := range c {
		if bits.OnesCount(c[row]) > 0 {
			floorOrHighestRock = row
			break
		}
	}
	return floorOrHighestRock
}

func (c Chamber) Window(endOfWindow, size int) []uint {
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

var shapes = []Shape{
	{
		getInt("0011110"), // Horizontal line
	},
	{
		getInt("0001000"), // Plus
		getInt("0011100"),
		getInt("0001000"),
	},
	{
		getInt("0000100"), // L (backwards)
		getInt("0000100"),
		getInt("0011100"),
	},
	{
		getInt("0010000"), // Vertical line
		getInt("0010000"),
		getInt("0010000"),
		getInt("0010000"),
	},
	{
		getInt("0011000"), // Square
		getInt("0011000"),
	},
}

func (s Shape) String() string {
	var str string
	for row := range s {
		str += fmt.Sprintf("%07b\n", s[row])
	}
	return str
}

func (s Shape) shift(direction string) Shape {
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

	chamber := Chamber{ // Initialize with floor
		getInt("1111111"),
	}

	jetIndex := 0
	for rocksFallen := 0; rocksFallen < 2022; rocksFallen++ {
		s := shapes[rocksFallen%len(shapes)]

		// Pad the top if we need room (This can potentially add more room to the top than needed, but I am not worrying about that now)
		// Find the row of the top rock/floor of the chamber
		floorOrHighestRock := chamber.floorOrHighestRock()

		for i := 0; i < (len(s)+2)-floorOrHighestRock; i++ {
			chamber = utils.Shift(chamber, getInt("0000000"))
		}

		floorOrHighestRock = chamber.floorOrHighestRock()

		// We start our shape 3 rows above the highest rock/floor of chamber
		bottomOfShape := floorOrHighestRock - 3

		for {
			jetDirection := jetPattern[jetIndex%len(jetPattern)]
			jetIndex++

			// Try to move left or right
			if !detectCollision(chamber.Window(bottomOfShape, len(s)), s.shift(jetDirection)) {
				s = s.shift(jetDirection)
			}

			// Try to move down
			if detectCollision(chamber.Window(bottomOfShape+1, len(s)), s) {
				for row := range s {
					chamber[bottomOfShape-len(s)+row] |= s[row]
				}
				break
			}

			bottomOfShape++
		}
	}

	return nil
}

func Part2() any {
	jetPattern := getInput()
	fmt.Printf("len(jetPattern) = %+v\n", len(jetPattern))

	// Run len(shapes) * len(jetPattern) times. Get the height at that point.
	// Multiply the height by how many times it can go into 1_000_000_000_000
	// Get the top few row of the chamber at that point.
	// Initialize the chamber with that top window,
	// Run the stuff again however many times are remaining from the initial division. See how high it go.
	// Add some numbers, and you are done!

	return nil
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 17: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 17: Part 2: = %+v\n", part2Solution)
}

func detectCollision(chamberWindow []uint, s Shape) bool {
	for row := range chamberWindow {
		if bits.OnesCount(chamberWindow[row]^s[row]) < bits.OnesCount(chamberWindow[row])+bits.OnesCount(s[row]) {
			return true
		}
	}
	return false
}

func getInt(binary string) uint {
	i, _ := strconv.ParseInt(binary, 2, 64)
	return uint(i)
}

func getInput() []string {
	contents, _ := utils.ReadContents(f, "input.txt")
	return strings.Split(strings.Trim(contents, "\n"), "")
}
