define GO_MOD_TEMPLATE
module adventofcode2022/day${day}

go 1.19
endef

define GO_FILE_TEMPLATE
package main

import (
	"embed"
	"fmt"
)

//go:embed input.txt
var f embed.FS

func Part1() any {
	return nil
}

func Part2() any {
	return nil
}

func main() {
	part1Solution := Part1()
	part2Solution := Part2()

  fmt.Printf("Day ${day}: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day ${day}: Part 2: = %+v\n", part2Solution)
}
endef

export GO_MOD_TEMPLATE
export GO_FILE_TEMPLATE

init:
	@mkdir day${day}
	@echo "$$GO_MOD_TEMPLATE" > day${day}/go.mod
	@echo "$$GO_FILE_TEMPLATE" > day${day}/day${day}.go
	@touch day${day}/input.txt
	@touch day${day}/README.md

new:
	@go run ./runner.go --command new

run-current:
	@go run ./runner.go --command runCurrent

run-all:
	@go run ./runner.go --command runAll

build-current:
	@go run ./runner.go --command buildCurrent

build-all:
	@go run ./runner.go --command buildAll

