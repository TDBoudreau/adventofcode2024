#!/usr/bin/env bash

# Usage: ./add_day.sh <day_number>
#
# Example:
#   ./add_day.sh 1
#   This will create structures for day01.
#
#   ./add_day.sh 09
#   This will create structures for day09.

# Exit on errors
set -e

# Check if day number is provided
if [ -z "$1" ]; then
    echo "Error: No day number provided."
    echo "Usage: $0 <day_number>"
    exit 1
fi

# Zero-pad the day number
DAY=$(printf "%02d" "$1")

# Define directories and files
CMD_DIR="cmd/day${DAY}"
PKG_DIR="pkg/day${DAY}"
INPUT_DIR="input/${DAY}"
PROBLEM_FILE="problems/${DAY}.html"

# Create directories if they don't exist
mkdir -p "$CMD_DIR"
mkdir -p "$PKG_DIR"
mkdir -p "$INPUT_DIR"
mkdir -p "$(dirname "$PROBLEM_FILE")"

# Create empty or template files
MAIN_GO_FILE="${CMD_DIR}/main.go"
if [ ! -f "$MAIN_GO_FILE" ]; then
    cat > "$MAIN_GO_FILE" <<EOF
package main

import (
    "fmt"
    "os"
    "path/filepath"

    "github.com/tdboudreau/adventofcode2024/pkg/day${DAY}"
)

func main() {
    // Example of how you might run your solution
    exePath, _ := os.Executable()
    dir := filepath.Dir(exePath)
    inputPath := filepath.Join(dir, "../../input/${DAY}", "input.txt")
    
    resultPart1, resultPart2, err := day${DAY}.Solve(inputPath)
    if err != nil {
      panic(err)
    }

    fmt.Printf("Day ${DAY}, Part 1: %v\n", resultPart1)
    fmt.Printf("Day ${DAY}, Part 2: %v\n", resultPart2)
}
EOF
    echo "Created $MAIN_GO_FILE"
else
    echo "$MAIN_GO_FILE already exists."
fi

SOLUTION_GO_FILE="${PKG_DIR}/solution.go"
if [ ! -f "$SOLUTION_GO_FILE" ]; then
    cat > "$SOLUTION_GO_FILE" <<EOF
package day${DAY}

import "github.com/tdboudreau/adventofcode2024/utils"

// Solve reads the input file and returns the results for both parts of the problem.
func Solve(inputPath string) (int, int, error) {
  // lines, err := utils.ReadLines("input/${DAY}/input.txt")
  lines, err := utils.ReadLines("input/${DAY}/ex1.txt")
  if err != nil {
    return 0, 0, err
  }

  // Implement your logic here.
  // For now, return placeholder values.
  ans1, err := part1()
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
*/
func part1() (int, error) {
  // Logic for part1

  return 0, nil
}

/*
-- Part 2 --
*/
func part2() (int, error) {
  // Logic for part2

  return 0, nil
}
EOF
    echo "Created $SOLUTION_GO_FILE"
else
    echo "$SOLUTION_GO_FILE already exists."
fi

INPUT_TXT_FILE="${INPUT_DIR}/input.txt"
if [ ! -f "$INPUT_TXT_FILE" ]; then
    touch "$INPUT_TXT_FILE"
    echo "Created $INPUT_TXT_FILE"
else
    echo "$INPUT_TXT_FILE already exists."
fi

INPUT_EX_TXT_FILE="${INPUT_DIR}/ex.txt"
if [ ! -f "$INPUT_EX_TXT_FILE" ]; then
    touch "$INPUT_EX_TXT_FILE"
    echo "Created $INPUT_EX_TXT_FILE"
else
    echo "$INPUT_EX_TXT_FILE already exists."
fi

if [ ! -f "$PROBLEM_FILE" ]; then
    cat > "$PROBLEM_FILE" <<EOF
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <link rel="stylesheet" href="../static/style.css" />
    <title>Problem ${DAY}</title>
  </head>
  <body>
    <main>
      <!-- Part 1 Article -->
      
      <!-- Part 2 Article -->
    </main>
  </body>
</html>
EOF
    echo "Created $PROBLEM_FILE"
else
    echo "$PROBLEM_FILE already exists."
fi

echo "Setup complete for Day ${DAY}."
