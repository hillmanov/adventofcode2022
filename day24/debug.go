package main

import "fmt"

func (v Valley) BlizzardsAtPosition(p Position, minute int) []Blizzard {
	blizzards := []Blizzard{}
	for _, b := range v.Blizzards[minute] {
		if b.Row == p.Row && b.Col == p.Col {
			blizzards = append(blizzards, b)
		}
	}
	return blizzards
}

func (v Valley) Dump(expedition Position, minute int) {
	for rowIndex := v.MinRow - 1; rowIndex < v.MaxRow+1; rowIndex++ {
		for colIndex := v.MinCol - 1; colIndex < v.MaxCol+1; colIndex++ {
			if (Position{Row: rowIndex, Col: colIndex}) == expedition {
				fmt.Print("E")
			} else if v.Walls[Position{Row: rowIndex, Col: colIndex}] {
				fmt.Print("#")
			} else if rowIndex == v.Start.Row && colIndex == v.Start.Col {
				fmt.Print("s")
			} else if rowIndex == v.End.Row && colIndex == v.End.Col {
				fmt.Print("e")
			} else {
				blizzardsHere := v.BlizzardsAtPosition(Position{Row: rowIndex, Col: colIndex}, minute)
				if len(blizzardsHere) > 1 {
					fmt.Print(len(blizzardsHere))
				} else if len(blizzardsHere) == 1 {
					fmt.Print(blizzardsHere[0].Direction)
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
