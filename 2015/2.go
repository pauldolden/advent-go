package _2015

import (
	"log"
	"strings"

	"strconv"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
)

func TwoOne(o config.Options) int {
	var total int
	scanner, file := utils.OpenFile(2015, 2, o)
	defer file.Close()

	for scanner.Scan() {
		l := scanner.Text()

		ss := strings.Split(l, "x")

		is, err := stringSliceToIntSlice(ss)

		if err != nil {
			log.Fatal(err)
		}

		lw := 2 * (is[0] * is[1])
		wh := 2 * (is[1] * is[2])
		lh := 2 * (is[0] * is[2])

		sm := min(lw, wh, lh) / 2

		total += (lw + wh + lh + sm)
	}

	return total
}

func TwoTwo(o config.Options) int {
	var total int
	scanner, file := utils.OpenFile(2015, 2, o)
	defer file.Close()

	for scanner.Scan() {
		l := scanner.Text()

		ss := strings.Split(l, "x")

		is, err := stringSliceToIntSlice(ss)

		if err != nil {
			log.Fatal(err)
		}

		lg := is[0]
		w := is[1]
		h := is[2]

		rb := lg * w * h

		sm, ssm := twoSmallestNumbers(lg, w, h)

		mrb := sm + sm + ssm + ssm

		total += rb + mrb
	}

	return total
}

func stringSliceToIntSlice(stringSlice []string) ([]int, error) {
	var intSlice []int
	for _, str := range stringSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			// Handle the error if the string cannot be converted to an integer
			return nil, err
		}
		intSlice = append(intSlice, num)
	}
	return intSlice, nil
}

func twoSmallestNumbers(a, b, c int) (int, int) {
	// Initially assume a and b are the smallest and second smallest
	smallest, secondSmallest := min(a, b), max(a, b)

	// Now compare with c
	if c < smallest {
		smallest, secondSmallest = c, smallest
	} else if c < secondSmallest {
		secondSmallest = c
	}

	return smallest, secondSmallest
}
