package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
	"sync"
)

//go:embed input.txt
var f embed.FS

type Square struct {
	Elevation int
	Row       int
	Col       int
}

type moveValidatorFunc func(int, int) bool

type terminatorFunc func(Square) bool

var cache sync.Map

type Grid [][]Square

func (s Square) ValidMoves(grid Grid, validateMove moveValidatorFunc) []Square {
	if moves, ok := cache.Load(s); ok {
		return moves.([]Square)
	}
	moves := []Square{}
	if s.Row > 0 && validateMove(grid[s.Row-1][s.Col].Elevation, s.Elevation) {
		moves = append(moves, grid[s.Row-1][s.Col])
	}
	if s.Row < len(grid)-1 && validateMove(grid[s.Row+1][s.Col].Elevation, s.Elevation) {
		moves = append(moves, grid[s.Row+1][s.Col])
	}
	if s.Col > 0 && validateMove(grid[s.Row][s.Col-1].Elevation, s.Elevation) {
		moves = append(moves, grid[s.Row][s.Col-1])
	}
	if s.Col < len(grid[0])-1 && validateMove(grid[s.Row][s.Col+1].Elevation, s.Elevation) {
		moves = append(moves, grid[s.Row][s.Col+1])
	}

	cache.Store(s, moves)

	return moves
}

func BFS(grid Grid, start Square, terminator terminatorFunc, validateMove moveValidatorFunc) (int, bool) {
	Q := []Square{start}
	visited := map[Square]bool{}
	distances := map[Square]int{}

	var current Square
	for {
		if len(Q) == 0 {
			return -1, false
		}

		current, Q = utils.Pop(Q)
		if !visited[current] {
			visited[current] = true
			if terminator(current) {
				return distances[current], true
			}
			for _, move := range current.ValidMoves(grid, validateMove) {
				if !visited[move] {
					Q = append(Q, move)
					distances[move] = distances[current] + 1
				}
			}
		}
	}
}

func Part1() any {
	cache = sync.Map{}
	start, end, grid := getInput()

	terminator := func(s Square) bool {
		return s == end
	}

	distance, _ := BFS(grid, start, terminator, up)

	return distance
}

func Part2() any {
	cache = sync.Map{}
	_, end, grid := getInput()
	terminator := func(s Square) bool {
		return grid[s.Row][s.Col].Elevation == 0
	}
	distance, _ := BFS(grid, end, terminator, down)

	return distance
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 12: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 12: Part 2: = %+v\n", part2Solution)
}

func getInput() (start Square, end Square, grid Grid) {
	lines, _ := utils.ReadLines(f, "input.txt")

	for row, line := range lines {
		grid = append(grid, []Square{})
		for col, char := range line {
			switch char {
			case 'S':
				start = Square{Elevation: int('a') - 97, Row: row, Col: col}
				grid[row] = append(grid[row], start)
			case 'E':
				end = Square{Elevation: int('z') - 97, Row: row, Col: col}
				grid[row] = append(grid[row], end)
			default:
				grid[row] = append(grid[row], Square{Elevation: int(char) - 97, Row: row, Col: col})
			}
		}
	}

	return start, end, grid
}

func up(a, b int) bool {
	return a <= b+1
}
func down(a, b int) bool {
	return b <= a+1
}
