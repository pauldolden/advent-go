package _2015

import (
	"fmt"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
)

func EightOne(o config.Options) int {
	scanner, file := utils.OpenFile(2015, 8, o)
	defer file.Close()

	rawCharLength := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		rawCharLength = append(rawCharLength, len(line))
	}

	fmt.Println(rawCharLength)

	return 0
}

func EightTwo() int {
	return 0
}
