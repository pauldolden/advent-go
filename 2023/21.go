package _2023

import (
	"slices"
	"strings"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
)

func TwentyOneOne(o config.Options) int {
	scanner, file := utils.OpenFile(2023, 21, o)
	defer file.Close()
	var grid [][]string
	var start utils.Node

	for scanner.Scan() {
		line := scanner.Text()

		grid = append(grid, strings.Split(line, ""))
	}

	for r, row := range grid {
		for c, col := range row {
			if col == "S" {
				start = utils.Node{Row: r, Col: c, Weight: 0}
			}
		}
	}

	return dijkstra21(grid, start)
}

func findAdjacentNodes21(grid [][]string, n utils.Node, seen [][2]int) []utils.Node {
	var nodes []utils.Node

	for _, dir := range utils.Directions {
		r := n.Row + dir[0]
		c := n.Col + dir[1]

		if utils.IsInBounds(r, c, grid) && grid[r][c] != "#" &&
			!slices.Contains(seen, [2]int{r, c}) {
			nodes = append(nodes, utils.Node{Row: r, Col: c, Weight: n.Weight + 1})
		}
	}

	return nodes
}

func dijkstra21(grid [][]string, start utils.Node) int {
	// create a priority queue
	var q utils.Queue[utils.Node]
	var ans []utils.Node
	var seen [][2]int

	q.Enqueue(start)

main:
	for len(q) > 0 {
		current := q.Dequeue()

		if current.Weight%2 == 0 {
			ans = append(ans, current)
		}

		if current.Weight == 26501365 {
			continue main
		}

		adjacentNodes := findAdjacentNodes21(grid, current, seen)

		for _, node := range adjacentNodes {
			seen = append(seen, [2]int{node.Row, node.Col})
			q.Enqueue(node)
		}
	}

	return len(ans) - 1
}

func TwentyOneTwo(o config.Options) int {
	return 0
}
