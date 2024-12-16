package day08

import (
	"fmt"
	"strings"

	"github.com/tdboudreau/adventofcode2024/utils"
)

type Position struct {
	x int
	y int
}

// type Antenna struct {
// 	Coordinates Position
// 	Symbol      string
// }

func createMap(lines []string) ([][]string, map[string][]Position) {
	var cityMap [][]string
	antennas := make(map[string][]Position)

	for y, line := range lines {
		lineParts := strings.Split(line, "")
		for x, c := range lineParts {
			if c != "." {
				antennas[c] = append(antennas[c], Position{x: x, y: y})
			}
		}
		fmt.Println(line)
		cityMap = append(cityMap, lineParts)
	}

	return cityMap, antennas
}

// Solve reads the input file and returns the results for both parts of the problem.
func Solve(inputPath string) (int, int, error) {
	lines, err := utils.ReadLines("input/08/input.txt")
	// lines, err := utils.ReadLines("input/08/ex.txt")
	if err != nil {
		return 0, 0, err
	}

	// Implement your logic here.
	// For now, return placeholder values.
	city, coords := createMap(lines)

	ans1, err := part1(city, coords)
	if err != nil {
		return 0, 0, err
	}
	ans2, err := part2(city, coords)
	if err != nil {
		return 0, 0, err
	}

	return ans1, ans2, nil
}

/*
-- Part 1 --
Antenna (character symbols other than #) create 'antinodes'.

	An antinode occurs at any point that is perfectly in line with two antennas of the
	same frequency - but only when one of the antennas is twice as far away as the other.

Example:
Answer:
_________________________
|. . . . . . # . . . . #|
|. . . # . . . . 0 . . .|
|. . . . # 0 . . . . # .|
|. . # . . . . 0 . . . .|
|. . . . 0 . . . . # . .|
|. # . . . . A . . . . .|
|. . . # . . . . . . . .|
|# . . . . . . # . . . .|
|. . . . . . . . A . . .|
|. . . . . . . . . A . .|
|. . . . . . . . . . # .|
|. . . . . . . . . . # .|
-------------------------

My output:
[. . . . . . . . . . # .]
[. . . . . . . . 0 . . #]
[. . . . . 0 . . . # . .]
[. . . # . . . 0 . . . .]
[# # . . 0 . . . . . . .]
[. . # . # . A . . . . .]
[# # . . . . . . . . . .]
[. . . # . . . . . . . .]
[. . . . . . . . A . . .]
[. . . . . . . . . A . .]
[. . . . . . . . . . . .]
[. . . . . . . . . . . .]

In this example, an antinode is overlapping the topmost A-frequency antenna.
*/
func part1(city [][]string, antennas map[string][]Position) (int, error) {
	// antinodes will represent all the antinodes at a single position
	antinodes := make(map[Position][]string)

	// map bounds, we don't count antinodes outside these bounds
	mapHeight := len(city)
	mapWidth := len(city[0])

	// create and map antinodes
	for key, value := range antennas {
		for y := 0; y < len(value)-1; y++ {
			for x := y + 1; x < len(value); x++ {
				antenna1 := value[y]
				antenna2 := value[x]
				// each antenna pair creates 2 antinodes
				antinodes = addAntinodes(antenna1, antenna2, antinodes, Position{x: mapWidth, y: mapHeight}, key)
			}
		}
	}

	// lets create a temp map and overlay our antinodes so we can compare against the example.
	newCity := make([][]string, len(city))
	for i := range city {
		newCity[i] = make([]string, len(city[i]))
		copy(newCity[i], city[i])
	}

	fmt.Println()
	for pos := range antinodes {
		if pos.x >= 0 && pos.x < mapWidth && pos.y >= 0 && pos.y < mapHeight {
			newCity[pos.y][pos.x] = "#"
		}
	}

	// print new city with the antinodes
	for _, line := range newCity {
		for _, c := range line {
			fmt.Printf("%s", c)
		}
		fmt.Printf("\n")
	}

	countAnti := 0
	for pos := range antinodes {
		if pos.x >= 0 && pos.x < mapWidth && pos.y >= 0 && pos.y < mapHeight {
			countAnti++
		}
	}
	return countAnti, nil
}

func addAntinodes(a1, a2 Position, antinodes map[Position][]string, mapBounds Position, freq string) map[Position][]string {
	// First antinode: 2*a1 - a2
	antinode1 := Position{x: 2*a1.x - a2.x, y: 2*a1.y - a2.y}
	_, has := antinodes[antinode1]
	if !has && antinode1.x >= 0 && antinode1.x < mapBounds.x && antinode1.y >= 0 && antinode1.y < mapBounds.y {
		antinodes[antinode1] = append(antinodes[antinode1], freq)
	}

	// Second antinode: 2*a2 - a1
	antinode2 := Position{x: 2*a2.x - a1.x, y: 2*a2.y - a1.y}
	_, has = antinodes[antinode2]
	if !has && antinode2.x >= 0 && antinode2.x < mapBounds.x && antinode2.y >= 0 && antinode2.y < mapBounds.y {
		antinodes[antinode2] = append(antinodes[antinode2], freq)
	}

	return antinodes
}

// func abs(a int) int {
// 	if a >= 0 {
// 		return a
// 	}
// 	return -a
// }

/*
1,9
2,6

anti1 = b.x-a.x, b.y+b.y
anti2 = b.x+a.x, b.y-b.y

anti1 = 1, 15
anti2 = 3, 3

*/

/*
-- Part 2 --
*/
func part2(city [][]string, antennas map[string][]Position) (int, error) {
	// antinodes will represent all the antinodes at a single position
	antinodes := make(map[Position][]string)

	// map bounds, we don't count antinodes outside these bounds
	mapHeight := len(city)
	mapWidth := len(city[0])

	// create and map antinodes
	for key, value := range antennas {
		for y := 0; y < len(value)-1; y++ {
			for x := y + 1; x < len(value); x++ {
				antenna1 := value[y]
				antenna2 := value[x]
				// each antenna pair creates 2 antinodes
				antinodes = addAntinodesLine(antenna1, antenna2, antinodes, Position{x: mapWidth, y: mapHeight}, key)
			}
		}
	}

	// lets create a temp map and overlay our antinodes so we can compare against the example.
	newCity := make([][]string, len(city))
	for i := range city {
		newCity[i] = make([]string, len(city[i]))
		copy(newCity[i], city[i])
	}

	fmt.Println()
	for pos := range antinodes {
		if pos.x >= 0 && pos.x < mapWidth && pos.y >= 0 && pos.y < mapHeight && newCity[pos.y][pos.x] == "." {
			newCity[pos.y][pos.x] = "#"
		}
	}

	// print new city with the antinodes
	for _, line := range newCity {
		for _, c := range line {
			fmt.Printf("%s", c)
		}
		fmt.Printf("\n")
	}

	countAnti := 0
	for pos := range antinodes {
		if pos.x >= 0 && pos.x < mapWidth && pos.y >= 0 && pos.y < mapHeight {
			countAnti++
		}
	}
	return countAnti, nil
}

func addAntinodesLine(a1, a2 Position, antinodes map[Position][]string, mapBounds Position, freq string) map[Position][]string {
	maxX, maxY := mapBounds.x, mapBounds.y

	// Add antinodes ontop of antenna
	antinodes[a1] = append(antinodes[a1], freq)
	antinodes[a2] = append(antinodes[a2], freq)

	// Direction vectors
	dxAB := a1.x - a2.x
	dyAB := a1.y - a2.y
	dxBA := a2.x - a1.x
	dyBA := a2.y - a1.y

	// First antinode on A side
	currentA := Position{x: 2*a1.x - a2.x, y: 2*a1.y - a2.y}
	// Keep stepping until out of bounds
	for currentA.x >= 0 && currentA.x < maxX && currentA.y >= 0 && currentA.y < maxY {
		antinodes[currentA] = append(antinodes[currentA], freq)
		// Move one step further out along the line in the A direction
		currentA.x += dxAB
		currentA.y += dyAB
	}

	// First antinode on B side
	currentB := Position{x: 2*a2.x - a1.x, y: 2*a2.y - a1.y}
	// Keep stepping until out of bounds
	for currentB.x >= 0 && currentB.x < maxX && currentB.y >= 0 && currentB.y < maxY {
		antinodes[currentB] = append(antinodes[currentB], freq)
		// Move one step further out along the line in the B direction
		currentB.x += dxBA
		currentB.y += dyBA
	}

	return antinodes
}
