package main

import (
	"adventofcode/utils"
	"container/heap"
	"embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var f embed.FS

type Position struct {
	Row int
	Col int
}

func (p *Position) ValidMoves(v Valley, minute int) []Position {
	moves := []Position{}

	here, up, down, left, right := *p, Position{p.Row - 1, p.Col}, Position{p.Row + 1, p.Col}, Position{p.Row, p.Col - 1}, Position{p.Row, p.Col + 1}

	if here == v.Start || here == v.End {
		moves = append(moves, here)
	}

	if v.PositionIsOpen(here, minute) {
		moves = append(moves, here)
	}
	if v.PositionIsOpen(up, minute) {
		moves = append(moves, up)
	}
	if v.PositionIsOpen(down, minute) {
		moves = append(moves, down)
	}
	if v.PositionIsOpen(left, minute) {
		moves = append(moves, left)
	}
	if v.PositionIsOpen(right, minute) {
		moves = append(moves, right)
	}

	if down == v.End {
		moves = append(moves, down)
	}

	if up == v.Start {
		moves = append(moves, up)
	}

	return moves
}

type PositionAndMinute struct {
	Position Position
	Minute   int
}

type Blizzard struct {
	Row       int
	Col       int
	Direction string
}

type Valley struct {
	Blizzards          map[int][]Blizzard // Blizzards move based on the current "minute". We need to be able to go "back in time" while backtracking
	BlizzardAtPosition map[int]map[Position]bool
	MinCol, MaxCol     int
	MinRow, MaxRow     int
	Start              Position
	End                Position
	Walls              map[Position]bool
}

func (v *Valley) MoveBlizzards(minute int) {
	if _, ok := v.Blizzards[minute]; ok {
		return
	}
	v.Blizzards[minute] = []Blizzard{}
	v.BlizzardAtPosition[minute] = map[Position]bool{}

	for _, blizzard := range v.Blizzards[minute-1] {
		switch blizzard.Direction {
		case "^":
			blizzard.Row--
		case "v":
			blizzard.Row++
		case ">":
			blizzard.Col++
		case "<":
			blizzard.Col--
		}

		// Wrap the blizzard around to the other side if it is out of bounds
		switch {
		case blizzard.Row <= v.MinRow-1:
			blizzard.Row = v.MaxRow - 1
		case blizzard.Row >= v.MaxRow:
			blizzard.Row = v.MinRow
		case blizzard.Col <= v.MinCol-1:
			blizzard.Col = v.MaxCol - 1
		case blizzard.Col >= v.MaxCol:
			blizzard.Col = v.MinCol
		}

		v.BlizzardAtPosition[minute][Position{Row: blizzard.Row, Col: blizzard.Col}] = true
		v.Blizzards[minute] = append(v.Blizzards[minute], blizzard)
	}
}

func (v Valley) PositionIsOpen(p Position, minute int) bool {
	return !v.BlizzardAtPosition[minute][p] && !v.Walls[p] && p.Row <= v.MaxRow && p.Col <= v.MaxCol && p.Row >= v.MinRow && p.Col >= v.MinCol
}

type Node struct {
	Position Position
	Minute   int
	Cost     int
	Open     bool
	Closed   bool
	Index    int
}

type NodeMap map[PositionAndMinute]*Node

func (nm NodeMap) Get(p PositionAndMinute) *Node {
	node, ok := nm[p]
	if !ok {
		node = &Node{Position: p.Position, Minute: p.Minute}
		nm[p] = node
	}
	return node
}

func Path(v Valley, start, end Position, minute int) (time int) {
	nodeMap := NodeMap{}
	pQ := &PriorityQueue{}
	heap.Init(pQ)

	fromNode := nodeMap.Get(PositionAndMinute{Position: start, Minute: minute})
	fromNode.Open = true
	heap.Push(pQ, fromNode)

	for {
		if pQ.Len() == 0 {
			panic("No path found!")
		}

		current := heap.Pop(pQ).(*Node)
		current.Open = false
		current.Closed = true

		if current.Position == end {
			return current.Minute
		}

		v.MoveBlizzards(current.Minute + 1)
		for _, neighbor := range current.Position.ValidMoves(v, current.Minute+1) {
			cost := current.Cost + 1 + current.Minute

			neighborNode := nodeMap.Get(PositionAndMinute{Position: neighbor, Minute: current.Minute + 1})

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
				heap.Push(pQ, neighborNode)
			}
		}
	}
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 12: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 12: Part 2: = %+v\n", part2Solution)
}

func Part1() any {
	valley := getInput()
	minutes := Path(valley, valley.Start, valley.End, 0)
	return minutes
}

func Part2() any {
	valley := getInput()
	there := Path(valley, valley.Start, valley.End, 0)
	back := Path(valley, valley.End, valley.Start, there)
	thereAgain := Path(valley, valley.Start, valley.End, back)

	return thereAgain
}

func getInput() Valley {
	lines, _ := utils.ReadLines(f, "input.txt")

	valley := Valley{
		Blizzards:          map[int][]Blizzard{},
		BlizzardAtPosition: map[int]map[Position]bool{},
		Walls:              map[Position]bool{},
		MinCol:             1,
		MaxCol:             len(lines[0]) - 1,
		MinRow:             1,
		MaxRow:             len(lines) - 1,
		Start: Position{
			Row: 0,
			Col: 1,
		},
		End: Position{
			Row: len(lines) - 1,
			Col: len(lines[0]) - 2,
		},
	}

	for rowIndex := range lines {
		for colIndex := range lines[rowIndex] {
			symbol := lines[rowIndex][colIndex]
			switch {
			case utils.IndexOf(strings.Split("^v<>", ""), string(symbol)) > -1:
				valley.Blizzards[0] = append(valley.Blizzards[0], Blizzard{
					Row:       rowIndex,
					Col:       colIndex,
					Direction: string(symbol),
				})
			case symbol == '#':
				valley.Walls[Position{
					Row: rowIndex,
					Col: colIndex,
				}] = true
			}
		}
	}

	return valley
}
