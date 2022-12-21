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

func (b Blueprint) GeodeProductionAmount(minutes, oreRobotLimit, clayRobotLimit, obsidianRobotLimit int) int {

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

	for i := 1; i <= minutes; i++ {
		makingOreRobot := false
		makingClayRobot := false
		makingObsidianRobot := false
		makingGeodeRobot := false

		fmt.Printf("== Minute: %+v ==\n", i)
		if resources.oreCount >= b.OreRobotCost && robots.oreRobots < oreRobotLimit {
			fmt.Printf("Spend %d ore to start building a ore-collecting robot.\n", b.OreRobotCost)
			resources.oreCount -= b.OreRobotCost
			makingOreRobot = true
		}

		if resources.oreCount >= b.ClayRobotCost && robots.clayRobots < clayRobotLimit {
			fmt.Printf("Spend %d ore to start building a clay-collecting robot.\n", b.ClayRobotCost)
			resources.oreCount -= b.ClayRobotCost
			makingClayRobot = true
		}

		if resources.oreCount >= b.ObsidianRobotOreCost && resources.clayCount >= b.ObsidianRobotClayCost && robots.obsidianRobots < obsidianRobotLimit {
			fmt.Printf("Spend %d ore and %d clay to start building a obsidian-collecting robot.\n", b.ObsidianRobotOreCost, b.ObsidianRobotClayCost)
			resources.oreCount -= b.ObsidianRobotOreCost
			resources.clayCount -= b.ObsidianRobotClayCost
			makingObsidianRobot = true
		}

		if resources.oreCount >= b.GeodeRobotOreCost && resources.obsidianCount >= b.GeodeRobotObsidianCost {
			fmt.Printf("Spend %d ore and %d obsidian to start building a geode-collecting robot.\n", b.GeodeRobotOreCost, b.GeodeRobotObsidianCost)
			resources.oreCount -= b.GeodeRobotOreCost
			resources.obsidianCount -= b.GeodeRobotObsidianCost
			makingGeodeRobot = true
		}

		resources.oreCount += robots.oreRobots
		fmt.Printf("%d ore collecting robots collects %d ore; You now have %d ore.\n", robots.oreRobots, robots.oreRobots, resources.oreCount)

		resources.clayCount += robots.clayRobots
		if robots.clayRobots > 0 {
			fmt.Printf("%d clay collecting robots collects %d clay; You now have %d clay.\n", robots.clayRobots, robots.clayRobots, resources.clayCount)
		}

		if robots.obsidianRobots > 0 {
			fmt.Printf("%d obsidian collecting robots collects %d obsidian; You now have %d obsidian.\n", robots.obsidianRobots, robots.obsidianRobots, resources.obsidianCount)
		}
		resources.obsidianCount += robots.obsidianRobots

		if robots.geodeRobots > 0 {
			fmt.Printf("%d geode producing robots produces %d geode; You now have %d geode.\n", robots.geodeRobots, robots.geodeRobots, resources.geodeCount)
		}
		resources.geodeCount += robots.geodeRobots

		if makingOreRobot {
			robots.oreRobots++
			fmt.Printf("The new ore collecting robot is ready. You now have %d of them.\n", robots.oreRobots)
		}
		if makingClayRobot {
			robots.clayRobots++
			fmt.Printf("The new clay collecting robot is ready. You now have %d of them.\n", robots.clayRobots)
		}
		if makingObsidianRobot {
			robots.obsidianRobots++
			fmt.Printf("The new obsidian collecting robot is ready. You now have %d of them.\n", robots.obsidianRobots)
		}
		if makingGeodeRobot {
			robots.geodeRobots++
			fmt.Printf("The new geode collecting robot is ready. You now have %d of them.\n", robots.geodeRobots)
		}

		fmt.Println()
	}

	return resources.geodeCount
}

func Part1() any {
	blueprints := getInput()

	blueprints[0].GeodeProductionAmount(24, 1, 3, 2)

	// for _, blueprint := range blueprints {
	// 	geodeAmount := getMaxConfiguration(blueprint)
	// 	fmt.Printf("geodeAmount = %+v\n", geodeAmount)
	// }

	return nil
}

func getMaxConfiguration(b Blueprint) int {
	max := 0
	for oreRobotLimit := 1; oreRobotLimit <= 24; oreRobotLimit++ {
		for clayRobotLimit := 1; clayRobotLimit <= 24; clayRobotLimit++ {
			for obsidianRobotLimit := 1; obsidianRobotLimit <= 24; obsidianRobotLimit++ {
				geodeAmount := b.GeodeProductionAmount(24, oreRobotLimit, clayRobotLimit, obsidianRobotLimit)
				if geodeAmount > max {
					max = geodeAmount
					fmt.Printf("new max: %d, oreRobotLimit: %d, clayRobotLimit: %d, obsidianRobotLimit: %d\n", max, oreRobotLimit, clayRobotLimit, obsidianRobotLimit)
				}
			}
		}
	}
	return max
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
