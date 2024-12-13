package day03

import (
	"strconv"

	"github.com/tdboudreau/adventofcode2024/utils"
)

func Part1() (int, error) {
	input, err := utils.ReadFile("input/03/input.txt")
	if err != nil {
		return 0, err
	}

	// Toggle for whether mul() calls count
	do := true
	var total int

	for i := 0; i < len(input); i++ {
		// Check for "don't()" and flip the toggle off
		if i+7 <= len(input) && input[i:i+7] == "don't()" {
			do = false
			i += 6
			continue
		}

		// Check for "do()" and flip the toggle on
		if i+4 <= len(input) && input[i:i+4] == "do()" {
			do = true
			i += 3
			continue
		}

		// If we're allowed to count mul(), go for it
		if do {
			if i+4 <= len(input) && input[i:i+4] == "mul(" {
				res, length := parseMulInstruction(input[i+4:])
				if length > 0 {
					total += res
					i += 3 + length
				}
			}
		}
	}

	return total, nil
}

// Attempts to parse "X,Y)" right after "mul("
func parseMulInstruction(s string) (int, int) {
	closeIndex := -1
	for j := 0; j < len(s); j++ {
		if s[j] == ')' {
			closeIndex = j
			break
		}
	}

	if closeIndex == -1 {
		return 0, 0
	}

	// inside is a slice containing X,Y up to, but not including, the closing ')'
	inside := s[:closeIndex]
	commaIndex := -1
	for idx, ch := range inside {
		if ch == ',' {
			if commaIndex == -1 {
				commaIndex = idx
			} else {
				// Multiple commas ',' - bad formatting, skip
				return 0, 0
			}
		}
	}

	// no comma found
	if commaIndex == -1 {
		return 0, 0
	}

	Xstr := inside[:commaIndex]
	Ystr := inside[commaIndex+1:]

	X, errX := strconv.Atoi(Xstr)
	Y, errY := strconv.Atoi(Ystr)
	if errX != nil || errY != nil {
		return 0, 0
	}

	length := closeIndex + 1
	return X * Y, length
}

// // Regex method:
// // Part1Regex parses a text file containing multiplication and control instructions,
// // calculating a total based on enabled multiplication operations
// func Part1Regex() (int, error) {
// 	// Read the input file into a string
// 	input, err := utils.ReadFile("inputs/03/input.txt")
// 	if err != nil {
// 		return 0, err
// 	}

// 	// Regex pattern matches three different instruction types:
// 	// 1. mul(X,Y)   - where X and Y are 1-3 digit numbers
// 	// 2. do()       - enables multiplication operations
// 	// 3. don't()    - disables multiplication operations
// 	//
// 	// Capture groups:
// 	// [1] Full mul(...) match
// 	// [2] First number in mul
// 	// [3] Second number in mul
// 	// [4] do() match
// 	// [5] don't() match
// 	re := regexp.MustCompile(`(mul\(([0-9]{1,3}),([0-9]{1,3})\))|(do\(\))|(don't\(\))`)

// 	// FindAllStringSubmatch returns all matches and their capture groups:
// 	// - Outer array contains all matches found in the input
// 	// - Inner array contains the full match and all capture groups for each match
// 	matches := re.FindAllStringSubmatch(input, -1)

// 	// Flag to track whether multiplication operations are enabled
// 	mulEnabled := true
// 	// Running total of all valid multiplication results
// 	total := 0

// 	// Process each match and its capture groups
// 	for _, m := range matches {
// 		// Debug print for the first match (can be removed)
// 		switch {
// 		case m[4] == "do()":
// 			// Enable multiplication operations
// 			mulEnabled = true
// 		case m[5] == "don't()":
// 			// Disable multiplication operations
// 			mulEnabled = false
// 		case m[1] != "":
// 			// Process multiplication instruction if present (non-empty capture group 1)
// 			X, errX := strconv.Atoi(m[2])
// 			Y, errY := strconv.Atoi(m[3])
// 			// Only multiply and add to total if:
// 			// - Both numbers parsed successfully
// 			// - Multiplication operations are currently enabled
// 			if errX == nil && errY == nil && mulEnabled {
// 				total += X * Y
// 			}
// 		}
// 	}

// 	return total, nil
// }
