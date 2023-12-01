package _2015

import (
	"github.com/pauldolden/advent-go/utils"
)

func OneOne() int {
	scanner, file := utils.OpenFile(2015, 1)
	defer file.Close()

	var floor int

	for scanner.Scan() {
		for _, char := range scanner.Text() {
			if char == '(' {
				floor++
			} else if char == ')' {
				floor--
			}
		}
	}

	return floor
}

func OneTwo() int {
	scanner, file := utils.OpenFile(2015, 1)
	defer file.Close()

	var floor int
	var position int

	for scanner.Scan() {
		for _, char := range scanner.Text() {
			position++
			if char == '(' {
				floor++
			} else if char == ')' {
				floor--
			}

			if floor == -1 {
				return position
			}
		}
	}

	return 0
}
