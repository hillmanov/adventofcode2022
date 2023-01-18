package main

import (
	"adventofcode/utils"
	"embed"
	"fmt"
)

//go:embed input.txt
var f embed.FS

type ObsidianRobotCost struct {
	OreCost  int
	ClayCost int
}

type GeodeRobotCost struct {
	OreCost      int
	ObsidianCost int
}

type Blueprint struct {
	ID                     int
	OreRobotCost           int
	ClayRobotCost          int
	ObsidianRobotOreCost   int
	ObsidianRobotClayCost  int
	GeodeRobotOreCost      int
	GeodeRobotObsidianCost int
}

type RobotCounts struct {
	oreRobots      int
	clayRobots     int
	obsidianRobots int
	geodeRobots    int
}

func (r RobotCounts) Copy() RobotCounts {
	return RobotCounts(r)
}

type ResourceCounts struct {
	oreCount      int
	clayCount     int
	obsidianCount int
	geodeCount    int
}

func (r ResourceCounts) Copy() ResourceCounts {
	return ResourceCounts(r)
}

var max = 0

type State struct {
	minute    int
	resources ResourceCounts
	robots    RobotCounts
	blueprint Blueprint
}

type history string

var maxes = map[history]int{}

func (b Blueprint) GeodeProductionAmount(minute int, resources ResourceCounts, robots RobotCounts, results *[]int, history string) int {

	for i := minute; i <= 24; i++ {
		resources.oreCount += robots.oreRobots
		resources.clayCount += robots.clayRobots
		resources.obsidianCount += robots.obsidianRobots
		resources.geodeCount += robots.geodeRobots

		// Keep track of what things we CAN make
		// Go through the list of each thing we can make and recursively call ourselves with the current state, minus the resources needed to make the thing

		if resources.oreCount >= b.GeodeRobotOreCost && resources.obsidianCount >= b.GeodeRobotObsidianCost {
			b.GeodeProductionAmount(i+1,
				ResourceCounts{
					oreCount:      resources.oreCount - b.GeodeRobotOreCost,
					clayCount:     resources.clayCount,
					obsidianCount: resources.obsidianCount - b.GeodeRobotObsidianCost,
					geodeCount:    resources.geodeCount,
				},
				RobotCounts{
					oreRobots:      robots.oreRobots,
					clayRobots:     robots.clayRobots,
					obsidianRobots: robots.obsidianRobots,
					geodeRobots:    robots.geodeRobots + 1,
				}, results, history+fmt.Sprintf("%d:%d", i, "geode"))
		}

		if resources.oreCount >= b.ObsidianRobotOreCost && resources.clayCount >= b.ObsidianRobotClayCost {
			b.GeodeProductionAmount(i+1,
				ResourceCounts{
					oreCount:      resources.oreCount - b.ObsidianRobotOreCost,
					clayCount:     resources.clayCount - b.ObsidianRobotClayCost,
					obsidianCount: resources.obsidianCount,
					geodeCount:    resources.geodeCount,
				},
				RobotCounts{
					oreRobots:      robots.oreRobots,
					clayRobots:     robots.clayRobots,
					obsidianRobots: robots.obsidianRobots + 1,
					geodeRobots:    robots.geodeRobots,
				}, results, history+fmt.Sprintf("%d:%d", i, "obsidian"))
		}

		if resources.oreCount >= b.ClayRobotCost {
			b.GeodeProductionAmount(i+1,
				ResourceCounts{
					oreCount:      resources.oreCount - b.ClayRobotCost,
					clayCount:     resources.clayCount,
					obsidianCount: resources.obsidianCount,
					geodeCount:    resources.geodeCount,
				},
				RobotCounts{
					oreRobots:      robots.oreRobots,
					clayRobots:     robots.clayRobots + 1,
					obsidianRobots: robots.obsidianRobots,
					geodeRobots:    robots.geodeRobots,
				}, results, history+fmt.Sprintf("%d:%d", i, "clay"))
		}
		if resources.oreCount >= b.OreRobotCost {
			b.GeodeProductionAmount(i+1,
				ResourceCounts{
					oreCount:      resources.oreCount - b.OreRobotCost,
					clayCount:     resources.clayCount,
					obsidianCount: resources.obsidianCount,
					geodeCount:    resources.geodeCount,
				},
				RobotCounts{
					oreRobots:      robots.oreRobots + 1,
					clayRobots:     robots.clayRobots,
					obsidianRobots: robots.obsidianRobots,
					geodeRobots:    robots.geodeRobots,
				}, results, history+fmt.Sprintf("%d:%d", i, "ore"))
		}
	}

	return resources.geodeCount
}

func Part1() any {
	blueprints := getInput()

	results := []int{}

	resources := ResourceCounts{
		oreCount:      0,
		clayCount:     0,
		obsidianCount: 0,
		geodeCount:    0,
	}

	robots := RobotCounts{
		oreRobots:      1,
		clayRobots:     0,
		obsidianRobots: 0,
		geodeRobots:    0,
	}

	blueprints[0].GeodeProductionAmount(1, resources, robots, &results, []Decision{})

	// for _, blueprint := range blueprints {
	// 	geodeAmount := getMaxConfiguration(blueprint)
	// 	fmt.Printf("geodeAmount = %+v\n", geodeAmount)
	// }

	return nil
}

func Part2() any {
	return nil
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

	fmt.Printf("Day 19: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 19: Part 2: = %+v\n", part2Solution)
}

func getInput() []Blueprint {
	lines, _ := utils.ReadLines(f, "input.txt")
	blueprints := []Blueprint{}
	for _, line := range lines {
		blueprint := Blueprint{}
		fmt.Sscanf(
			line,
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&blueprint.ID,
			&blueprint.OreRobotCost,
			&blueprint.ClayRobotCost,
			&blueprint.ObsidianRobotOreCost,
			&blueprint.ObsidianRobotClayCost,
			&blueprint.GeodeRobotOreCost,
			&blueprint.GeodeRobotObsidianCost,
		)
		blueprints = append(blueprints, blueprint)
	}

	return blueprints

}
