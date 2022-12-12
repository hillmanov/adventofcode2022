package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strings"
)

type Monkey struct {
	ID          int
	Items       []int
	Operation   func(a int) int
	TestDivisor int
	TrueMonkey  int
	FalseMonkey int
}

//go:embed input.txt
var f embed.FS

func Part1() any {
	monkeys := getInput()
	inspectionCount := make([]int, len(monkeys))

	for i := 0; i < 20; i++ {
		for m := range monkeys {
			monkey := monkeys[m]

			for len(monkey.Items) > 0 {
				inspectionCount[m]++
				item := monkey.Items[0]
				monkey.Items = monkey.Items[1:]
				worryLevel := int(float64(monkey.Operation(item))) / 3
				if worryLevel%monkey.TestDivisor == 0 {
					monkeys[monkey.TrueMonkey].Items = append(monkeys[monkey.TrueMonkey].Items, worryLevel)
				} else {
					monkeys[monkey.FalseMonkey].Items = append(monkeys[monkey.FalseMonkey].Items, worryLevel)
				}
			}
		}
	}

	sort.Slice(inspectionCount, func(i, j int) bool {
		return inspectionCount[i] > inspectionCount[j]
	})

	return inspectionCount[0] * inspectionCount[1]
}

func Part2() any {
	monkeys := getInput()
	inspectionCount := make([]int, len(monkeys))

	for i := 0; i < 20; i++ {
		for m := range monkeys {
			monkey := monkeys[m]

			for len(monkey.Items) > 0 {
				inspectionCount[m]++
				item := monkey.Items[0]
				monkey.Items = monkey.Items[1:]

				worryLevel := int(float64(monkey.Operation(item)))
				worryLevel = int(math.Logb(float64(worryLevel)))

				if worryLevel%monkey.TestDivisor == 0 {
					monkeys[monkey.TrueMonkey].Items = append(monkeys[monkey.TrueMonkey].Items, item)
				} else {
					monkeys[monkey.FalseMonkey].Items = append(monkeys[monkey.FalseMonkey].Items, item)
				}
			}
		}
	}

	sort.Slice(inspectionCount, func(i, j int) bool {
		return inspectionCount[i] > inspectionCount[j]
	})

	return nil
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 11: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 11: Part 2: = %+v\n", part2Solution)
}

func getInput() []*Monkey {
	isNumber := regexp.MustCompile(`^\d+$`)
	contents, _ := utils.ReadContents(f, "input.txt")

	monkeys := []*Monkey{}

	sections := strings.Split(contents, "\n\n")
	for _, section := range sections {
		lines := strings.Split(section, "\n")
		monkey := Monkey{}

		fmt.Sscanf(strings.Trim(lines[0], " "), "Monkey %d:", &monkey.ID)
		monkey.Items = ExtractItems(strings.TrimPrefix(strings.TrimLeft(lines[1], " "), "Starting items: "))

		var argA, operator, argB string
		fmt.Sscanf(strings.Trim(lines[2], " "), "Operation: new = %s %s %s", &argA, &operator, &argB)
		switch {
		case argA == "old" && isNumber.MatchString(argB):
			switch operator {
			case "+":
				monkey.Operation = func(old int) int {
					return old + utils.ParseInt(argB)
				}
			case "*":
				monkey.Operation = func(old int) int {
					return old * utils.ParseInt(argB)
				}
			}
		case argA == "old" && argB == "old":
			switch operator {
			case "+":
				monkey.Operation = func(old int) int {
					return old + old
				}
			case "*":
				monkey.Operation = func(old int) int {
					return old * old
				}
			}
		}

		fmt.Sscanf(strings.Trim(lines[3], " "), "Test: divisible by %d:", &monkey.TestDivisor)
		fmt.Sscanf(strings.Trim(lines[4], " "), "If true: throw to monkey %d", &monkey.TrueMonkey)
		fmt.Sscanf(strings.Trim(lines[5], " "), "If false: throw to monkey %d", &monkey.FalseMonkey)

		monkeys = append(monkeys, &monkey)
	}

	return monkeys
}

func ExtractItems(strItems string) []int {
	items := []int{}
	for _, strItem := range strings.Split(strItems, ", ") {
		items = append(items, utils.ParseInt(strItem))
	}
	return items
}
