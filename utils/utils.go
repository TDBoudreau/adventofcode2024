package utils

import (
	"bufio"
	"os"
)

func ReadFile(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func Day05LoadInstructions(filename string) ([]string, []string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, []string{}, err
	}
	defer file.Close()

	var rules, pages []string
	var doneWithRules bool
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			doneWithRules = true
			continue
		}

		if !doneWithRules {
			rules = append(rules, text)
		} else {
			pages = append(pages, text)
		}
	}

	return rules, pages, nil
}
