package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
	"sort"
)

//go:embed input.txt
var f embed.FS

type Range struct {
	Start, End int
}

type Sensor struct {
	Row, Col int
	Reach    int
}

func (s Sensor) RangeAtRow(row int) (r Range, inBounds bool) {
	colDelta := s.Col - (s.Col + utils.Abs(s.Row-row) - s.Reach)
	if colDelta < 0 {
		return Range{s.Col, s.Col}, false
	}
	return Range{s.Col - colDelta, s.Col + colDelta}, true
}

func Part1() any {
	sensors := getInput()

	ranges := getMergedRanges(sensors, 2_000_000)
	takenSpots := 0
	for _, r := range ranges {
		takenSpots += r.End - r.Start
	}

	return takenSpots
}

func Part2() any {
	sensors := getInput()

	result := -1
	searchRow := func(row int) {
		ranges := getMergedRanges(sensors, row)
		if len(ranges) == 2 { // We found the row with the gap, we can find the column by looking at the ranges
			result = ((ranges[0].End + 1) * 4_000_000) + row
		}
	}

	// Search from both sides
	for row := 0; row < 4_000_000/2; row++ {
		if result == -1 {
			go searchRow(row)
			go searchRow(4_000_000 - row)
		}
	}

	return result
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 15: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 15: Part 2: = %+v\n", part2Solution)
}

func getMergedRanges(sensors []Sensor, row int) []Range {
	// Gather all the minCol/maxCol ranges for the given row
	ranges := []Range{}
	for _, s := range sensors {
		r, inBounds := s.RangeAtRow(row)
		if inBounds {
			ranges = append(ranges, r)
		}
	}

	// Sort the ranges by their minCol, otherwise the merging below will not work
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	// Simplify the ranges by merging those that overlap
	// We can do this in a single pass now that the ranges have been sorted
	merged := []Range{}
	for _, r := range ranges {
		if len(merged) == 0 {
			merged = append(merged, r)
			continue
		}

		last := merged[len(merged)-1]
		if r.Start <= last.End+1 {
			merged[len(merged)-1].End = utils.Max(last.End, r.End)
		} else {
			merged = append(merged, r)
		}
	}
	return merged
}

func getInput() []Sensor {
	lines, _ := utils.ReadLines(f, "input.txt")

	sensors := []Sensor{}
	for _, line := range lines {
		var sensorRow, sensorCol, beaconRow, beaconCol, reach int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorCol, &sensorRow, &beaconCol, &beaconRow)

		reach = manhattanDistance(sensorCol, sensorRow, beaconCol, beaconRow)
		sensors = append(sensors, Sensor{
			Row:   sensorRow,
			Col:   sensorCol,
			Reach: reach,
		})
	}
	return sensors
}

func manhattanDistance(col1, row1, col2, row2 int) int {
	return utils.Abs(row1-row2) + utils.Abs(col1-col2)
}
