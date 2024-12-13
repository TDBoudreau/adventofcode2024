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
	// lines, err := utils.ReadLines("input/08/input.txt")
	lines, err := utils.ReadLines("input/08/ex.txt")
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
	maxX := mapBounds.x
	maxY := mapBounds.y

	// Compute the two antinodes:
	p1 := Position{x: 2*a1.x - a2.x, y: 2*a1.y - a2.y}
	p2 := Position{x: 2*a2.x - a1.x, y: 2*a2.y - a1.y}

	dx := p2.x - p1.x
	dy := p2.y - p1.y

	// Handle special cases: vertical or horizontal lines
	if dx == 0 && dy == 0 {
		// Both antinodes are at the same point (very unusual), just add that one if inside
		if p1.x >= 0 && p1.x < maxX && p1.y >= 0 && p1.y < maxY {
			antinodes[p1] = append(antinodes[p1], freq)
		}
		return antinodes
	}

	if dx == 0 {
		// Vertical line: x = p1.x
		x := p1.x
		if x >= 0 && x < maxX {
			// Covers entire height
			for y := 0; y < maxY; y++ {
				antinodes[Position{x, y}] = append(antinodes[Position{x, y}], freq)
			}
		}
		return antinodes
	}

	if dy == 0 {
		// Horizontal line: y = p1.y
		y := p1.y
		if y >= 0 && y < maxY {
			// Covers entire width
			for x := 0; x < maxX; x++ {
				antinodes[Position{x, y}] = append(antinodes[Position{x, y}], freq)
			}
		}
		return antinodes
	}

	// For a general line, find intersections with the map boundary:
	tValues := []float64{}

	// Compute intersection t values if possible:
	// Left boundary (X=0)
	if dx != 0 {
		tLeft := float64(-p1.x) / float64(dx)
		Yleft := float64(p1.y) + tLeft*float64(dy)
		if Yleft >= 0 && Yleft <= float64(maxY-1) {
			tValues = append(tValues, tLeft)
		}

		// Right boundary (X=maxX-1)
		tRight := float64((maxX-1)-p1.x) / float64(dx)
		Yright := float64(p1.y) + tRight*float64(dy)
		if Yright >= 0 && Yright <= float64(maxY-1) {
			tValues = append(tValues, tRight)
		}
	}

	// Top boundary (Y=0)
	if dy != 0 {
		tTop := float64(-p1.y) / float64(dy)
		Xtop := float64(p1.x) + tTop*float64(dx)
		if Xtop >= 0 && Xtop <= float64(maxX-1) {
			tValues = append(tValues, tTop)
		}

		// Bottom boundary (Y=maxY-1)
		tBottom := float64((maxY-1)-p1.y) / float64(dy)
		Xbottom := float64(p1.x) + tBottom*float64(dx)
		if Xbottom >= 0 && Xbottom <= float64(maxX-1) {
			tValues = append(tValues, tBottom)
		}
	}

	if len(tValues) == 0 {
		// No intersection within the map? The line might be entirely outside.
		return antinodes
	}

	// We need the min and max t that define the segment inside the map
	minT, maxT := minMax(tValues)

	// Convert these t-values into actual points
	startX := int(float64(p1.x) + minT*float64(dx))
	startY := int(float64(p1.y) + minT*float64(dy))
	endX := int(float64(p1.x) + maxT*float64(dx))
	endY := int(float64(p1.y) + maxT*float64(dy))

	// Now use Bresenham's line algorithm from (startX,startY) to (endX,endY)
	linePoints := bresenhamLine(startX, startY, endX, endY)

	// Add all these points to antinodes
	for _, pt := range linePoints {
		if pt.x >= 0 && pt.x < maxX && pt.y >= 0 && pt.y < maxY {
			antinodes[pt] = append(antinodes[pt], freq)
		}
	}

	return antinodes
}

func minMax(arr []float64) (float64, float64) {
	min := arr[0]
	max := arr[0]
	for _, v := range arr[1:] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func bresenhamLine(x0, y0, x1, y1 int) []Position {
	var points []Position

	dx := absInt(x1 - x0)
	sx := 1
	if x0 > x1 {
		sx = -1
	}
	dy := -absInt(y1 - y0)
	sy := 1
	if y0 > y1 {
		sy = -1
	}
	err := dx + dy

	x, y := x0, y0
	for {
		points = append(points, Position{x, y})
		if x == x1 && y == y1 {
			break
		}
		e2 := 2 * err
		if e2 >= dy {
			err += dy
			x += sx
		}
		if e2 <= dx {
			err += dx
			y += sy
		}
	}

	return points
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}