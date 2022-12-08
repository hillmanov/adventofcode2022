package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var f embed.FS

type location struct {
	row, col int
}

func Part1() any {
	forest := getInput()
	visibleTrees := map[location]bool{}

	// Visible from top
	for col := 0; col < len(forest[0]); col++ {
		tallestTreeInColSoFar := -1
		for row := 0; row < len(forest); row++ {
			if forest[row][col] > tallestTreeInColSoFar {
				visibleTrees[location{row, col}] = true
				tallestTreeInColSoFar = forest[row][col]
			}
			if forest[row][col] == 9 {
				break
			}
		}
	}

	// Visible from bottom
	for col := 0; col < len(forest); col++ {
		tallestTreeInColSoFar := -1
		for row := len(forest[0]) - 1; row >= 0; row-- {
			if forest[row][col] > tallestTreeInColSoFar {
				visibleTrees[location{row, col}] = true
				tallestTreeInColSoFar = forest[row][col]
			}
			if forest[row][col] == 9 {
				break
			}
		}
	}

	// Visible from left
	for row := 0; row < len(forest); row++ {
		tallestTreeInRowSoFar := -1
		for col := 0; col < len(forest); col++ {
			if forest[row][col] > tallestTreeInRowSoFar {
				visibleTrees[location{row, col}] = true
				tallestTreeInRowSoFar = forest[row][col]
			}
			if forest[row][col] == 9 {
				break
			}
		}
	}

	// Visible from right
	for row := 0; row < len(forest); row++ {
		tallestTreeInRowSoFar := -1
		for col := len(forest) - 1; col >= 0; col-- {
			if forest[row][col] > tallestTreeInRowSoFar {
				visibleTrees[location{row, col}] = true
				tallestTreeInRowSoFar = forest[row][col]
			}
			if forest[row][col] == 9 {
				break
			}
		}
	}

	return len(visibleTrees)
}

func Part2() any {
	forest := getInput()
	scenicScores := []int{}

	for col := 0; col < len(forest); col++ {
		for row := 0; row < len(forest); row++ {
			scenicScores = append(scenicScores, getScenicScore(location{row, col}, forest))
		}
	}

	return utils.MaxOf(scenicScores)
}

func getScenicScore(l location, forest [][]int) int {
	treeHeight := forest[l.row][l.col]

	up, down, left, right := 0, 0, 0, 0

	// Look up
	lookUpCoords := location(l)
	for {
		lookUpCoords.row = lookUpCoords.row - 1
		if lookUpCoords.row < 0 {
			break
		}
		if forest[lookUpCoords.row][lookUpCoords.col] < treeHeight {
			up++
			continue
		}
		if forest[lookUpCoords.row][lookUpCoords.col] >= treeHeight {
			up++
			break
		}
	}

	// Look down
	lookDownCoords := location(l)
	for {
		lookDownCoords.row++
		if lookDownCoords.row >= len(forest) {
			break
		}
		if forest[lookDownCoords.row][lookDownCoords.col] < treeHeight {
			down++
			continue
		}
		if forest[lookDownCoords.row][lookDownCoords.col] >= treeHeight {
			down++
			break
		}
	}

	// Look left
	lookLeftCoords := location(l)
	for {
		lookLeftCoords.col--
		if lookLeftCoords.col < 0 {
			break
		}
		if forest[lookLeftCoords.row][lookLeftCoords.col] < treeHeight {
			left++
			continue
		}
		if forest[lookLeftCoords.row][lookLeftCoords.col] >= treeHeight {
			left++
			break
		}
	}

	// Look right
	lookRightCoords := location(l)
	for {
		lookRightCoords.col++
		if lookRightCoords.col >= len(forest) {
			break
		}
		if forest[lookRightCoords.row][lookRightCoords.col] < treeHeight {
			right++
			continue
		}
		if forest[lookRightCoords.row][lookRightCoords.col] >= treeHeight {
			right++
			break
		}
	}

	return up * down * left * right
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 08: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 08: Part 2: = %+v\n", part2Solution)
}

func getInput() [][]int {
	lines, _ := utils.ReadLines(f, "input.txt")
	trees := [][]int{}

	for _, line := range lines {
		parts := strings.Split(line, "")

		row := []int{}
		for _, part := range parts {
			row = append(row, utils.ParseInt(part))
		}
		trees = append(trees, row)
	}

	return trees
}
