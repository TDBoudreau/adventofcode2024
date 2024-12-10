package main

import (
	"fmt"

	"github.com/tdboudreau/adventofcode2024/pkg/day02"
)

func main() {
	solution, err := day02.Part1()
	if err != nil {
		panic(err)
	}

	fmt.Println("Number of safe reports:", solution)

	fmt.Println("--------------------------")

	solution, err = day02.Part2()
	if err != nil {
		panic(err)
	}

	fmt.Println("Number of safe dampened reports:", solution)
}
