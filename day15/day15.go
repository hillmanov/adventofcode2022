package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
	"math"
	"sort"
)

//go:embed input.txt
var f embed.FS

type Bounds struct {
	MinCol, MaxCol int
	MinRow, MaxRow int
}

type Sensor struct {
	Row, Col int
	Reach    int
}

func (s Sensor) BoundsAtRow(row int) (minCol int, maxCol int, inBounds bool) {
	colDelta := s.Col - (s.Col + utils.Abs(s.Row-row) - s.Reach)
	if colDelta < 0 {
		return s.Col, s.Col, false
	}
	return s.Col - colDelta, s.Col + colDelta, true
}

func Part1() any {
	sensors, _ := getInput()

	ranges := getMergedRanges(sensors, 2_000_000)
	takenSpots := 0
	for _, r := range ranges {
		takenSpots += r[1] - r[0]
	}

	return takenSpots
}

func Part2() any {
	sensors, _ := getInput()

	searchRow := func(row int, c chan int) {
		ranges := getMergedRanges(sensors, row)
		if len(ranges) == 2 { // We found the row, we can find the column by look at the ranges
			beaconRow := row
			beaconCol := ranges[0][1] + 1
			c <- beaconCol*4_000_000 + beaconRow
		}
	}

	c := make(chan int, 1)
	for row := 0; row < 4_000_000/2; row++ {
		go searchRow(row, c)
		go searchRow(4_000_000-row, c)
	}

	select {
	case <-c:
		fmt.Println("Done!")
	}

	return 1
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 15: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 15: Part 2: = %+v\n", part2Solution)
}

func getMergedRanges(sensors []Sensor, row int) [][2]int {
	// Gather all the minCol/maxCol ranges for the given row
	ranges := [][2]int{}
	for _, s := range sensors {
		minCol, maxCol, inBounds := s.BoundsAtRow(row)
		if inBounds {
			ranges = append(ranges, [2]int{minCol, maxCol})
		}
	}

	// Sort the ranges by their minCol, otherwise the merging below will not work
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	// Simplify the ranges by merging those that overlap
	// We can do this in a single pass now that the ranges have been sorted
	merged := [][2]int{}
	for _, r := range ranges {
		if len(merged) == 0 {
			merged = append(merged, r)
			continue
		}

		last := merged[len(merged)-1]
		if r[0] <= last[1]+1 {
			merged[len(merged)-1][1] = utils.Max(last[1], r[1])
		} else {
			merged = append(merged, r)
		}
	}
	return merged
}

func getInput() ([]Sensor, Bounds) {
	lines, _ := utils.ReadLines(f, "input.txt")

	sensors := []Sensor{}
	bounds := Bounds{
		MinCol: math.MaxInt,
		MaxCol: math.MinInt,
		MinRow: math.MaxInt,
		MaxRow: math.MinInt,
	}
	for _, line := range lines {
		var sensorRow, sensorCol, beaconRow, beaconCol, reach int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorCol, &sensorRow, &beaconCol, &beaconRow)

		reach = manhattanDistance(sensorCol, sensorRow, beaconCol, beaconRow)
		sensors = append(sensors, Sensor{
			Row:   sensorRow,
			Col:   sensorCol,
			Reach: reach,
		})

		bounds.MinCol = utils.Min(bounds.MinCol, sensorCol-reach)
		bounds.MaxCol = utils.Max(bounds.MaxCol, sensorCol+reach)
		bounds.MinRow = utils.Min(bounds.MinRow, sensorRow-reach)
		bounds.MaxRow = utils.Max(bounds.MaxRow, sensorRow+reach)
	}
	return sensors, bounds
}

func square(x int) float64 {
	return float64(x * x)
}

func manhattanDistance(col1, row1, col2, row2 int) int {
	return utils.Abs(row1-row2) + utils.Abs(col1-col2)
}
