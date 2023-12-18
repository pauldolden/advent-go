package _2023

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
)

type node struct {
	x, y, weight, dir, dirMoves int
}

type priorityQueue []node

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
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

	for _, row := range grid {
		fmt.Println(row)
	}

	return 0
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
