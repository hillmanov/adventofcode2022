package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
	"sort"
	"sync"
)

//go:embed input.txt
var f embed.FS

type Square struct {
	Elevation int
	Row       int
	Col       int
}

var cache sync.Map

type Grid [][]Square

func (s Square) ValidMoves(grid Grid) []Square {
	if moves, ok := cache.Load(s); ok {
		return moves.([]Square)
	}
	moves := []Square{}
	if s.Row > 0 && grid[s.Row-1][s.Col].Elevation-s.Elevation <= 1 {
		moves = append(moves, grid[s.Row-1][s.Col])
	}
	if s.Row < len(grid)-1 && grid[s.Row+1][s.Col].Elevation-s.Elevation <= 1 {
		moves = append(moves, grid[s.Row+1][s.Col])
	}
	if s.Col > 0 && grid[s.Row][s.Col-1].Elevation-s.Elevation <= 1 {
		moves = append(moves, grid[s.Row][s.Col-1])
	}
	if s.Col < len(grid[0])-1 && grid[s.Row][s.Col+1].Elevation-s.Elevation <= 1 {
		moves = append(moves, grid[s.Row][s.Col+1])
	}

	cache.Store(s, moves)

	return moves
}

func BFS(grid Grid, start, end Square) (int, bool) {
	Q := []Square{start}
	visited := map[Square]bool{}
	distances := map[Square]int{}

	var current Square
	for {
		if len(Q) == 0 {
			return -1, false
		}

		current, Q = pop(Q)
		if !visited[current] {
			visited[current] = true
			if current == end {
				return distances[end], true
			}
			for _, move := range current.ValidMoves(grid) {
				if !visited[move] {
					Q = append(Q, move)
					distances[move] = distances[current] + 1
				}
			}
		}
	}
}

func Part1() any {
	start, end, grid := getInput()
	distance, _ := BFS(grid, start, end)

	return distance
}

func Part2() any {
	_, end, grid := getInput()

	potentialStartingPoints := []Square{}
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col].Elevation == 0 {
				potentialStartingPoints = append(potentialStartingPoints, grid[row][col])
			}
		}
	}

	distances := []int{}
	// Easy concurrency!
	wg := sync.WaitGroup{}
	for _, start := range potentialStartingPoints {
		wg.Add(1)
		go func(start Square) {
			distance, found := BFS(grid, start, end)
			if found {
				distances = append(distances, distance)
			}
			wg.Done()
		}(start)
	}
	wg.Wait()

	sort.Slice(distances, func(i, j int) bool {
		return distances[i] < distances[j]
	})

	return distances[0]
}

func main() {
	cache = sync.Map{}
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

func pop(list []Square) (Square, []Square) {
	return list[0], list[1:]
}
