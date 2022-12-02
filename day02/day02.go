package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
)

const (
	lose     = 0
	draw     = 3
	win      = 6
	rock     = 1
	paper    = 2
	scissors = 3
)

type Shape = string

type Round struct {
	Them Shape
	Me   Shape
}

type Scoring struct {
	Shape   int
	Outcome int
}

func (s Scoring) TotalScore() int {
	return s.Outcome + s.Shape
}

type RuleSet map[string]Scoring

func (r Round) Result(rules RuleSet) Scoring {
	return rules[r.Them+":"+r.Me]
}

//go:embed input.txt
var f embed.FS

func Part1() any {
	rounds := getInput()

	strategy := RuleSet{
		"A:X": Scoring{Outcome: draw, Shape: rock},
		"B:X": Scoring{Outcome: lose, Shape: rock},
		"C:X": Scoring{Outcome: win, Shape: rock},
		"A:Y": Scoring{Outcome: win, Shape: paper},
		"B:Y": Scoring{Outcome: draw, Shape: paper},
		"C:Y": Scoring{Outcome: lose, Shape: paper},
		"A:Z": Scoring{Outcome: lose, Shape: scissors},
		"B:Z": Scoring{Outcome: win, Shape: scissors},
		"C:Z": Scoring{Outcome: draw, Shape: scissors},
	}

	totalScore := 0
	for _, round := range rounds {
		totalScore += round.Result(strategy).TotalScore()
	}

	return totalScore
}

func Part2() any {
	rounds := getInput()

	totalScore := 0
	for _, round := range rounds {
		switch round.Me {
		case "X": // We need to lose
			switch round.Them {
			case "A":
				totalScore += round.Result(RuleSet{"A:X": Scoring{Outcome: lose, Shape: scissors}}).TotalScore()
			case "B":
				totalScore += round.Result(RuleSet{"B:X": Scoring{Outcome: lose, Shape: rock}}).TotalScore()
			case "C":
				totalScore += round.Result(RuleSet{"C:X": Scoring{Outcome: lose, Shape: paper}}).TotalScore()
			}
		case "Y": // We need to draw
			switch round.Them {
			case "A":
				totalScore += round.Result(RuleSet{"A:Y": Scoring{Outcome: draw, Shape: rock}}).TotalScore()
			case "B":
				totalScore += round.Result(RuleSet{"B:Y": Scoring{Outcome: draw, Shape: paper}}).TotalScore()
			case "C":
				totalScore += round.Result(RuleSet{"C:Y": Scoring{Outcome: draw, Shape: scissors}}).TotalScore()
			}
		case "Z": // We need to win
			switch round.Them {
			case "A":
				totalScore += round.Result(RuleSet{"A:Z": Scoring{Outcome: win, Shape: paper}}).TotalScore()
			case "B":
				totalScore += round.Result(RuleSet{"B:Z": Scoring{Outcome: win, Shape: scissors}}).TotalScore()
			case "C":
				totalScore += round.Result(RuleSet{"C:Z": Scoring{Outcome: win, Shape: rock}}).TotalScore()
			}
		}
	}
	return totalScore
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 02: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 02: Part 2: = %+v\n", part2Solution)
}

func getInput() []Round {
	lines, _ := utils.ReadLines(f, "input.txt")
	rounds := []Round{}
	for _, line := range lines {
		rounds = append(rounds, Round{line[0:1], line[2:3]})
	}

	return rounds
}
