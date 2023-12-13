package _2023

import (
	"os"
	"strings"

	"github.com/pauldolden/advent-go/config"
)

func ThirteenTwo(o config.Options) int {
	bs, err := os.ReadFile("./2023/13.txt")
	if err != nil {
		panic(err)
	}

	s := string(bs)

	ss := strings.Split(s, "\n\n")

	var total int

	for _, block := range ss {
		grid := parseBlock(block)

		mirror := findMirror(grid)

		total += 100 * mirror

		// Rotate the grid
		transposed := transpose(grid)
		mirror = findMirror(transposed)

		total += mirror
	}
	return total
}

func transpose(grid [][]string) [][]string {
	if len(grid) == 0 {
		return nil
	}

	transposed := make([][]string, len(grid[0]))
	for i := range transposed {
		transposed[i] = make([]string, len(grid))
		for j := range grid {
			transposed[i][j] = grid[j][i]
		}
	}

	return transposed
}

func findMirror(grid [][]string) int {
	for r := 1; r < len(grid); r++ {
		above := reverseGrid(grid[:r])
		below := grid[r:]

		// slice above and below so their lengths match
		if len(above) > len(below) {
			above = above[:len(below)]
		}

		if len(below) > len(above) {
			below = below[:len(above)]
		}

		diffCount := 0
		for i := 0; i < len(above); i++ {
			for j := 0; j < len(above[i]); j++ {
				if above[i][j] != below[i][j] {
					diffCount++
				}
			}
		}

		if diffCount == 1 {
			return r
		}
	}
	return 0
}

func parseBlock(block string) [][]string {
	lines := strings.Split(block, "\n")
	var grid [][]string

	for _, line := range lines {
		if line != "" {
			grid = append(grid, strings.Split(line, ""))
		}
	}
	return grid
}

func reverseGrid(grid [][]string) [][]string {
	length := len(grid)
	reversed := make([][]string, length)
	for i := range grid {
		reversed[length-i-1] = grid[i]
	}
	return reversed
}
