package main

import (
	"adventofcode/utils"
	"embed"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var f embed.FS

type Num = float64
type Packet = any
type Pair = [2]Packet
type List = []Packet

func Part1() any {
	pairs := getInput()

	sum := 0
	for i, pair := range pairs {
		if isEqual, _ := isOrdered(pair[0], pair[1]); isEqual {
			sum += i + 1
		}
	}

	return sum
}

func Part2() any {
	pairs := getInput()

	packets := []Packet{}
	for _, pair := range pairs {
		packets = append(packets, pair[0])
		packets = append(packets, pair[1])
	}

	divider1 := List{List{float64(2)}}
	divider2 := List{List{float64(6)}}

	packets = append(packets, divider1)
	packets = append(packets, divider2)

	sort.Slice(packets, func(i, j int) bool {
		ordered, _ := isOrdered(packets[i], packets[j])
		return ordered
	})

	divider1Index := -1
	divider2Index := -1
	for i, packet := range packets {
		if fmt.Sprintf("%v", packet) == fmt.Sprintf("%v", divider1) {
			divider1Index = i + 1
		}
		if fmt.Sprintf("%v", packet) == fmt.Sprintf("%v", divider2) {
			divider2Index = i + 1
		}
	}

	return divider1Index * divider2Index
}

func isOrdered(a, b interface{}) (bool, bool) {
	// Try to cast to see what we are working with
	aNum, aIsNumber := a.(Num)
	bNum, bIsNumber := b.(Num)
	aList, aIsList := a.(List)
	bList, bIsList := b.(List)

	ordered := true
	shouldContinue := true
	switch {
	case aIsList && bIsList:
		for i := 0; i < utils.Max(len(aList), len(bList)) && ordered && shouldContinue; i++ {
			if ordered && shouldContinue && (i >= len(aList) || i >= len(bList)) {
				return len(aList) <= len(bList), false
			}
			ordered, shouldContinue = isOrdered(aList[i], bList[i])
		}

		return ordered, shouldContinue

	case aIsNumber && bIsNumber:
		switch {
		case aNum < bNum:
			return true, false
		case aNum == bNum:
			return true, true
		case aNum > bNum:
			return false, false
		}

	case aIsNumber && bIsList:
		return isOrdered(List{aNum}, b)

	case aIsList && bIsNumber:
		return isOrdered(a, List{bNum})
	}

	return true, false
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 13: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 13: Part 2: = %+v\n", part2Solution)
}

func getInput() []Pair {
	contents, _ := utils.ReadContents(f, "input.txt")
	pairs := []Pair{}

	sections := strings.Split(contents, "\n\n")
	for _, section := range sections {
		pair := Pair{}
		parts := strings.Split(section, "\n")
		json.Unmarshal([]byte(parts[0]), &pair[0])
		json.Unmarshal([]byte(parts[1]), &pair[1])
		pairs = append(pairs, pair)
	}

	return pairs
}
