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
		return Range{0, 0}, false
	}
	return Range{
		s.Col - colDelta,
		s.Col + colDelta,
	}, true
}

func Part1() any {
	sensors := getInput()
	ranges := getMergedRanges(sensors, 2_000_000)
	return ranges[0].End - ranges[0].Start
}

func Part2() any {
	sensors := getInput()

	result := -1
	min, max := 0, 4_000_000

	// Slightly faster, but much more complex way for finding the x
	searchRow := func(row int) {
		for x := min; x <= max; {
			sensorFound := false

			for _, s := range sensors {
				if rangeForRow, inBounds := s.RangeAtRow(row); inBounds && between(x, rangeForRow.Start, rangeForRow.End) {
					x = utils.Max(x, rangeForRow.End+1)
					sensorFound = true
				}
			}

			if !sensorFound {
				result = x*max + row
				break
			}
		}
	}

	// Search from both sides
	for row := 0; row < max/2; row++ {
		if result == -1 {
			go searchRow(row)
			go searchRow(max - row)
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
		if r, inBounds := s.RangeAtRow(row); inBounds {
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

func between(needle, start, end int) bool {
	return needle >= start && needle <= end
}
