package day05

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tdboudreau/adventofcode2024/utils"
)

// mapRules takes in a list of rules and maps them, the result maps each page number to
// a slice of pages numbers it must be printed *before.
func mapRules(rules []string) (map[int][]int, error) {
	mappedRules := make(map[int][]int)

	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		if len(parts) != 2 {
			return make(map[int][]int), fmt.Errorf("error mapping rules, issue splitting following rule: %s", rule)
		}

		pageX, err := strconv.Atoi(parts[0])
		if err != nil {
			return make(map[int][]int), err
		}
		pageY, err := strconv.Atoi(parts[1])
		if err != nil {
			return make(map[int][]int), err
		}

		_, ok := mappedRules[pageX]
		if !ok {
			mappedRules[pageX] = []int{}
		}

		val := mappedRules[pageX]
		val = append(val, pageY)
		mappedRules[pageX] = val
	}

	return mappedRules, nil
}

func convertPrintsToIntSlices(prints []string) ([][]int, error) {
	printSlice := [][]int{}

	for _, print := range prints {
		pages := strings.Split(print, ",")

		tempPrint := []int{}
		for _, page := range pages {
			p, err := strconv.Atoi(page)
			if err != nil {
				return [][]int{}, err
			}
			tempPrint = append(tempPrint, p)
		}

		printSlice = append(printSlice, tempPrint)
	}

	return printSlice, nil
}

/* Day 05 - Print Queue

We need to read in the printing rules.
All the printing rules are "X|Y", which just means
 page "X" comes before page "Y"

First approach:
	Use a map to store the page number and a list of page numbers it must appear before.
		- surprisingly worked.
*/

func AllParts() (int, int, error) {
	rulesInput, printsInput, err := utils.Day05LoadInstructions("input/05/input.txt")
	if err != nil {
		return 0, 0, err
	}

	rules, err := mapRules(rulesInput)
	if err != nil {
		return 0, 0, err
	}

	// fmt.Printf("rules:\t%v\n", rules)

	prints, err := convertPrintsToIntSlices(printsInput)
	if err != nil {
		return 0, 0, err
	}

	// fmt.Printf("\nprints:\t%v\n\n", prints)

	answer1, incorrectUpdates, err := Part1(rules, prints)
	if err != nil {
		return 0, 0, err
	}

	answer2, err := Part2(rules, incorrectUpdates)
	if err != nil {
		return 0, 0, err
	}

	return answer1, answer2, nil
}

func Part1(rules map[int][]int, prints [][]int) (int, [][]int, error) {
	var midSum int
	var correctPrints, incorrectPrints [][]int
	for i := 0; i < len(prints); i++ {
		correctOrder := true
		print := prints[i]

		for j := 0; j < len(print); j++ {
			page := print[j]
			// fmt.Printf("%v\t", page)
			pageRules, ok := rules[page]
			if !ok {
				// fmt.Println("No rules for page ", page)
				continue
			}

			for _, beforePage := range pageRules {
				for k := 0; k < j; k++ {
					if print[k] == beforePage {
						correctOrder = false
						break
					}
				}
				if !correctOrder {
					break
				}
			}
		}
		// fmt.Printf("\n")

		if correctOrder {
			correctPrints = append(correctPrints, print)
		} else {
			incorrectPrints = append(incorrectPrints, print)
		}
	}

	for _, print := range correctPrints {
		midIdx := len(print) / 2 // int division ('/') already gets floor, neat!
		midInt := print[midIdx]
		midSum += midInt
	}

	return midSum, incorrectPrints, nil
}

/*
Part 2 - Fix the ordering for incorrectly-ordered updates,

	then calculate the middle numbers from those, now correctly-ordered, updates.
*/
func Part2(rules map[int][]int, incorrectUpdates [][]int) (int, error) {
	for i := 0; i < len(incorrectUpdates); i++ {
		incorrectUpdates[i] = correctPrintOrder(rules, incorrectUpdates[i])
	}
	correctedUpdates := incorrectUpdates

	var correctedTotal int
	for _, update := range correctedUpdates {
		mid := len(update) / 2
		correctedTotal += update[mid]
	}

	return correctedTotal, nil
	// 4839 -- too high
	// 4828 **winner winner**
}

/*
How to correct the ordering using the provided rules map?.. hmm
*/
func correctPrintOrder(rules map[int][]int, print []int) []int {
	// Determine the set of pages in this update
	pageSet := make(map[int]bool)
	for _, p := range print {
		pageSet[p] = true
	}

	// Build a graph (adjacency list) and in-degree map for only these pages
	adj := make(map[int][]int)
	inDegree := make(map[int]int)
	for p := range pageSet {
		adj[p] = []int{}
		inDegree[p] = 0
	}

	// Fill in edges for relevant pages
	for x, ys := range rules {
		if pageSet[x] {
			for _, y := range ys {
				if pageSet[y] {
					adj[x] = append(adj[x], y)
					inDegree[y]++
				}
			}
		}
	}

	// Kahn's Algorithm for Topological Sort:
	var queue []int
	for node, deg := range inDegree {
		if deg == 0 {
			queue = append(queue, node)
		}
	}

	var sorted []int
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		sorted = append(sorted, current)

		for _, neighbor := range adj[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// If sorted doesn't contain all pages, there's a cycle or error.
	if len(sorted) != len(print) {
		// Fallback, though it shouldn't happen if input is valid.
		return print
	}

	return sorted
}
