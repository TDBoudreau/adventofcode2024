package day03

import (
	"strconv"

	"github.com/tdboudreau/adventofcode2024/utils"
)

func Part1() (int, error) {
	input, err := utils.ReadFile("inputs/day03.txt")
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

// Just tries to parse something like "X,Y)" right after "mul("
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

	inside := s[:closeIndex]
	commaIndex := -1
	for idx, ch := range inside {
		if ch == ',' {
			if commaIndex == -1 {
				commaIndex = idx
			} else {
				return 0, 0
			}
		}
	}

	if commaIndex == -1 {
		return 0, 0
	}

	Xstr := inside[:commaIndex]
	Ystr := inside[commaIndex+1:]

	if !isValidNumber(Xstr) || !isValidNumber(Ystr) {
		return 0, 0
	}

	X, errX := strconv.Atoi(Xstr)
	Y, errY := strconv.Atoi(Ystr)
	if errX != nil || errY != nil {
		return 0, 0
	}

	length := closeIndex + 1
	return X * Y, length
}

// Just checks that the string is digits, length 1-3
func isValidNumber(numStr string) bool {
	if len(numStr) < 1 || len(numStr) > 3 {
		return false
	}
	for _, ch := range numStr {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
}

// // Regex method
// func Part1() (int, error) {
// 	input, err := utils.ReadFile("inputs/03/input.txt")
// 	if err != nil {
// 		return 0, err
// 	}

// 	re := regexp.MustCompile(`(mul\(([0-9]{1,3}),([0-9]{1,3})\))|(do\(\))|(don't\(\))`)
// 	matches := re.FindAllStringSubmatch(input, -1)

// 	mulEnabled := true
// 	total := 0

// 	for i, m := range matches {
// 		if i == 0 {

// 			fmt.Println(m)
// 			fmt.Println(m[1])
// 		}
// 		switch {
// 		case m[4] == "do()":
// 			// Enable future mul instructions
// 			mulEnabled = true
// 		case m[5] == "don't()":
// 			// Disable future mul instructions
// 			mulEnabled = false
// 		case m[1] != "":
// 			// This is a mul instruction
// 			X, errX := strconv.Atoi(m[2])
// 			Y, errY := strconv.Atoi(m[3])
// 			if errX == nil && errY == nil && mulEnabled {
// 				total += X * Y
// 			}
// 		}
// 	}

// 	return total, nil
// }
