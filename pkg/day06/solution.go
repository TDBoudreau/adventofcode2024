package day06

import (
	"fmt"
	"strings"
	"sync"

	"github.com/tdboudreau/adventofcode2024/utils"
)

// Solve reads the input file and returns the results for both parts of the problem.
func Solve(inputPath string) (int, int, error) {
	lines, err := utils.ReadLines("input/06/input.txt")
	if err != nil {
		return 0, 0, err
	}

	var startPos Position

	var guardMap [][]string
	for i, line := range lines {
		lineParts := strings.Split(line, "")
		for j, char := range lineParts {
			if char == "^" {
				startPos = Position{y: i, x: j, direction: "^"}
			} else if char == ">" {
				startPos = Position{y: i, x: j, direction: ">"}
			} else if char == "v" {
				startPos = Position{y: i, x: j, direction: "V"}
			} else if char == "<" {
				startPos = Position{y: i, x: j, direction: "<"}
			}
		}
		guardMap = append(guardMap, lineParts)
	}
	// fmt.Println("Start Position: ", startPos)

	answer1, err := part1(startPos, guardMap)
	if err != nil {
		return 0, 0, err
	}

	answer2, err := part2(guardMap, startPos)
	if err != nil {
		return 0, 0, err
	}

	return answer1, answer2, nil
}

type Position struct {
	x         int
	y         int
	direction string
}

type ObstaclePos struct {
	x int
	y int
}

func (p *Position) rotate() {
	switch p.direction {
	case "^":
		p.direction = ">"
	case ">":
		p.direction = "v"
	case "v":
		p.direction = "<"
	case "<":
		p.direction = "^"
	}
}

func part1(startPos Position, guardMap [][]string) (int, error) {
	finalMap, _ := runGame(guardMap, startPos)

	var countX int
	for _, line := range finalMap {
		for _, c := range line {
			if c == "X" {
				countX++
			}
		}
	}

	return countX, nil
}

func runGame(guardMap [][]string, pos Position) ([][]string, Position) {
	running := true

	for running {
		// printMap(pos, guardMap)

		switch pos.direction {
		case "^":
			if pos.y > 0 {
				if guardMap[pos.y-1][pos.x] != "#" {
					guardMap[pos.y-1][pos.x] = "^"
					guardMap[pos.y][pos.x] = "X"
					pos.y = pos.y - 1
				} else {
					pos.rotate()
				}
			} else {
				guardMap[pos.y][pos.x] = "X"
				running = false
			}
		case ">":
			if pos.x < len(guardMap[0])-1 {
				if guardMap[pos.y][pos.x+1] != "#" {
					guardMap[pos.y][pos.x+1] = ">"
					guardMap[pos.y][pos.x] = "X"
					pos.x = pos.x + 1
				} else {
					pos.rotate()
				}
			} else {
				guardMap[pos.y][pos.x] = "X"
				running = false
			}
		case "v":
			if pos.y < len(guardMap)-1 {
				if guardMap[pos.y+1][pos.x] != "#" {
					guardMap[pos.y+1][pos.x] = "v"
					guardMap[pos.y][pos.x] = "X"
					pos.y = pos.y + 1
				} else {
					pos.rotate()
				}
			} else {
				guardMap[pos.y][pos.x] = "X"
				running = false
			}
		case "<":
			if pos.x > 0 {
				if guardMap[pos.y][pos.x-1] != "#" {
					guardMap[pos.y][pos.x-1] = "<"
					guardMap[pos.y][pos.x] = "X"
					pos.x = pos.x - 1
				} else {
					pos.rotate()
				}
			} else {
				guardMap[pos.y][pos.x] = "X"
				running = false
			}
		}
	}

	return guardMap, pos
}

/*
-- Part 2 --
Add a single obstacle to create infinite loops.

How many options are there to create invite loops?

How to do this?...
*/
func part2(originalMap [][]string, startPos Position) (int, error) {
	// Identify candidate cells
	candidates := []ObstaclePos{}
	for y := range originalMap {
		for x := range originalMap[y] {
			// Consider placing an obstacle only on walkable spaces
			if originalMap[y][x] == "." || originalMap[y][x] == "X" {
				candidates = append(candidates, ObstaclePos{x: x, y: y})
			}
		}
	}

	// Use a goroutine for each candidate to check if it causes a loop.
	// If there are a very large number of candidates, consider a worker pool.
	resultCh := make(chan int, len(candidates))
	var wg sync.WaitGroup
	wg.Add(len(candidates))

	for _, c := range candidates {
		go func(c ObstaclePos) {
			defer wg.Done()
			guardMapCopy := copyMap(originalMap)
			guardMapCopy[c.y][c.x] = "#"
			if doesCauseLoop(guardMapCopy, startPos) {
				resultCh <- 1
			} else {
				resultCh <- 0
			}
		}(c)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	close(resultCh)

	loopCounter := 0
	for val := range resultCh {
		loopCounter += val
	}

	return loopCounter, nil
}

func doesCauseLoop(guardMap [][]string, startPos Position) bool {
	visitedStates := make(map[string]bool)
	pos := startPos

	for {
		stateKey := fmt.Sprintf("%d,%d,%s", pos.x, pos.y, pos.direction)
		if visitedStates[stateKey] {
			// loop detected
			return true
		}
		visitedStates[stateKey] = true

		nextPos, ok := moveGuard(guardMap, pos)
		if !ok {
			// no loop if guard stops
			return false
		}

		pos = nextPos
	}
}

func moveGuard(guardMap [][]string, pos Position) (Position, bool) {
	maxY := len(guardMap)
	maxX := len(guardMap[0])

	switch pos.direction {
	case "^":
		if pos.y > 0 {
			if guardMap[pos.y-1][pos.x] != "#" {
				pos.y = pos.y - 1
			} else {
				pos.rotate()
			}
		} else {
			return pos, false
		}

	case ">":
		if pos.x < maxX-1 {
			if guardMap[pos.y][pos.x+1] != "#" {
				pos.x = pos.x + 1
			} else {
				pos.rotate()
			}
		} else {
			return pos, false
		}

	case "v":
		if pos.y < maxY-1 {
			if guardMap[pos.y+1][pos.x] != "#" {
				pos.y = pos.y + 1
			} else {
				pos.rotate()
			}
		} else {
			return pos, false
		}

	case "<":
		if pos.x > 0 {
			if guardMap[pos.y][pos.x-1] != "#" {
				pos.x = pos.x - 1
			} else {
				pos.rotate()
			}
		} else {
			return pos, false
		}
	}

	return pos, true
}

func copyMap(m [][]string) [][]string {
	newMap := make([][]string, len(m))
	for i := range m {
		newMap[i] = make([]string, len(m[i]))
		copy(newMap[i], m[i])
	}
	return newMap
}
