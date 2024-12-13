package day07

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/tdboudreau/adventofcode2024/utils"
)

func parseEquations(lines []string) (map[int][]int, error) {
	equations := make(map[int][]int)

	for i, line := range lines {
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			return equations, fmt.Errorf("invalid format at line: %d", i+1)
		}

		testValue, err := strconv.Atoi(parts[0])

		_, has := equations[testValue]
		if has {
			panic("input has duplicate test values, cannot use map")
		}

		if err != nil {
			return equations, err
		}

		numbers := strings.Split(parts[1], " ")
		if len(numbers) < 1 {
			return equations, fmt.Errorf("invalid number of input numbers at line: %d", i)
		}

		var nums []int
		for _, number := range numbers {
			n, err := strconv.Atoi(number)
			if err != nil {
				return equations, err
			}

			nums = append(nums, n)
		}

		equations[testValue] = nums
	}

	return equations, nil
}

// Solve reads the input file and returns the results for both parts of the problem.
func Solve(inputPath string) (int, int, error) {
	lines, err := utils.ReadLines("input/07/input.txt")
	// lines, err := utils.ReadLines("input/07/ex1.txt")
	if err != nil {
		return 0, 0, err
	}

	equations, err := parseEquations(lines)
	if err != nil {
		return 0, 0, err
	}

	ans1, err := part1(equations)
	if err != nil {
		return 0, 0, err
	}
	ans2, err := part2(equations)

	return ans1, ans2, err
}

/*
-- Part 1 --

To brute force check all operator positions, we would

	need to check at least 2(n - 1) combinations.
*/
func part1(equations map[int][]int) (int, error) {
	var validTests []int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for test, values := range equations {
		wg.Add(1)

		go func(test int, values []int) {
			defer wg.Done()
			if tryCombination(values, 0, 0, test) {
				mu.Lock()
				validTests = append(validTests, test)
				mu.Unlock()
			}
		}(test, values)
	}

	wg.Wait()

	var sum int
	for _, test := range validTests {
		sum += test
	}

	return sum, nil
	// 28703334998584 -- too low
	// 28730327770375 ** winner winner **
}

func tryCombination(numbers []int, position, currentResult, target int) bool {
	if position == len(numbers) {
		return currentResult == target
	}

	if tryCombination(numbers, position+1, currentResult+numbers[position], target) {
		return true
	}

	if tryCombination(numbers, position+1, currentResult*numbers[position], target) {
		return true
	}

	return false
}

func part2(equations map[int][]int) (int, error) {
	var validTests []int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for test, values := range equations {
		wg.Add(1)

		go func(test int, values []int) {
			defer wg.Done()
			if tryCombinationWithConcat(values, test) {
				mu.Lock()
				validTests = append(validTests, test)
				mu.Unlock()
			}
		}(test, values)
	}

	wg.Wait()

	var sum int
	for _, test := range validTests {
		sum += test
	}

	return sum, nil
}

func tryCombinationWithConcat(numbers []int, target int) bool {
	// Start the recursion from the second number since
	// we take the first number as our initial result.
	if len(numbers) == 1 {
		// If there's only one number, just check it directly
		return numbers[0] == target
	}
	return recurseWithConcat(numbers, 1, numbers[0], target)
}

// recurseWithConcat tries all operators between the previous computed result
// and numbers[index], returning true if it can eventually reach the target.
func recurseWithConcat(numbers []int, index int, currentResult int, target int) bool {
	// If we've placed operators between all pairs of numbers:
	if index == len(numbers) {
		// Check if the currentResult matches the target
		return currentResult == target
	}

	nextNum := numbers[index]

	// Try addition
	if recurseWithConcat(numbers, index+1, currentResult+nextNum, target) {
		return true
	}

	// Try multiplication
	if recurseWithConcat(numbers, index+1, currentResult*nextNum, target) {
		return true
	}

	// Try concatenation
	// Concatenation means merging the currentResult and nextNum as if they were digits in a number.
	// For example: currentResult = 48, nextNum = 6 => newResult = 486
	// One way: convert to string, concatenate, then convert back.
	// Or do mathematically:
	// newResult = currentResult * (10^(digit_count(nextNum))) + nextNum
	concatResult := concatNumbers(currentResult, nextNum)
	return recurseWithConcat(numbers, index+1, concatResult, target)
}

// concatNumbers concatenates two integers as if they were strings of digits.
// e.g. concatNumbers(48, 6) = 486
func concatNumbers(a, b int) int {
	// Count digits of b
	digits := 0
	temp := b
	if temp == 0 {
		digits = 1
	} else {
		for temp > 0 {
			temp /= 10
			digits++
		}
	}

	// a * 10^digits + b
	multiplier := 1
	for i := 0; i < digits; i++ {
		multiplier *= 10
	}

	return a*multiplier + b
}
