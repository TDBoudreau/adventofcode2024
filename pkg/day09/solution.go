package day09

import (
	"fmt"
	"strconv"

	"github.com/tdboudreau/adventofcode2024/utils"
)

// Solve reads the input file and returns the results for both parts of the problem.
func Solve(inputPath string) (int, int, error) {
	discMap, err := utils.ReadFile("input/09/input.txt")
	// discMap, err := utils.ReadFile("input/09/ex.txt")
	if err != nil {
		return 0, 0, err
	}

	// Implement your logic here.
	// For now, return placeholder values.
	ans1, err := part1(discMap)
	if err != nil {
		return 0, 0, err
	}
	ans2, err := part2(discMap)
	if err != nil {
		return 0, 0, err
	}

	return ans1, ans2, nil
}

type Block struct {
	Index  int
	FileID int // -1 for spaces
	Length int
}

/*
-- Part 1 --

The disk map uses a dense format to represent the layout of files and
free space on the disk. The digits alternate between indicating the
length of a file and the length of free space.

So, a disk map like 12345 would represent a one-block file, two blocks of
free space, a three-block file, four blocks of free space, and then a
five-block file. A disk map like 90909 would represent three nine-block
files in a row (with no free space between them).

Each file on disk also has an ID number based on the order of the files as
they appear before they are rearranged, starting with ID 0. So, the disk
map 12345 has three files: a one-block file with ID 0, a three-block file
with ID 1, and a five-block file with ID 2. Using one character for each
block where digits are the file ID and . is free space.
*/
func part1(discMap string) (int, error) {
	// Logic for part1
	fileBlocks, _, _, err := parseBlocks(discMap)
	if err != nil {
		return 0, err
	}

	// fmt.Println(fileBlocks)

	// compactedFiles := compactBlocks(fileBlocks)

	// fmt.Println(compactedFiles)

	checksum := calculateChecksum(fileBlocks)

	return checksum, nil
	// 88217448737			-- too low? oooh, multi-digit fileIDs, how annoying
	// 154181994596019	-- too high.. hmmm
	// 6200294120911		-- **winner, winner**
}

func parseBlocks(discMap string) ([]Block, map[int]Block, int, error) {
	blockMap := make(map[int]Block)

	var blocks []Block
	var currentPos int
	fileID := 0

	for currentPos < len(discMap) {
		// Get block length
		length, err := strconv.Atoi(string(discMap[currentPos]))
		if err != nil {
			return nil, nil, 0, err
		}
		currentPos++

		// Add file block
		if length > 0 {
			block := Block{
				Index:  len(blocks),
				FileID: fileID,
				Length: length,
			}
			blocks = append(blocks, block)
			blockMap[fileID] = block

			fileID++
		}

		// Handle spaces if we're not at the end
		if currentPos < len(discMap) {
			spaces, err := strconv.Atoi(string(discMap[currentPos]))
			if err != nil {
				return nil, nil, 0, err
			}
			if spaces > 0 {
				blocks = append(blocks, Block{
					Index:  len(blocks),
					FileID: -1,
					Length: spaces,
				})
			}
			currentPos++
		}
	}

	fileID--

	return blocks, blockMap, fileID, nil
}

func compactBlocks(blocks []Block) []Block {
	// Convert blocks to positions
	var positions []Block
	for _, block := range blocks {
		for i := 0; i < block.Length; i++ {
			positions = append(positions, Block{FileID: block.FileID, Length: 1})
		}
	}

	// Compact from right to left
	for i := len(positions) - 1; i > 0; i-- {
		if positions[i].FileID == -1 {
			continue
		}

		for j := 0; j < i; j++ {
			if positions[j].FileID != -1 {
				continue
			}

			positions[j].FileID = positions[i].FileID
			positions[i].FileID = -1
			break
		}
	}

	// Merge consecutive blocks with same FileID
	var compacted []Block
	if len(positions) > 0 {
		current := positions[0]
		for i := 1; i < len(positions); i++ {
			if positions[i].FileID == current.FileID {
				current.Length++
			} else {
				if current.FileID != -1 {
					compacted = append(compacted, current)
				}
				current = positions[i]
			}
		}
		if current.FileID != -1 {
			compacted = append(compacted, current)
		}
	}

	return compacted
}

func calculateChecksum(blocks []Block) int {
	checksum := 0
	position := 0

	for _, block := range blocks {
		if block.FileID != -1 {
			// For each position in the block, add fileID * position to checksum
			for i := 0; i < block.Length; i++ {
				checksum += block.FileID * (position + i)
			}
		}
		position += block.Length
	}

	return checksum
}

/*
-- Part 2 --
This time, attempt to move whole files to the leftmost span of free space
blocks that could fit the file. Attempt to move each file exactly once
in order of decreasing file ID number starting with the file with the
highest file ID number. If there is no span of free space to the left of
a file that is large enough to fit the file, the file does not move.
*/
func part2(discMap string) (int, error) {
	// Logic for part2
	blocks, blockMap, maxID, err := parseBlocks(discMap)
	if err != nil {
		return 0, err
	}

	compactedBlocks := compactWholeBlocks(blocks, blockMap, maxID)
	fmt.Println(compactedBlocks)

	return 0, nil
}

func compactWholeBlocks(blocks []Block, blockMap map[int]Block, maxID int) []Block {
	fmt.Println("here")
	var positions []Block
	var compactedBlocks []Block
	for _, block := range blocks {
		compactedBlocks = append(compactedBlocks, block)

		for i := 0; i < block.Length; i++ {
			positions = append(positions, Block{FileID: block.FileID, Length: 1})
		}
	}

	for maxID > 0 {
		selectedBlock := blockMap[maxID]
		for i := 0; i < len(compactedBlocks); i++ {
			if compactedBlocks[i].FileID == -1 && compactedBlocks[i].Length <= selectedBlock.Length && i > 0 {
				if compactedBlocks[i].Length == selectedBlock.Length {
					// replace with selectedBlock
					compactedBlocks[i].FileID = selectedBlock.FileID
					compactedBlocks[selectedBlock.Index].FileID = -1
				} else {
					diff := compactedBlocks[i].Length - selectedBlock.Length
					if diff < 0 {
						diff = -diff
					}

					// replace with selectedBlock & update length
					compactedBlocks[i].FileID = selectedBlock.FileID
					compactedBlocks[i].Length = selectedBlock.Length

					// insert new FileID: -1 Block after block[i]
					preBlock := compactedBlocks[:i+1]
					postBlocks := compactedBlocks[i:]
					preBlock = append(preBlock, Block{FileID: -1, Length: diff, Index: i + 1})
					compactedBlocks = append(preBlock, postBlocks...)

					// update selected block index to reflect spaces
					selectedBlock.FileID = -1
				}
				break
			}
		}

		maxID--
	}

	return compactedBlocks
}
