package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var f embed.FS

type Node struct {
	Type string
	Name string

	// Used only for "file" types
	Size int

	// Used onlye for "dir" types
	Children map[string]*Node

	Parent *Node
}

func Part1() any {
	commands := getInput()
	root := getTree(commands)

	_, sizes := sizer(root)
	limitedSizesSum := 0
	for _, size := range sizes {
		if size <= 100_000 {
			limitedSizesSum += size
		}
	}

	return limitedSizesSum
}

func Part2() any {
	commands := getInput()
	root := getTree(commands)

	totalSize, sizes := sizer(root)

	neededSpace := 30_000_000
	freeSpace := 70_000_000 - totalSize
	minSizeNeeded := neededSpace - freeSpace

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] < sizes[j]
	})

	for _, size := range sizes {
		if size >= minSizeNeeded {
			return size
		}
	}

	return nil
}

func getTree(commands []string) *Node {
	root := &Node{
		Type:     "dir",
		Name:     "/",
		Children: map[string]*Node{},
	}

	currentDir := root
	for _, command := range commands {
		switch {
		case strings.HasPrefix(command, "$"):
			command = strings.TrimPrefix(command, "$ ")
			switch {
			case strings.HasPrefix(command, "cd"):
				targetDir := strings.TrimPrefix(command, "cd ")
				switch targetDir {
				case "/":
					currentDir = root
				case "..":
					currentDir = currentDir.Parent
				default:
					currentDir = currentDir.Children[targetDir]
				}
			case command == "ls":
				// We can essentially ignore this command
			}
		default:
			parts := strings.Split(command, " ")
			switch parts[0] {
			case "dir":
				currentDir.Children[parts[1]] = &Node{
					Type:     "dir",
					Name:     parts[1],
					Children: map[string]*Node{},
					Parent:   currentDir,
				}
			default:
				currentDir.Children[parts[1]] = &Node{
					Type:   "file",
					Name:   parts[1],
					Size:   utils.ParseInt(parts[0]),
					Parent: currentDir,
				}
			}
		}
	}

	return root
}

func sizer(n *Node) (int, []int) {
	dirSizes := []int{}
	var sizer func(node *Node) int
	sizer = func(n *Node) int {
		switch n.Type {
		case "file":
			return n.Size
		case "dir":
			size := 0
			for _, child := range n.Children {
				size += sizer(child)
			}
			dirSizes = append(dirSizes, size)
			return size
		default:
			panic("Unknown type")
		}
	}
	totalSize := sizer(n)
	return totalSize, dirSizes
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 07: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 07: Part 2: = %+v\n", part2Solution)
}

func getInput() []string {
	lines, _ := utils.ReadLines(f, "input.txt")
	return lines
}
