package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed input.txt
var f embed.FS

type Point struct {
	Row, Col int
}

type PointWithDirection struct {
	Row, Col  int
	Direction Delta
}

type Delta struct {
	Row, Col int
}

type Movement struct {
	Direction rune
	Amount    int
}

type Grove struct {
	MaxRow, MaxCol     int
	StartRow, StartCol int
	Map                map[Point]rune
}

func (g Grove) String(pathFollowed PathFollowed) string {
	var s strings.Builder
	s.WriteRune('\n')
	for row := 0; row < g.MaxRow; row++ {
		for col := 0; col < g.MaxCol; col++ {
			if r, ok := g.Map[Point{row, col}]; ok {
				if row == g.StartRow && col == g.StartCol {
					s.WriteRune('@')
				} else {
					// Get path followed element
					runeToWrite := r
					for i, p := range pathFollowed {

						if p.Row == row && p.Col == col {
							if i == len(pathFollowed)-1 {
								runeToWrite = 'X'
							} else {
								switch p.Direction {
								case UP:
									runeToWrite = '^'
								case DOWN:
									runeToWrite = 'v'
								case LEFT:
									runeToWrite = '<'
								case RIGHT:
									runeToWrite = '>'
								}
							}
						}
					}

					s.WriteRune(runeToWrite)
				}
			} else {
				s.WriteRune(' ')
			}
		}
		s.WriteRune('\n')
	}
	return s.String()
}

func (g Grove) GetWrapAroundRow(row int, col int, delta int) int {
	for {
		if r, ok := g.Map[Point{row, col}]; ok {
			if r == '#' || r == '.' {
				return row
			}
		}
		row += delta
		if row > g.MaxRow {
			row = 0
		}
		if row < 0 {
			row = g.MaxRow
		}
	}
}

func (g Grove) GetWrapAroundCol(row int, col int, delta int) int {
	for {
		if r, ok := g.Map[Point{row, col}]; ok {
			if r == '#' || r == '.' {
				return col
			}
		}
		col += delta
		if col > g.MaxCol {
			col = 0
		}
		if col < 0 {
			col = g.MaxCol
		}
	}
}

type Path []Movement
type PathFollowed []PointWithDirection

var (
	VERTICAL   = 0
	HORIZONTAL = 1
	VOID       = new(rune)
	UP         = Delta{Row: -1, Col: 0}
	DOWN       = Delta{Row: 1, Col: 0}
	LEFT       = Delta{Row: 0, Col: -1}
	RIGHT      = Delta{Row: 0, Col: 1}
)

func (d Delta) Turn(direction rune) Delta {
	switch d {
	case UP:
		switch direction {
		case 'L':
			return LEFT
		case 'R':
			return RIGHT
		}
	case DOWN:
		switch direction {
		case 'L':
			return RIGHT
		case 'R':
			return LEFT
		}
	case LEFT:
		switch direction {
		case 'L':
			return DOWN
		case 'R':
			return UP
		}
	case RIGHT:
		switch direction {
		case 'L':
			return UP
		case 'R':
			return DOWN
		}
	}
	panic("Invalid direction")
}

func (d Delta) Plane() int {
	switch d {
	case UP, DOWN:
		return VERTICAL
	case LEFT, RIGHT:
		return HORIZONTAL
	}
	panic("Invalid direction")
}

func Part1() any {
	grove, path := getInput()

	currentPoint := Point{Row: grove.StartRow, Col: grove.StartCol}
	movementDelta := UP

	for _, m := range path {

		movementDelta = movementDelta.Turn(m.Direction)

		for step := 0; step < m.Amount; step++ {
			nextPoint := Point{
				Row: currentPoint.Row + movementDelta.Row,
				Col: currentPoint.Col + movementDelta.Col,
			}
			nextTile := grove.Map[nextPoint]

			if movementDelta.Plane() == VERTICAL && (nextPoint.Row < 0 || nextPoint.Row >= grove.MaxRow || nextTile == ' ' || nextTile == *VOID) {
				nextPoint.Row = grove.GetWrapAroundRow(nextPoint.Row, nextPoint.Col, movementDelta.Row)
			}

			if movementDelta.Plane() == HORIZONTAL && (nextPoint.Col < 0 || nextPoint.Col >= grove.MaxCol || nextTile == ' ' || nextTile == *VOID) {
				nextPoint.Col = grove.GetWrapAroundCol(nextPoint.Row, nextPoint.Col, movementDelta.Col)
			}

			if grove.Map[nextPoint] == '#' {
				nextPoint = currentPoint
			}

			currentPoint = nextPoint
		}
	}

	sum := (currentPoint.Row + 1) * 1000
	sum += (currentPoint.Col + 1) * 4
	switch movementDelta {
	case UP:
		sum += 3
	case DOWN:
		sum += 1
	case LEFT:
		sum += 2
	case RIGHT:
		sum += 0
	}

	return sum
}

func Part2() any {
	return nil
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 22: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 22: Part 2: = %+v\n", part2Solution)
}

func getInput() (Grove, Path) {
	contents, _ := utils.ReadContents(f, "input.txt")

	grove := Grove{
		Map: map[Point]rune{},
	}
	parts := strings.Split(contents, "\n\n")

	rows := strings.Split(parts[0], "\n")
	for row := range rows {
		for col := range strings.Split(rows[row], "") {
			v := rune(rows[row][col])
			if v != ' ' && grove.StartCol == 0 && grove.StartRow == 0 {
				grove.StartRow = row
				grove.StartCol = col
			}
			grove.Map[Point{Row: row, Col: col}] = v
			grove.MaxRow = utils.Max(grove.MaxRow, row+1)
			grove.MaxCol = utils.Max(grove.MaxCol, col+1)
		}
	}

	// Start by going to the right. Easier to initialize here
	parts[1] = "R" + parts[1]

	splitter := regexp.MustCompile(`\d+|\w`)
	pathParts := splitter.FindAllStringSubmatch(parts[1], -1)

	path := Path{}
	for i := 0; i < len(pathParts)-1; i += 2 {
		path = append(path, Movement{
			Direction: rune(pathParts[i][0][0]),
			Amount:    utils.ParseInt(pathParts[i+1][0]),
		})
	}

	return grove, path
}
