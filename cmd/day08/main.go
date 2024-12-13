package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tdboudreau/adventofcode2024/pkg/day08"
)

func main() {
	// Example of how you might run your solution
	exePath, _ := os.Executable()
	dir := filepath.Dir(exePath)
	inputPath := filepath.Join(dir, "../../input/08", "input.txt")

	resultPart1, resultPart2, err := day08.Solve(inputPath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Day 08, Part 1: %v\n", resultPart1)
	fmt.Printf("Day 08, Part 2: %v\n", resultPart2)
}
