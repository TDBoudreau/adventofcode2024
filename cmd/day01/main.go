package main

import (
	"fmt"

	"github.com/tdboudreau/adventofcode2024/pkg/day01"
)

func main() {
	solution, err := day01.Part1()
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", solution)

	solution, err = day01.Part2()
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 2:", solution)
}
