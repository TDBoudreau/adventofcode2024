package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tdboudreau/adventofcode2024/utils"
)

func parseData(lines []string) ([][]int, error) {
	// Pre-allocate the slice based on input size
	numsSlice := make([][]int, 0, len(lines))
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 0 {
			return nil, fmt.Errorf("error splitting line: %s", line)
		}

		// Pre-allocate the nums slice
		nums := make([]int, 0, len(parts))
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("error: couldn't convert part: %s", part)
			}
			nums = append(nums, num)
		}
		numsSlice = append(numsSlice, nums)
	}
	return numsSlice, nil
}

// The levels are either 'all increasing' or 'all decreasing'.
// Any two adjacent levels differ by 'at least one' and 'at most three'.

func checkReport(report []int) bool {
	if len(report) < 2 {
		return true
	}

	decreasing := report[0] > report[1]
	prev := report[0]

	for i := 1; i < len(report); i++ {
		curr := report[i]
		difference := diff(curr, prev)

		if difference == 0 || difference > 3 ||
			(decreasing && curr > prev) || (!decreasing && curr < prev) {
			return false
		}

		prev = curr
	}

	return true
}

func diff(x, y int) int {
	if x > y {
		return x - y
	}

	return y - x
}

func modifyReport(report []int) bool {
	if len(report) <= 2 {
		return true
	}

	// Create a reusable slice with capacity for our modified reports
	newReport := make([]int, 0, len(report)-1)

	for i := 0; i < len(report); i++ {
		// Reset the slice length while keeping capacity
		newReport = newReport[:0]

		// Build modified report without the current element
		newReport = append(newReport, report[:i]...)   // add everything before i: [4, 5, 6, i, ...]
		newReport = append(newReport, report[i+1:]...) // add everything after i: [..., i, 8, 9, 10]

		if checkReport(newReport) {
			return true
		}
	}
	return false
}

func Part1() (int, error) {
	lines, err := utils.ReadLines("inputs/day02.txt")
	if err != nil {
		return 0, err
	}

	numSlices, err := parseData(lines)
	if err != nil {
		return 0, err
	}

	numSafeReports := 0
	for _, numSlice := range numSlices {
		if checkReport(numSlice) {
			numSafeReports++
		}
	}
	return numSafeReports, nil
}

func Part2() (int, error) {
	lines, err := utils.ReadLines("inputs/day02.txt")
	if err != nil {
		return 0, err
	}

	numSlices, err := parseData(lines)
	if err != nil {
		return 0, err
	}

	numSafeReports := 0
	for _, numSlice := range numSlices {
		if modifyReport(numSlice) {
			numSafeReports++
		}
	}
	return numSafeReports, nil
}
