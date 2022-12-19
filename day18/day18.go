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
	X, Y, Z int
	Exposed uint
}

func (c *Cube) Init() {
	c.Exposed = TOP | BOTTOM | LEFT | RIGHT | FRONT | BACK
}

func (c *Cube) Cover(side uint) {
	c.Exposed &= ^side
}

func (c Cube) ExposedSides() int {
	return bits.OnesCount(c.Exposed)
}

type Container struct {
	MinX   int
	MaxX   int
	MinY   int
	MaxY   int
	MinZ   int
	MaxZ   int
	Filled map[Location]bool
}

func (c Container) ExpandedVolume() int {
	return (c.MaxX - c.MinX) * (c.MaxY - c.MinY) * (c.MaxZ - c.MinZ)
}

func (c Container) InnerVolume() int {
	return ((c.MaxX - c.MinX) - 2) * ((c.MaxY - c.MinY) - 2) * ((c.MaxZ - c.MinZ) - 2)
}

func (c Container) GetVolumeDiff() int {
	return c.ExpandedVolume() - c.InnerVolume()
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
	cubes := getInput()
	return getTotalSurfaceArea(cubes)
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

	fmt.Printf("minX: %+v, maxX: %+v, minY: %+v, maxY: %+v, minZ: %+v, maxZ: %+v\n", minX, maxX, minY, maxY, minZ, maxZ)

	container := Container{
		MinX:   minX - 1,
		MaxX:   maxX + 1,
		MinY:   minY - 1,
		MaxY:   maxY + 1,
		MinZ:   minZ - 1,
		MaxZ:   maxZ + 1,
		Filled: make(map[Location]bool),
	}

	for location := range obsidian {
		container.Filled[location] = true
	}

	var currentLocation Location
findStartingLocation:
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			for z := minZ; z <= maxZ; z++ {
				if _, ok := obsidian[Location{x, y, z}]; !ok {
					currentLocation = Location{x, y, z}
					break findStartingLocation
				}
			}
		}
	}

	fmt.Printf("startingLocation = %+v\n", currentLocation)

	volume := spread(currentLocation, container) - container.GetVolumeDiff()
	fmt.Printf("volume = %+v\n", volume)

	fmt.Printf("container.ExpandedVolume() = %+v\n", container.ExpandedVolume())
	fmt.Printf("container.InnerVolumne() = %+v\n", container.InnerVolume())

	totalSurfaceArea := getTotalSurfaceArea(obsidian)

	// NOT 3250 - that is too high, 2008 is too low
	return container.InnerVolume() - totalSurfaceArea
}

func spread(location Location, container Container) int {
	if location.X < container.MinX || location.X > container.MaxX || location.Y < container.MinY || location.Y > container.MaxY || location.Z < container.MinZ || location.Z > container.MaxZ {
		return 0
	}
	if _, ok := container.Filled[location]; ok {
		return 0
	}
	container.Filled[location] = true
	right := spread(Location{location.X + 1, location.Y, location.Z}, container)
	left := spread(Location{location.X - 1, location.Y, location.Z}, container)
	up := spread(Location{location.X, location.Y + 1, location.Z}, container)
	down := spread(Location{location.X, location.Y - 1, location.Z}, container)
	front := spread(Location{location.X, location.Y, location.Z + 1}, container)
	back := spread(Location{location.X, location.Y, location.Z - 1}, container)

	return 1 + right + left + up + down + front + back
}

func getTotalSurfaceArea(obsidian Obsidian) int {
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
