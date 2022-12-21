package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var f embed.FS

type Valve struct {
	Name     string
	FlowRate int
	Mask     int
}

type Valves map[string]Valve
type Connections map[string][]string

func Part1() any {
	allValves, connections := getInput()

	valves := initializeRelevantValves(allValves)
	minutesBetween := getMinutesBetweenMap(connections, allValves)
	pressuresAchieved := map[int]int{}
	pressuresAchieved = visit("AA", 30, 0, 0, pressuresAchieved, minutesBetween, valves)

	max := 0
	for _, pressure := range pressuresAchieved {
		max = utils.Max(max, pressure)
	}

	return max
}

func Part2() any {
	allValves, connections := getInput()

	valves := initializeRelevantValves(allValves)
	minutesBetween := getMinutesBetweenMap(connections, allValves)
	pressuresAchieved := map[int]int{}
	pressuresAchieved = visit("AA", 26, 0, 0, pressuresAchieved, minutesBetween, valves)

	valveCombos := []int{}
	for valveComb := range pressuresAchieved {
		valveCombos = append(valveCombos, valveComb)
	}

	maxPressureTogetherAchieved := 0
	for i := 0; i < len(valveCombos); i++ {
		myValves := valveCombos[i]
		for j := i + 1; j < len(valveCombos); j++ {
			elephantValves := valveCombos[j]
			if myValves&elephantValves == 0 { // Make sure there are no shared valves between the two
				maxPressureTogetherAchieved = utils.Max(maxPressureTogetherAchieved, pressuresAchieved[myValves]+pressuresAchieved[elephantValves])
			}
		}
	}
	return maxPressureTogetherAchieved
}

func initializeRelevantValves(allValves Valves) Valves {
	valves := Valves{}
	i := 0
	for _, valve := range allValves {
		if valve.FlowRate > 0 {
			valve.Mask = 1 << i
			valves[valve.Name] = valve
			i++
		}
	}
	return valves
}

func getMinutesBetweenMap(connections Connections, valves Valves) map[string]int {
	minutesBetween := map[string]int{}
	for i := range valves {
		for j := range valves {
			minutesBetween[i+j] = minutesTo(connections, i, j)
		}
	}
	return minutesBetween
}

func visit(
	currentValve string,
	minutesRemaining int,
	valvesState int,
	currentFlow int,
	pressuresAchieved map[int]int,
	minutesBetween map[string]int,
	valves Valves,
) map[int]int {

	pressuresAchieved[valvesState] = utils.Max(pressuresAchieved[valvesState], currentFlow)
	for nextValve := range valves {
		newMinutesRemaining := minutesRemaining - minutesBetween[currentValve+nextValve] - 1

		if valves[nextValve].Mask&valvesState > 0 || newMinutesRemaining <= 0 { // Check to see if we have been there, and if we can make it on time
			continue
		}

		visit(
			nextValve,
			newMinutesRemaining,
			valvesState|valves[nextValve].Mask,
			currentFlow+newMinutesRemaining*valves[nextValve].FlowRate,
			pressuresAchieved,
			minutesBetween,
			valves,
		)
	}
	return pressuresAchieved
}

func minutesTo(connections Connections, start string, end string) int {
	Q := []string{start}
	visited := map[string]bool{}
	minutesBetween := map[string]int{}

	var current string
	for {
		current, Q = utils.Pop(Q)
		if !visited[current] {
			visited[current] = true
			if current == end {
				return minutesBetween[current]
			}
			for _, move := range connections[current] {
				if !visited[move] {
					Q = append(Q, move)
					minutesBetween[move] = minutesBetween[current] + 1
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
