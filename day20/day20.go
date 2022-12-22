package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
)

//go:embed input.txt
var f embed.FS

type Number struct {
	Value    int
	Previous *Number
	Next     *Number
}

type LinkedList struct {
	First   *Number
	Last    *Number
	Zero    *Number
	Numbers []*Number
}

func swap(left, right *Number) {
	previous := left.Previous
	next := right.Next

	next.Previous = left
	left.Previous = right
	left.Next = next
	right.Previous = previous
	right.Next = left
	previous.Next = right
}

func Part1() any {
	list := getInput()
	return mix(list, 1, 1)
}

func Part2() any {
	list := getInput()
	return mix(list, 811589153, 10)
}

func mix(list LinkedList, key int, mixAmount int) int {
	// Apply the key first
	for i := range list.Numbers {
		list.Numbers[i].Value *= key
	}

	for mixIndex := 0; mixIndex < mixAmount; mixIndex++ {
		for i := range list.Numbers {
			moveBy := utils.Abs(list.Numbers[i].Value) % (len(list.Numbers) - 1)
			value := list.Numbers[i].Value

			if moveBy > len(list.Numbers)/2 {
				moveBy = len(list.Numbers) - moveBy - 1
				value = -value
			}

			// Start moving the number
			switch {
			case value < 0:
				for swaps := 0; swaps < moveBy; swaps++ {
					swap(list.Numbers[i].Previous, list.Numbers[i])
				}
			case value > 0:
				for swaps := 0; swaps < moveBy; swaps++ {
					swap(list.Numbers[i], list.Numbers[i].Next)
				}
			}
		}
	}

	sum := 0
	current := list.Zero
	for i := 0; i <= 3000; i++ {
		if i%1000 == 0 {
			sum += current.Value
		}
		current = current.Next
	}

	return sum
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 20: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 20: Part 2: = %+v\n", part2Solution)
}

func getInput() LinkedList {
	lines, _ := utils.ReadLines(f, "input.txt")
	l := LinkedList{}

	for _, line := range lines {
		number := Number{Value: utils.ParseInt(line)}
		l.Numbers = append(l.Numbers, &number)

		if l.First == nil {
			l.First = &number
		} else if number.Value == 0 {
			l.Zero = &number
		}

		if l.Last != nil {
			number.Previous = l.Last
			l.Last.Next = &number
		}

		l.Last = &number
	}

	l.Last.Next = l.First
	l.First.Previous = l.Last

	return l
}
