package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tdboudreau/adventofcode2024/pkg/day06"
)

func main() {
	// Example of how you might run your solution
	exePath, _ := os.Executable()
	dir := filepath.Dir(exePath)
	inputPath := filepath.Join(dir, "../../input/06", "input.txt")

	resultPart1, resultPart2, err := day06.Solve(inputPath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Day 06, Part 1: %v\n", resultPart1)
	fmt.Printf("Day 06, Part 2: %v\n", resultPart2)
}
