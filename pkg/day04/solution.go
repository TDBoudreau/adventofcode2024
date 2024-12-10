package day04

import (
	"strings"

	"github.com/tdboudreau/adventofcode2024/utils"
)

/*
-- Word Search --
Takes in text like:
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX

And solves (forward, backwards, up, down, diagonal, etc) for 'XMAS'.

To solve:
Let us take in the input, then parse, line by line, into a multidimensional array of [][]string

From there we can loop through the puzzle and look in a +3 range around our index for "XMAS".
Since "XMAS" is not a palindrome, if we search around each "X", we should not have any issues
 with counting duplicates.

Directions to check:
	Horizontal:
		- Forward
		- Backward
	Vertical:
		- Up
		- Down
	Diagonal:
		Down:
			- Left
			- Right
		Up:
			- Left
			- Right
*/

func loadPuzzle() ([][]string, error) {
	lines, err := utils.ReadLines("inputs/04/input.txt")
	if err != nil {
		return [][]string{}, err
	}

	var puzzle [][]string
	for _, line := range lines {
		chars := strings.Split(line, "")
		puzzle = append(puzzle, chars)
	}

	return puzzle, nil
}

func Part1() (int, error) {
	puzzle, err := loadPuzzle()
	if err != nil {
		return 0, err
	}

	var xmas int
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[i]); j++ {
			if puzzle[i][j] == "X" {
				// Check around X in each direction.

				// -- Horizontal --
				// Forward
				if j+4 <= len(puzzle[i]) {
					if isXmas(puzzle[i][j : j+4]) {
						// fmt.Printf("i=%v\tj=%v\n", i, j)
						xmas += 1
					}
				}
				// Backward
				if j >= 3 {
					backward := puzzle[i][j-3 : j+1]

					if strings.Join(backward, "") == "SAMX" {
						xmas += 1
					}
				}
				// -- Horizontal --

				// -- Vertical --
				// Up
				if i >= 3 {
					vertSlice := []string{
						puzzle[i][j],
						puzzle[i-1][j],
						puzzle[i-2][j],
						puzzle[i-3][j],
					}

					if isXmas(vertSlice) {
						xmas += 1
					}
				}

				// Down
				if i+4 <= len(puzzle) {

					vertSlice := []string{
						puzzle[i][j],
						puzzle[i+1][j],
						puzzle[i+2][j],
						puzzle[i+3][j],
					}

					if isXmas(vertSlice) {
						xmas += 1
					}
				}
				// -- Vertical --

				// -- Diagonal --
				// Up Left
				if i >= 3 && j >= 3 {
					diagSlice := []string{
						puzzle[i][j],
						puzzle[i-1][j-1],
						puzzle[i-2][j-2],
						puzzle[i-3][j-3],
					}

					if isXmas(diagSlice) {
						xmas += 1
					}
				}

				// Up Right
				if i >= 3 && j+3 < len(puzzle[i]) {
					diagSlice := []string{
						puzzle[i][j],
						puzzle[i-1][j+1],
						puzzle[i-2][j+2],
						puzzle[i-3][j+3],
					}

					if isXmas(diagSlice) {
						xmas += 1
					}
				}

				// Down Left
				if i+3 < len(puzzle) && j >= 3 {
					diagSlice := []string{
						puzzle[i][j],
						puzzle[i+1][j-1],
						puzzle[i+2][j-2],
						puzzle[i+3][j-3],
					}

					if isXmas(diagSlice) {
						xmas += 1
					}
				}

				// Down Right
				if i+3 < len(puzzle) && j+3 < len(puzzle[i]) {
					diagSlice := []string{
						puzzle[i][j],
						puzzle[i+1][j+1],
						puzzle[i+2][j+2],
						puzzle[i+3][j+3],
					}

					if isXmas(diagSlice) {
						xmas += 1
					}
				}
				// -- Diagonal --
			}
		}
	}

	return xmas, nil

	// 2665 too high
	// 2601 too high
	// ** 2500 ** winner winner
}

func isXmas(input []string) bool {
	return strings.Join(input, "") == "XMAS"
}

/*
Part 2 X-MAS, not XMAS 🤦‍♂️

We now need to find instances where two "MAS" appear in an X (diagonally).
For example:
	M.S
	.A.
	M.S

Let's use the same basis we did before, but, instead of looking for "X"
 we look for the common denominator "A". From there we can check if "A" is
 in the middle of an X-MAS.
*/

func Part2() (int, error) {
	puzzle, err := loadPuzzle()
	if err != nil {
		return 0, err
	}

	var x_mas int

	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[i]); j++ {
			// fmt.Printf("i=%v\tj=%v\n", i, j)
			if puzzle[i][j] == "A" {
				if i-1 >= 0 && j-1 >= 0 && i+2 <= len(puzzle) && j+2 <= len(puzzle[i]) {
					top := puzzle[i-1][j-1 : j+2]
					// mid := puzzle[i][j-1 : j+2]
					bot := puzzle[i+1][j-1 : j+2]

					// fmt.Printf("%v\n%v\n%v\n", top, mid, bot)

					if isXMAS(top, bot) {
						x_mas++
					}
				}
			}
			// fmt.Println()
		}
	}

	return x_mas, nil
	// ** 1933 ** winner winner
}

type Directions struct {
	dirX int
	dirY int
}

var directions = []Directions{
	{-1, 1}, // Top-Left <-> Bottom-Right
	{1, -1}, // Top-Right <-> Bottom-Left
}

func isXMAS(top, bot []string) bool {
	var mas_count int

	for _, dir := range directions {
		if checkMAS(top[1+dir.dirX], bot[1+dir.dirY]) {
			mas_count++
		}
	}

	return mas_count == 2
}

// Checks both forwards and backwards for a given diagonal.
func checkMAS(a, b string) bool {
	return (a+"A"+b == "MAS") || (b+"A"+a == "MAS")
}
