package _2023

import (
	"bufio"
	"fmt"
	"slices"
	"strings"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
)

func SixteenOne(o config.Options) int {
	scanner, file := utils.OpenFile(2023, 16, o)
	defer file.Close()

	grid := parse16(scanner)

	return travel(grid, [3]int{0, 0, 1})
}

func SixteenTwo(o config.Options) int {
	scanner, file := utils.OpenFile(2023, 16, o)
	defer file.Close()

	grid := parse16(scanner)

	// we are going to start from every edge or if we are a corner then we will start from the corner and go both ways
	// we will then count the number of tiles we have seen and return the biggest number

	var numbers []int
	// top and bottom
	for i := 0; i < len(grid); i++ {
		numbers = append(numbers, travel(grid, [3]int{i, 0, 1}))
		numbers = append(numbers, travel(grid, [3]int{i, len(grid[0]) - 1, 3}))
	}

	// left and right
	for i := 0; i < len(grid[0]); i++ {
		numbers = append(numbers, travel(grid, [3]int{0, i, 2}))
		numbers = append(numbers, travel(grid, [3]int{len(grid) - 1, i, 0}))
	}

	return slices.Max(numbers)
}

func travel(grid [][]string, start [3]int) int {
	// Dirs: up, right, down, left
	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	queue := make([][3]int, 0)

	// starting point
	queue = append(queue, start)
	// empowered tiles
	seen := [][3]int{}

	for len(queue) > 0 {
		current, remaining := dequeue(queue)

		dir := current[2]

		// if we are within grid bounds then add the current tile to the seen list
		if current[0] >= 0 && current[0] < len(grid) &&
			current[1] >= 0 && current[1] < len(grid[0]) {

			tile := grid[current[0]][current[1]]

			// mark the tile as seen
			// if it's not already in seen then add it
			found := false
			for _, s := range seen {
				if s[0] == current[0] && s[1] == current[1] && s[2] == dir {
					found = true
				}
			}
			if !found {
				seen = append(seen, [3]int{current[0], current[1], dir})
			} else {
				queue = remaining
				continue
			}

			switch tile {
			case `/`:
				switch dir {
				// if we are going right we need to go up
				case 1:
					remaining = append(
						remaining,
						[3]int{current[0] + dirs[0][0], current[1] + dirs[0][1], 0},
					)
					// if we are going down we need to go left
				case 2:
					remaining = append(
						remaining,
						[3]int{current[0] + dirs[3][0], current[1] + dirs[3][1], 3},
					)
					// if we are going left we need to go down
				case 3:
					remaining = append(
						remaining,
						[3]int{current[0] + dirs[2][0], current[1] + dirs[2][1], 2},
					)
					// if we are going up we need to go right
				case 0:
					remaining = append(
						remaining,
						[3]int{current[0] + dirs[1][0], current[1] + dirs[1][1], 1},
					)
				}
			case `\`:
				switch dir {
				// if we are going right we need to go down
				case 1:
					remaining = append(
						remaining,
						[3]int{current[0] + dirs[2][0], current[1] + dirs[2][1], 2},
					)
					// if we are going down we need to go right
				case 2:
					remaining = append(
						remaining,
						[3]int{current[0] + dirs[1][0], current[1] + dirs[1][1], 1},
					)
					// if we are going left we need to go up
				case 3:
					remaining = append(
						remaining,
						[3]int{current[0] + dirs[0][0], current[1] + dirs[0][1], 0},
					)
					// if we are going up we need to go left
				case 0:
					remaining = append(
						remaining,
						[3]int{current[0] + dirs[3][0], current[1] + dirs[3][1], 3},
					)
				}
			case ".":
				// continue in the same direction
				remaining = append(
					remaining,
					[3]int{current[0] + dirs[dir][0], current[1] + dirs[dir][1], dir},
				)
			case "-":

				// if we are going left or right then we can continue in the same direction
				switch dir {
				case 1, 3:
					// continue in the same direction
					remaining = append(
						remaining,
						[3]int{current[0] + dirs[dir][0], current[1] + dirs[dir][1], dir},
					)
				case 0, 2:
					// if we are going up or down then we need to split and go both ways left and right
					remaining = append(
						remaining,
						[3]int{current[0] + dirs[1][0], current[1] + dirs[1][1], 1},
						[3]int{current[0] + dirs[3][0], current[1] + dirs[3][1], 3},
					)
				}

			case "|":
				switch dir {
				// if we are going up or down then we can continue in the same direction
				case 0, 2:
					// continue in the same direction
					remaining = append(
						remaining,
						[3]int{current[0] + dirs[dir][0], current[1] + dirs[dir][1], dir},
					)
				case 1, 3:
					remaining = append(
						remaining,
						[3]int{current[0] + dirs[0][0], current[1] + dirs[0][1], 0},
						[3]int{current[0] + dirs[2][0], current[1] + dirs[2][1], 2},
					)
				}
			default:
				panic(fmt.Sprintf("Unknown tile: %s", tile))
			}

			queue = remaining

		} else {
			queue = remaining
			continue
		}
	}

	count := 0
	deduped := [][2]int{}
	for _, s := range seen {
		found := false
		for _, d := range deduped {
			if d[0] == s[0] && d[1] == s[1] {
				found = true
			}
		}
		if !found {
			deduped = append(deduped, [2]int{s[0], s[1]})
			count++
		}
	}

	return count
}

func dequeue(queue [][3]int) ([3]int, [][3]int) {
	next := queue[0]

	queue = queue[1:]

	return next, queue
}

func parse16(scanner *bufio.Scanner) [][]string {
	var sss [][]string
	for scanner.Scan() {
		line := scanner.Text()

		ss := strings.Split(line, "")

		sss = append(sss, ss)
	}

	return sss
}
