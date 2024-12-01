package _2024

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
)

func OneOne(o config.Options) int {
	scanner, file := utils.OpenFile(2024, 1, o)
	defer file.Close()
	var left []int
	var right []int

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		ss := strings.Fields(line)
		ls, rs := ss[0], ss[1]
		li, _ := strconv.Atoi(ls)
		ri, _ := strconv.Atoi(rs)

		left = append(left, li)
		right = append(right, ri)
	}
	slices.Sort(left)
	slices.Sort(right)

	for i := range len(left) {
		diff := math.Abs(float64(left[i] - right[i]))

		count += int(diff)
	}

	return count
}

func OneTwo(o config.Options) int {
	scanner, file := utils.OpenFile(2024, 1, o)
	defer file.Close()
	var left []int
	var right []int

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		ss := strings.Fields(line)
		ls, rs := ss[0], ss[1]
		li, _ := strconv.Atoi(ls)
		ri, _ := strconv.Atoi(rs)

		left = append(left, li)
		right = append(right, ri)
	}
	slices.Sort(left)
	slices.Sort(right)

	for i := range len(left) {
		multiplier := 0
		for j := range len(right) {
			if right[j] > left[i] {
				break
			}
			if right[j] == left[i] {
				multiplier++
			}
		}
		count += left[i] * multiplier
	}

	return count
}
