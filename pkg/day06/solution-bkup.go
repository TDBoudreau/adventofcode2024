package day06

// func part2(guardMap [][]string, startPos Position) (int, error) {
// 	running := true
// 	var loopCounter int
// 	pos := startPos // Create a local copy of the position

// 	for running {
// 		// Create a deep copy of the map for testing
// 		testMap := make([][]string, len(guardMap))
// 		for i := range guardMap {
// 			testMap[i] = make([]string, len(guardMap[i]))
// 			copy(testMap[i], guardMap[i])
// 		}

// 		// Create a copy of current position for testing
// 		testPos := Position{
// 			x:         pos.x,
// 			y:         pos.y,
// 			direction: pos.direction,
// 		}

// 		// Test for loop with copies
// 		testMap = addObstacle(testMap, testPos)
// 		if checkForLoop(testMap, testPos) {
// 			loopCounter++
// 			fmt.Printf("Loop found: %v\tloops: %d\n", pos, loopCounter)
// 		}

// 		// Update the actual position
// 		moved := false
// 		switch pos.direction {
// 		case "^":
// 			if pos.y > 0 && guardMap[pos.y-1][pos.x] != "#" {
// 				pos.y--
// 				moved = true
// 			}
// 		case ">":
// 			if pos.x < len(guardMap[0])-1 && guardMap[pos.y][pos.x+1] != "#" {
// 				pos.x++
// 				moved = true
// 			}
// 		case "v":
// 			if pos.y < len(guardMap)-1 && guardMap[pos.y+1][pos.x] != "#" {
// 				pos.y++
// 				moved = true
// 			}
// 		case "<":
// 			if pos.x > 0 && guardMap[pos.y][pos.x-1] != "#" {
// 				pos.x--
// 				moved = true
// 			}
// 		}

// 		if !moved {
// 			if (pos.direction == "^" && pos.y == 0) ||
// 				(pos.direction == ">" && pos.x == len(guardMap[0])-1) ||
// 				(pos.direction == "v" && pos.y == len(guardMap)-1) ||
// 				(pos.direction == "<" && pos.x == 0) {
// 				running = false
// 			} else {
// 				pos.rotate()
// 			}
// 		}
// 	}

// 	return loopCounter, nil
// }

// func addObstacle(guardMap [][]string, startPosition Position) [][]string {
// 	loopMap := guardMap
// 	pos := startPosition

// 	switch pos.direction {
// 	case "^":
// 		if pos.y > 0 && loopMap[pos.y-1][pos.x] != "#" {
// 			loopMap[pos.y-1][pos.x] = "#"
// 		}
// 	case ">":
// 		if pos.x < len(loopMap[0])-1 && loopMap[pos.y][pos.x+1] != "#" {
// 			loopMap[pos.y][pos.x+1] = "#"
// 		}
// 	case "v":
// 		if pos.y < len(loopMap)-1 && loopMap[pos.y+1][pos.x] != "#" {
// 			loopMap[pos.y+1][pos.x] = "#"
// 		}
// 	case "<":
// 		if pos.x > 0 && loopMap[pos.y][pos.x-1] != "#" {
// 			loopMap[pos.y][pos.x-1] = "#"
// 		}
// 	}

// 	return loopMap
// }

// func checkForLoop(guardMap [][]string, startPostition Position) bool {
// 	pos := startPostition
// 	running := true

// 	var looping bool

// 	for i := 0; running; i++ {
// 		switch pos.direction {
// 		case "^":
// 			if pos.y > 0 {
// 				if guardMap[pos.y-1][pos.x] != "#" {
// 					guardMap[pos.y-1][pos.x] = "^"
// 					guardMap[pos.y][pos.x] = "X"
// 					pos.y = pos.y - 1
// 				} else {
// 					pos.rotate()
// 				}
// 			} else {
// 				guardMap[pos.y][pos.x] = "X"
// 				running = false
// 			}
// 		case ">":
// 			if pos.x < len(guardMap[0])-1 {
// 				if guardMap[pos.y][pos.x+1] != "#" {
// 					guardMap[pos.y][pos.x+1] = ">"
// 					guardMap[pos.y][pos.x] = "X"
// 					pos.x = pos.x + 1
// 				} else {
// 					pos.rotate()
// 				}
// 			} else {
// 				guardMap[pos.y][pos.x] = "X"
// 				running = false
// 			}
// 		case "v":
// 			if pos.y < len(guardMap)-1 {
// 				if guardMap[pos.y+1][pos.x] != "#" {
// 					guardMap[pos.y+1][pos.x] = "v"
// 					guardMap[pos.y][pos.x] = "X"
// 					pos.y = pos.y + 1
// 				} else {
// 					pos.rotate()
// 				}
// 			} else {
// 				guardMap[pos.y][pos.x] = "X"
// 				running = false
// 			}
// 		case "<":
// 			if pos.x > 0 {
// 				if guardMap[pos.y][pos.x-1] != "#" {
// 					guardMap[pos.y][pos.x-1] = "<"
// 					guardMap[pos.y][pos.x] = "X"
// 					pos.x = pos.x - 1
// 				} else {
// 					pos.rotate()
// 				}
// 			} else {
// 				guardMap[pos.y][pos.x] = "X"
// 				running = false
// 			}
// 		}

// 		if i > 0 && pos.isEqual(startPostition) {
// 			looping = true
// 			running = false
// 		}
// 	}

// 	return looping
// }
