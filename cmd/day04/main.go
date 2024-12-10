package main

import (
	"fmt"

	"github.com/tdboudreau/adventofcode2024/pkg/day04"
)

func main() {
	solution, err := day04.Part1()
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", solution)

	solution, err = day04.Part2()
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 2:", solution)
}
