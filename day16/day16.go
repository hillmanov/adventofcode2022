package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
	"math"
	"sort"
	"strings"
)

//go:embed input.txt
var f embed.FS
var max = math.MinInt
var seen = map[int]bool{}

type Valve struct {
	Name           string
	FlowRate       int
	Open           bool
	OpenedAtMinute int
}

type Move struct {
	Valve         string
	MinutesAway   int
	TotalPressure int
}

func (v Valve) TotalPressure() int {
	if v.Open {
		return v.FlowRate * (30 - v.OpenedAtMinute)
	}
	return 0
}

type Connections map[string][]string
type Valves map[string]Valve

func Part1() any {
	valves, connections := getInput()
	minutes := 1

	currentState := "AA"

	for minutes <= 30 {
		if currentState != "AA" {
			newState := valves[currentState]
			newState.Open = true
			newState.OpenedAtMinute = minutes
			valves[currentState] = newState
			minutes++
		}

		moves := []Move{}

		for _, valve := range valves {
			if valve.Name != currentState && !valve.Open { // No need to look at valves that are already open
				move := Move{}
				move.Valve = valve.Name
				minutesAway, _ := minutesTo(connections, currentState, valve.Name)
				move.MinutesAway = minutesAway
				move.TotalPressure = valve.FlowRate * (30 - minutes - minutesAway)
				moves = append(moves, move)
			}
		}

		sort.Slice(moves, func(i, j int) bool {
			return moves[i].TotalPressure > moves[j].TotalPressure
		})

		if len(moves) > 0 {
			bestMove := moves[0]
			currentState = bestMove.Valve
			minutes += bestMove.MinutesAway
		} else {
			minutes++
		}
	}

	sum := 0
	for _, valve := range valves {
		fmt.Printf("valve = %+v\n", valve)
		sum += valve.TotalPressure()
	}
	fmt.Printf("sum = %+v\n", sum)

	return nil
}

func Part2() any {
	return nil
}

func minutesTo(connections Connections, start string, end string) (int, bool) {
	Q := []string{start}
	visited := map[string]bool{}
	distances := map[string]int{}

	var current string
	for {
		if len(Q) == 0 {
			return -1, false
		}

		current, Q = utils.Pop(Q)
		if !visited[current] {
			visited[current] = true
			if current == end {
				return distances[current], true
			}
			for _, move := range connections[current] {
				if !visited[move] {
					Q = append(Q, move)
					distances[move] = distances[current] + 1
				}
			}
		}
	}
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 16: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 16: Part 2: = %+v\n", part2Solution)
}

func getInput() (Valves, Connections) {
	lines, _ := utils.ReadLines(f, "input.txt")

	valves := Valves{}
	connections := Connections{}

	for _, line := range lines {
		valve := Valve{}
		parts := strings.Split(line, ";")

		fmt.Sscanf(parts[0], "Valve %s has flow rate=%d", &valve.Name, &valve.FlowRate)
		tunnels := strings.Split(strings.SplitN(parts[1], " ", 6)[5], ", ")

		connections[valve.Name] = tunnels
		valves[valve.Name] = valve
	}

	return valves, connections
}
