package _2023

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
)

type node struct {
	x, y, dweight, hweight, dir, dirMoves int
}

type priorityQueue []node

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].dweight < pq[j].dweight
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(n interface{}) {
	*pq = append(*pq, n.(node))
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

func SeventeenOne(o config.Options) int {
	scanner, file := utils.OpenFile(2023, 17, o)
	defer file.Close()

	grid := parse17(scanner)
	start := node{
		x:        0,
		y:        0,
		dweight:  0,
		hweight:  0,
		dir:      2,
		dirMoves: 0,
	}

	end := node{
		x:        len(grid) - 1,
		y:        len(grid[0]) - 1,
		dweight:  0,
		hweight:  0,
		dir:      2,
		dirMoves: 0,
	}

	return dijkstra(grid, start, end)
}

func manhattanDistanceFromTarget(x1, y1 int, grid [][]int) int {
	// returns the manhattan distance from the target bottom right corner
	targetX := len(grid) - 1
	targetY := len(grid[0]) - 1

	return int(math.Abs(float64(x1-targetX)) + math.Abs(float64(y1-targetY)))
}

func isInBounds(x, y int, grid [][]int) bool {
	// returns true if the node is within the grid
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}

func dijkstra(grid [][]int, start, end node) int {
	// Dirs: up, right, down, left
	queue := priorityQueue{}

	// starting point
	queue.Push(start)

	for queue.Len() > 0 {
		current := queue.Pop().(node)

		if current.x == end.x && current.y == end.y {
			return current.hweight
		}

		adjNodes := findAdjacentNodes(current, grid)
		fmt.Println("adjNodes", adjNodes)

		for _, adjNode := range adjNodes {
			queue.Push(adjNode)

			fmt.Println(adjNode)
		}
	}

	return 0
}

func findAdjacentNodes(n node, grid [][]int) []node {
	var nodes []node
	// Dirs: up, right, down, left
	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for d, dir := range dirs {
		adjX := n.x + dir[0]
		adjY := n.y + dir[1]

		if isInBounds(adjX, adjY, grid) && isValidDirection(n.dir, d, n.dirMoves) {
			// if we are going in the same direction then we can increase the dirMoves otherwise we reset it
			currentHWeight := n.hweight
			currentWeight := n.dweight
			cellValue := grid[adjX][adjY]
			mdist := manhattanDistanceFromTarget(adjX, adjY, grid)
			weight := currentWeight + cellValue + mdist
			moves := n.dirMoves
			hweight := currentHWeight + cellValue
			if n.dir == d {
				moves++
			} else {
				moves = 0
			}

			adjNode := node{
				x:        adjX,
				y:        adjY,
				dir:      d,
				dweight:  weight,
				hweight:  hweight,
				dirMoves: moves,
			}

			nodes = append(nodes, adjNode)
		}

	}

	return nodes
}

func isValidDirection(currentDir int, targetDir int, dirMoves int) bool {
	// returns false if th direction is backwards from the current direction
	if currentDir == 0 && targetDir == 2 {
		return false
	}

	if currentDir == 1 && targetDir == 3 {
		return false
	}

	if currentDir == 2 && targetDir == 0 {
		return false
	}

	if currentDir == 3 && targetDir == 1 {
		return false
	}

	// if dirs are the same we can only move in this direction again if we have moved in this direction less than 3 times
	if currentDir == targetDir && dirMoves >= 3 {
		return false
	}

	return true
}

func parse17(scanner *bufio.Scanner) [][]int {
	var iss [][]int
	for scanner.Scan() {
		line := scanner.Text()

		ss := strings.Split(line, "")
		var is []int
		for _, s := range ss {
			i, _ := strconv.Atoi(s)

			is = append(is, i)
		}
		iss = append(iss, is)
	}

	return iss
}
