package _2023

import (
	"bufio"
	"container/heap"
	"strconv"
	"strings"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
)

type node struct {
	r, c, weight, dir, dirMoves int
}

type priorityQueue []node

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool { return pq[i].weight < pq[j].weight }

func (pq priorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *priorityQueue) Push(n interface{}) { *pq = append(*pq, n.(node)) }

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
		r:        0,
		c:        0,
		weight:   0,
		dir:      2,
		dirMoves: 0,
	}

	end := node{
		r:        len(grid) - 1,
		c:        len(grid[0]) - 1,
		weight:   0,
		dir:      2,
		dirMoves: 0,
	}

	return dijkstra(grid, start, end)
}

func isInBounds(r, c int, grid [][]int) bool {
	// returns true if the node is within the grid
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}

func dijkstra(grid [][]int, start, end node) int {
	var queue priorityQueue
	heap.Init(&queue) // Initialize the priority queue as a heap
	seen := []node{}
	heap.Push(&queue, start)

	for queue.Len() > 0 {
		current := heap.Pop(&queue).(node)

		found := false
		for _, s := range seen {
			if s.r == current.r && s.c == current.c && s.dir == current.dir && s.dirMoves == current.dirMoves {
				found = true
				break
			}
		}

		if found {
			continue
		} else {
			seen = append(seen, current)
		}

		if current.r == end.r && current.c == end.c {
			return current.weight
		}

		adjNodes := findAdjacentNodes(current, grid)

		for _, adjNode := range adjNodes {
			heap.Push(&queue, adjNode)
		}
	}

	return 0
}

func findAdjacentNodes(n node, grid [][]int) []node {
	var nodes []node
	// Dirs: up, right, down, left
	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for i, dir := range dirs {
		if isInBounds(n.r+dir[0], n.c+dir[1], grid) {
			if i == n.dir {
				if n.dirMoves < 3 {
					cellValue := grid[n.r+dir[0]][n.c+dir[1]]
					nodes = append(nodes, node{
						r:        n.r + dir[0],
						c:        n.c + dir[1],
						weight:   n.weight + cellValue,
						dir:      i,
						dirMoves: n.dirMoves + 1,
					})
				}
			} else if isValidDirection(n.dir, i, n.dirMoves) {
				cellValue := grid[n.r+dir[0]][n.c+dir[1]]
				nodes = append(nodes, node{
					r:        n.r + dir[0],
					c:        n.c + dir[1],
					weight:   n.weight + cellValue,
					dir:      i,
					dirMoves: 1,
				})
			}
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
