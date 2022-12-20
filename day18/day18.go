package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
	"math"
	"math/bits"
)

//go:embed input.txt
var f embed.FS

const (
	NONE   = 0
	TOP    = 1
	BOTTOM = 1 << 2
	LEFT   = 1 << 3
	RIGHT  = 1 << 4
	FRONT  = 1 << 5
	BACK   = 1 << 6
)

type Location struct {
	X, Y, Z int
}

type Cube struct {
	X, Y, Z        int
	Exposed        uint
	ExposedToWater uint
}

func (c *Cube) Init() {
	c.Exposed = TOP | BOTTOM | LEFT | RIGHT | FRONT | BACK
	c.ExposedToWater = 0
}

func (c *Cube) Cover(side uint) {
	c.Exposed &= ^side
}

func (c *Cube) CoverWithWater(side uint) {
	c.ExposedToWater |= side
}

func (c Cube) ExposedSides() int {
	return bits.OnesCount(c.Exposed)
}

func (c Cube) ExposedSidesToWater() int {
	return bits.OnesCount(c.ExposedToWater)
}

const (
	OBSIDIAN = 1
	WATER    = 2
)

type Container struct {
	MinX   int
	MaxX   int
	MinY   int
	MaxY   int
	MinZ   int
	MaxZ   int
	Filled map[Location]int
}

type Obsidian map[Location]*Cube

func GetSharedSides(c1, c2 *Cube) (uint, uint) {
	switch {
	case (c1.X-c2.X) == 1 && c1.Y == c2.Y && c1.Z == c2.Z: // c1 is to the right of c2
		return LEFT, RIGHT
	case (c1.X-c2.X) == -1 && c1.Y == c2.Y && c1.Z == c2.Z: // c1 is to the right of c2
		return RIGHT, LEFT
	case (c1.Y-c2.Y) == 1 && c1.X == c2.X && c1.Z == c2.Z: // c1 is on top of c2
		return BOTTOM, TOP
	case (c1.Y-c2.Y) == -1 && c1.X == c2.X && c1.Z == c2.Z: // c1 is on bottom of c2
		return TOP, BOTTOM
	case (c1.Z-c2.Z) == 1 && c1.X == c2.X && c1.Y == c2.Y: // c1 is in front of c2
		return BACK, FRONT
	case (c1.Z-c2.Z) == -1 && c1.X == c2.X && c1.Y == c2.Y: // c1 is behind c2
		return FRONT, BACK
	}
	return NONE, NONE
}

func Part1() any {
	obsidian := getInput()

	for location, cube := range obsidian {
		if _, ok := obsidian[Location{X: location.X + 1, Y: location.Y, Z: location.Z}]; ok {
			cube.Cover(RIGHT)
		}
		if _, ok := obsidian[Location{X: location.X - 1, Y: location.Y, Z: location.Z}]; ok {
			cube.Cover(LEFT)
		}
		if _, ok := obsidian[Location{X: location.X, Y: location.Y + 1, Z: location.Z}]; ok {
			cube.Cover(TOP)
		}
		if _, ok := obsidian[Location{X: location.X, Y: location.Y - 1, Z: location.Z}]; ok {
			cube.Cover(BOTTOM)
		}
		if _, ok := obsidian[Location{X: location.X, Y: location.Y, Z: location.Z + 1}]; ok {
			cube.Cover(FRONT)
		}
		if _, ok := obsidian[Location{X: location.X, Y: location.Y, Z: location.Z - 1}]; ok {
			cube.Cover(BACK)
		}
	}

	surfaceArea := 0
	for _, cube := range obsidian {
		surfaceArea += cube.ExposedSides()
	}

	return surfaceArea
}

func Part2() any {
	obsidian := getInput()

	minX, maxX, minY, maxY, minZ, maxZ := math.MaxInt, math.MinInt, math.MaxInt, math.MinInt, math.MaxInt, math.MinInt

	for location := range obsidian {
		minX = utils.Min(minX, location.X)
		maxX = utils.Max(maxX, location.X)
		minY = utils.Min(minY, location.Y)
		maxY = utils.Max(maxY, location.Y)
		minZ = utils.Min(minZ, location.Z)
		maxZ = utils.Max(maxZ, location.Z)
	}

	container := Container{
		MinX:   minX - 1,
		MaxX:   maxX + 1,
		MinY:   minY - 1,
		MaxY:   maxY + 1,
		MinZ:   minZ - 1,
		MaxZ:   maxZ + 1,
		Filled: make(map[Location]int),
	}

	for location := range obsidian {
		container.Filled[location] = OBSIDIAN
	}

	currentLocation := Location{container.MinX, container.MinY, container.MinZ}

	spread(currentLocation, container)

	// Now go through each cube in the obsidian and see which ones are exposed to water, and mark the sides.
	for _, cube := range obsidian {
		// Check to see if cube shares the left side with a water cube in the container
		if container.Filled[Location{X: cube.X - 1, Y: cube.Y, Z: cube.Z}] == WATER {
			cube.CoverWithWater(LEFT)
		}
		if container.Filled[Location{X: cube.X + 1, Y: cube.Y, Z: cube.Z}] == WATER {
			cube.CoverWithWater(RIGHT)
		}
		if container.Filled[Location{X: cube.X, Y: cube.Y - 1, Z: cube.Z}] == WATER {
			cube.CoverWithWater(BOTTOM)
		}
		if container.Filled[Location{X: cube.X, Y: cube.Y + 1, Z: cube.Z}] == WATER {
			cube.CoverWithWater(TOP)
		}
		if container.Filled[Location{X: cube.X, Y: cube.Y, Z: cube.Z - 1}] == WATER {
			cube.CoverWithWater(BACK)
		}
		if container.Filled[Location{X: cube.X, Y: cube.Y, Z: cube.Z + 1}] == WATER {
			cube.CoverWithWater(FRONT)
		}
	}

	externalSurfaceArea := 0
	for _, cube := range obsidian {
		externalSurfaceArea += cube.ExposedSidesToWater()
	}

	return externalSurfaceArea
}

func spread(location Location, container Container) int {
	if location.X < container.MinX || location.X > container.MaxX || location.Y < container.MinY || location.Y > container.MaxY || location.Z < container.MinZ || location.Z > container.MaxZ {
		return 0
	}

	if _, ok := container.Filled[location]; ok {
		return 0
	}

	container.Filled[location] = WATER

	right := spread(Location{X: location.X + 1, Y: location.Y, Z: location.Z}, container)
	left := spread(Location{X: location.X - 1, Y: location.Y, Z: location.Z}, container)
	up := spread(Location{X: location.X, Y: location.Y + 1, Z: location.Z}, container)
	down := spread(Location{X: location.X, Y: location.Y - 1, Z: location.Z}, container)
	front := spread(Location{X: location.X, Y: location.Y, Z: location.Z + 1}, container)
	back := spread(Location{X: location.X, Y: location.Y, Z: location.Z - 1}, container)

	return 1 + right + left + up + down + front + back
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 18: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 18: Part 2: = %+v\n", part2Solution)
}

func getInput() Obsidian {
	lines, _ := utils.ReadLines(f, "input.txt")
	obsidian := Obsidian{}
	for _, line := range lines {
		cube := Cube{}
		cube.Init()
		fmt.Sscanf(line, "%d,%d,%d", &cube.X, &cube.Y, &cube.Z)

		obsidian[Location{cube.X, cube.Y, cube.Z}] = &cube
	}

	return obsidian
}
