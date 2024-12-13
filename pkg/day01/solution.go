package day01

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/tdboudreau/adventofcode2024/utils"
)

func loadSlices(lines []string) ([]int, []int, error) {
	var leftSlice, rightSlice []int
	var left, right int
	var err error

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {

			return []int{}, []int{}, fmt.Errorf("error splitting into parts: %s", line)
		}

		left, err = strconv.Atoi(parts[0])
		if err != nil {
			return []int{}, []int{}, err
		}
		right, err = strconv.Atoi(parts[1])
		if err != nil {
			return []int{}, []int{}, err
		}

		leftSlice = append(leftSlice, left)
		rightSlice = append(rightSlice, right)
	}

	return leftSlice, rightSlice, nil
}

func diff(x, y int) int {
	if x > y {
		return x - y
	}

	return y - x
}

func Part1() (int, error) {
	lines, err := utils.ReadLines("input/01/input.txt")
	if err != nil {
		return 0, err
	}

	leftSlice, rightSlice, err := loadSlices(lines)
	if err != nil {
		return 0, err
	}

	slices.Sort(leftSlice)
	slices.Sort(rightSlice)

	var sum int

	for i := range leftSlice {
		sum += diff(leftSlice[i], rightSlice[i])
	}

	return sum, nil
}

func Part2() (int, error) {
	lines, err := utils.ReadLines("inputs/day01.txt")
	if err != nil {
		return 0, err
	}

	leftSlice, rightSlice, err := loadSlices(lines)
	if err != nil {
		return 0, err
	}

	distribution := make(map[int]int)

	for _, num := range rightSlice {
		_, exists := distribution[num]
		if !exists {
			distribution[num] = 1
			continue
		}

		distribution[num] += 1
	}

	var similarityScore int

	for _, num := range leftSlice {
		value, exists := distribution[num]
		if !exists {
			continue
		}

		similarityScore += num * value
	}

	return similarityScore, nil
}
