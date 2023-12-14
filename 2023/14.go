package _2023

import (
	"strings"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
)

func FourteenOne(o config.Options) int {
	scanner, file := utils.OpenFile(2023, 14, o)
	defer file.Close()

	var grid [][]string

	for scanner.Scan() {
		s := scanner.Text()

		grid = append(grid, strings.Split(s, ""))
	}

	moveUp(grid)

	total := 0
	for i, row := range grid {
		factor := len(grid) - i
		for _, col := range row {
			if col == "O" {
				total += factor * 1
			}
		}
	}

	return total
}

func FourteenTwo(o config.Options) int {
	scanner, file := utils.OpenFile(2023, 14, o)
	defer file.Close()

	var grid [][]string
	for scanner.Scan() {
		s := scanner.Text()
		grid = append(grid, strings.Split(s, ""))
	}

	directions := []string{"north", "east", "south", "west"}
	for i := 0; i < 1000; i++ {
		for _, dir := range directions {
			switch dir {
			case "north":
				grid = moveUp(grid)
			case "west":
				grid = rotateCounterClockwise(grid)
				grid = moveUp(grid)
				grid = rotateClockwise(grid)
			case "south":
				grid = rotateClockwise(grid)
				grid = rotateClockwise(grid)
				grid = moveUp(grid)
				grid = rotateCounterClockwise(grid)
				grid = rotateCounterClockwise(grid)
			case "east":
				grid = rotateClockwise(grid)
				grid = moveUp(grid)
				grid = rotateCounterClockwise(grid)
			}
		}
	}

	total := 0
	for i, row := range grid {
		factor := len(grid) - i
		for _, col := range row {
			if col == "O" {
				total += factor * 1
			}
		}
	}

	return total
}

func rotateClockwise(grid [][]string) [][]string {
	n := len(grid)
	rotated := make([][]string, n)
	for i := 0; i < n; i++ {
		rotated[i] = make([]string, n)
		for j := 0; j < n; j++ {
			rotated[i][j] = grid[n-j-1][i]
		}
	}
	return rotated
}

func rotateCounterClockwise(grid [][]string) [][]string {
	n := len(grid)
	rotated := make([][]string, n)
	for i := 0; i < n; i++ {
		rotated[i] = make([]string, n)
		for j := 0; j < n; j++ {
			rotated[i][j] = grid[j][n-i-1]
		}
	}
	return rotated
}

func moveUp(grid [][]string) [][]string {
	for i, row := range grid {
		if i == 0 {
			continue
		}

		for j, col := range row {
			if col == "." || col == "#" {
				continue
			}

			if grid[i-1][j] == "." {
				for k := i - 1; k >= 0; k-- {
					if grid[k][j] == "." {
						grid[k][j] = "O"
						grid[k+1][j] = "."
					} else {
						break
					}
				}
			}
		}
	}
	return grid
}
