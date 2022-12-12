package main

import (
	"adventofcode/utils"
	"container/heap"
	"embed"
	"fmt"
	"sort"
)

//go:embed input.txt
var f embed.FS

type Square struct {
	Elevation int
	Row       int
	Col       int
}

type Grid [][]Square

func (s *Square) ValidMoves(grid Grid) []Square {
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

	return moves
}

func (s *Square) MovementCost(grid Grid, to Square) int {
	return 1 // Movement always costs 1 (For part 1)
}

func (s *Square) EstimatedCost(grid Grid, to Square) int {
	return grid[to.Row][to.Col].Elevation - s.Elevation // This can be just about anything
}

type Node struct {
	Square Square
	Cost   int
	Rank   int
	Parent *Node
	Open   bool
	Closed bool
	Index  int // Used by the heap
}

type NodeMap map[Square]*Node

func (nm NodeMap) Get(square Square) *Node {
	node, ok := nm[square]
	if !ok {
		node = &Node{Square: square}
		nm[square] = node
	}
	return node
}

type PriorityQueue []*Node

func (pQ PriorityQueue) Len() int {
	return len(pQ)
}

func (pQ PriorityQueue) Less(i, j int) bool {
	return pQ[i].Rank < pQ[j].Rank
}

func (pQ PriorityQueue) Swap(i, j int) {
	pQ[i], pQ[j] = pQ[j], pQ[i]
	pQ[i].Index = i
	pQ[j].Index = j
}

func (pQ *PriorityQueue) Push(x any) {
	n := len(*pQ)
	no := x.(*Node)
	no.Index = n
	*pQ = append(*pQ, no)
}

func (pQ *PriorityQueue) Pop() any {
	old := *pQ
	oldLength := len(old)
	last := old[oldLength-1]
	last.Index = -1
	*pQ = old[0 : oldLength-1]
	return last
}

func Path(grid Grid, from Square, to Square) (path []Square, distance int, pathFound bool) {
	nodeMap := NodeMap{}
	pQ := &PriorityQueue{}
	heap.Init(pQ)

	fromNode := nodeMap.Get(from)
	fromNode.Open = true
	heap.Push(pQ, fromNode)

	for {
		if pQ.Len() == 0 {
			return path, distance, false
		}

		current := heap.Pop(pQ).(*Node)
		current.Open = false
		current.Closed = true

		if current == nodeMap.Get(to) {
			// We found our path!
			path := []Square{}
			c := current
			for c != nil {
				path = append(path, c.Square)
				c = c.Parent
			}
			return path, current.Cost, true
		}

		for _, neighbor := range current.Square.ValidMoves(grid) {
			cost := current.Cost + current.Square.MovementCost(grid, neighbor)
			neighborNode := nodeMap.Get(neighbor)
			if cost < neighborNode.Cost {
				if neighborNode.Open {
					heap.Remove(pQ, neighborNode.Index)
				}
				neighborNode.Open = false
				neighborNode.Closed = false
			}
			if !neighborNode.Open && !neighborNode.Closed {
				neighborNode.Cost = cost
				neighborNode.Open = true
				neighborNode.Rank = cost + neighbor.EstimatedCost(grid, to)
				neighborNode.Parent = current
				heap.Push(pQ, neighborNode)
			}
		}
	}
}

func Part1() any {
	start, end, grid := getInput()
	_, distance, _ := Path(grid, start, end)

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
	for _, start := range potentialStartingPoints {
		_, distance, found := Path(grid, start, end)
		if found {
			distances = append(distances, distance)
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i] < distances[j]
	})

	return distances[0]
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
