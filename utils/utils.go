package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/pauldolden/advent-go/config"
)

func OpenFile(year int, day int, o config.Options) (*bufio.Scanner, *os.File) {
	var path string
	if o.Test {
		if o.SplitInputs {
			path = fmt.Sprintf("./%d_test_%d.txt", day, o.TestPart)
		} else {
			path = fmt.Sprintf("./%d_test.txt", day)
		}
	} else {
		path = fmt.Sprintf("./%d/%d.txt", year, day)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	return scanner, file
}

type Node struct {
	Row, Col, Weight int
}

type PriorityQueue []Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Weight < pq[j].Weight }

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(n interface{}) { *pq = append(*pq, n.(Node)) }

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

// Dirs: up, right, down, left
var Directions = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func IsInBounds(r, c int, grid [][]string) bool {
	// returns true if the node is within the grid
	// if we are moving out of bounds, we want to copy the grid and expand it in the direction we are moving
	return false
}

type Queue[T any] []T

func (q *Queue[T]) Enqueue(b T) {
	*q = append(*q, b)
}

func (q *Queue[T]) Dequeue() T {
	b := (*q)[0]
	*q = (*q)[1:]
	return b
}
