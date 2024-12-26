package day10

import (
	"strconv"
	"strings"

	"github.com/tdboudreau/adventofcode2024/utils"
)

type Position struct {
	x int
	y int
}

func createMap(lines []string) ([][]int, error) {
	var hikingMap [][]int

	for i, line := range lines {
		trailPoints := strings.Split(line, "")
		var trailPointsInts []int
		for _, char := range trailPoints {
			integer, err := strconv.Atoi(char)
			if err != nil {
				return [][]int{}, err
			}
			trailPointsInts = append(trailPointsInts, integer)
		}
		hikingMap[i] = append(hikingMap[i], trailPointsInts...)
	}

	return hikingMap, nil
}

// Solve reads the input file and returns the results for both parts of the problem.
func Solve(inputPath string) (int, int, error) {
	// lines, err := utils.ReadLines("input/10/input.txt")
	lines, err := utils.ReadLines("input/10/ex1.txt")
	if err != nil {
		return 0, 0, err
	}

	hikingMap, err := createMap(lines)
	if err != nil {
		return 0, 0, err
	}

	// Implement your logic here.
	// For now, return placeholder values.
	ans1, err := part1(hikingMap)
	if err != nil {
		return 0, 0, err
	}
	ans2, err := part2()
	if err != nil {
		return 0, 0, err
	}

	return ans1, ans2, nil
}

/*
-- Part 1 --
The reindeer is holding a book titled "Lava Island Hiking Guide". However,
when you open the book, you discover that most of it seems to have been
scorched by lava! As you're about to ask how you can help, the reindeer
brings you a blank topographic map of the surrounding area (your puzzle input)
and looks up at you excitedly.

Perhaps you can help fill in the missing hiking trails?

The topographic map indicates the height at each position using a scale
from 0 (lowest) to 9 (highest). For example:
0123
1234
8765
9876

Based on un-scorched scraps of the book, you determine that a good hiking trail
is as long as possible and has an even, gradual, uphill slope. For all practical
purposes, this means that a hiking trail is any path that starts at height 0,
ends at height 9, and always increases by a height of exactly 1 at each step.
Hiking trails never include diagonal steps - only up, down, left, or right
(from the perspective of the map).

You look up from the map and notice that the reindeer has helpfully begun to
construct a small pile of pencils, markers, rulers, compasses, stickers, and
other equipment you might need to update the map with hiking trails.

A trailhead is any position that starts one or more hiking trails -
here, these positions will always have height 0. Assembling more fragments of
pages, you establish that a trailhead's score is the number of 9-height
positions reachable from that trailhead via a hiking trail. In the above
example, the single trailhead in the top left corner has a score of 1 because
it can reach a single 9 (the one in the bottom left).
*/
func part1(hikingMap [][]int) (int, error) {
	// A trailhead is any position that starts one or more hiking trails.
	// A trailhead's score is the number of 9-height positions reachable from that

	// 1. Loop through each line and look for a 0 (starting-position).
	// 2. Look at the surrounding tiles (up/down/left/right), and traverse valid paths.
	// -- Note: if two different trails end up at the same '9', then they only get counted once to the trail's score.
	// 3. Count how many (unique) 9's are reachable via that trailhead.
	for i := 0; i < len(hikingMap); i++ {
		for j := 0; j < len(hikingMap); j++ {
			if hikingMap[i][j] != 0 {
				continue
			}

			// check surrounding tiles and find valid trails.
		}
	}

	return 0, nil
}

func traverseTrail(hikingMap [][]int, pos Position, currentHeight int) (Position, bool) {
	nextHeight := currentHeight + 1

	// Up
	if pos.y > 0 {
		if hikingMap[pos.y-1][pos.x] == nextHeight {
			traverseTrail(hikingMap, Position{y: pos.y - 1, x: pos.x}, nextHeight)
		}
	}

	// Down
	if pos.y < len(hikingMap)-1 {
		if hikingMap[pos.y+1][pos.x] == nextHeight {
			traverseTrail(hikingMap, Position{y: pos.y + 1, x: pos.x}, nextHeight)
		}

	}

	// Left
	if pos.x > 0 {
		if hikingMap[pos.y][pos.x-1] == nextHeight {
			traverseTrail(hikingMap, Position{y: pos.y, x: pos.x - 1}, nextHeight)
		}
	}

	// Right
	if pos.x < len(hikingMap[pos.y])-1 {
		if hikingMap[pos.y][pos.x+1] == nextHeight {
			traverseTrail(hikingMap, Position{y: pos.y, x: pos.x + 1}, nextHeight)
		}
	}

	return Position{}, false
}

/*
-- Part 2 --
*/
func part2() (int, error) {
	// Logic for part2

	return 0, nil
}
