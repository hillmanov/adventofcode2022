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
	// fmt.Printf("size = %+v\n", size)
	// if size == 1 {
	// 	return []uint{c[endOfWindow]}
	// }
	return c[endOfWindow-size : endOfWindow]
}

func (c Chamber) Dump() {
	for row := range c {
		fmt.Printf("|%07b|\n", c[row])
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

func (s Shape) shiftRight() Shape {
	shifted := utils.CopyOf(s)
	for row := range s {
		if bits.OnesCount(s[row]>>1) != bits.OnesCount(s[row]) {
			return s
		}
	}

	for row := range s {
		shifted[row] >>= 1
	}

	return shifted
}

func (s Shape) shiftLeft() Shape {
	shifted := utils.CopyOf(s)
	for row := range s {
		if bits.OnesCount(s[row]<<1&mask) != bits.OnesCount(s[row]) {
			return s
		}
	}

	for row := range s {
		shifted[row] <<= 1
	}

	return shifted
}

func Part1() any {
	jetPattern := getInput()
	// Steps:
	// Get next shape
	// Position 3 above top of rock/floor of chamber
	// Repeat:
	// Shift left/right
	// Move down
	// Check for collision with rock/floow of chamber

	chamber := Chamber{
		getInt("1111111"),
	}

	rocksFallen := 0
	jetIndex := 0
	for rocksFallen < 3 {
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

			switch jetDirection {
			case ">":
				s = s.shiftRight()
			case "<":
				s = s.shiftLeft()
			}

			fmt.Printf("Shape: = \n%s\n", s)
			if detectCollision(chamber.Window(bottomOfShape+1, len(s)), s) {
				// We've hit something!
				// We can't move down, so we need to add the shape to the chamber
				for row := range s {
					chamber[bottomOfShape-len(s)+row] |= s[row]
				}
				break
			}

			bottomOfShape++
		}
		rocksFallen++
	}

	chamber.Dump()
	return nil
}

func Part2() any {
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
	return strings.Split(contents, "")
}
