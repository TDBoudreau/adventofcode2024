package main

import (
	"fmt"

	"github.com/tdboudreau/adventofcode2024/pkg/day03"
)

func main() {
	solution, err := day03.Part1()
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", solution)
}
