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

type Tile int

const (
	AIR Tile = iota
	ROCK
	SAND
)

type Point struct {
	Y, X int
}

type RockPath struct {
	Start Point
	End   Point
}

type Cave struct {
	Map        map[Point]Tile
	Normalizer int
	MaxY       int
	MinX       int
	MaxX       int
}

func (c *Cave) GetTile(p Point) Tile {
	return c.Map[p]
}

func (c Cave) Dump() {
	minX := math.MaxInt
	maxX := math.MinInt
	for p := range c.Map {
		minX = utils.Min(minX, p.X)
		maxX = utils.Max(maxX, p.X)
	}

	fmt.Println("--------------------")
	for row := 0; row <= c.MaxY; row++ {
		for col := minX; col < maxX; col++ {
			tile, ok := c.Map[Point{Y: row, X: col}]
			if !ok {
				fmt.Print(".")
			} else {
				switch tile {
				case ROCK:
					fmt.Print("#")
				case AIR:
					fmt.Print(".")
				case SAND:
					fmt.Print("o")
				}
			}
		}
		fmt.Println()
	}
}

func simulateSand(cave Cave, start Point) (resting Point, inCave bool) {
	current := start
	for {
		if current.Y == cave.MaxY-1 {
			return current, false
		}
		if cave.GetTile(Point{Y: current.Y + 1, X: current.X}) == AIR {
			current.Y++
		} else {
			break
		}
	}

	// Look left
	if current.Y+1 == cave.MaxY || current.X-1 == cave.MinX-1 {
		return current, false
	}
	if cave.GetTile(Point{Y: current.Y + 1, X: current.X - 1}) == AIR {
		current.Y++
		current.X--
		return simulateSand(cave, current)
	}

	// Look right
	if current.Y+1 == cave.MaxY || current.X+1 == cave.MaxX {
		return current, false
	}
	if cave.GetTile(Point{Y: current.Y + 1, X: current.X + 1}) == AIR {
		current.Y++
		current.X++
		return simulateSand(cave, current)
	}

	return current, true
}

func Part1() any {
	cave := getInput()

	i := 0
	for {
		restingPoint, inCave := simulateSand(cave, Point{Y: 0, X: 500 - cave.Normalizer})
		if !inCave {
			break
		}
		cave.Map[restingPoint] = SAND
		i++
	}

	return i
}

func Part2() any {
	cave := getInput()

	// Setup the floor, and infinite sides
	cave.MaxX = math.MaxInt
	cave.MinX = math.MinInt
	cave.MaxY = cave.MaxY + 1

	sandOrigin := Point{Y: 0, X: 500 - cave.Normalizer}

	i := 0
	for {
		restingPoint, _ := simulateSand(cave, sandOrigin)
		cave.Map[restingPoint] = SAND
		i++
		if restingPoint == sandOrigin {
			break
		}
	}

	return i
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 14: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 14: Part 2: = %+v\n", part2Solution)
}

func getInput() Cave {
	lines, _ := utils.ReadLines(f, "input.txt")

	cave := Cave{
		Map: make(map[Point]Tile),
	}

	// Gather rock paths
	rockPaths := []RockPath{}
	for _, line := range lines {
		rockPathSections := strings.Split(line, " -> ")
		for i := 0; i < len(rockPathSections)-1; i++ {
			startParts := strings.Split(rockPathSections[i], ",")
			endParts := strings.Split(rockPathSections[i+1], ",")
			rockPath := RockPath{
				Start: Point{
					X: utils.ParseInt(startParts[0]),
					Y: utils.ParseInt(startParts[1]),
				},
				End: Point{
					X: utils.ParseInt(endParts[0]),
					Y: utils.ParseInt(endParts[1]),
				},
			}
			rockPaths = append(rockPaths, rockPath)
		}
	}

	// Find min/max of X Y
	maxX := math.MinInt
	minX := math.MaxInt
	maxY := math.MinInt
	for _, rockPath := range rockPaths {
		minX = utils.MinOf([]int{minX, rockPath.Start.X, rockPath.End.X})
		maxX = utils.MaxOf([]int{maxX, rockPath.Start.X, rockPath.End.X})
		maxY = utils.MaxOf([]int{maxY, rockPath.Start.Y, rockPath.End.Y})
	}

	maxX = maxX - minX

	cave.MaxY = maxY + 1
	cave.MinX = 0
	cave.MaxX = maxX + 1
	cave.Normalizer = minX

	// Normalize Xs
	for i := range rockPaths {
		rockPaths[i].Start.X -= minX
		rockPaths[i].End.X -= minX
	}

	for _, rockPath := range rockPaths {
		current := rockPath.Start
		target := rockPath.End

		deltaX := getDelta(target.X - current.X)
		deltaY := getDelta(target.Y - current.Y)

		cave.Map[current] = ROCK
		for current.X != target.X || current.Y != target.Y {
			current.X += deltaX
			current.Y += deltaY
			cave.Map[current] = ROCK
		}
	}

	return cave
}

func getDelta(v int) int {
	if v == 0 {
		return 0
	}
	if v >= 1 {
		return 1
	}
	return -1
}
