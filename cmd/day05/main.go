package main

import (
	"fmt"

	"github.com/tdboudreau/adventofcode2024/pkg/day05"
)

func main() {
	part1, part2, err := day05.AllParts()
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
