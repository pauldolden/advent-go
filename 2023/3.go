package _2023

import (
	"regexp"
	"strconv"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
)

type xLocation struct {
	start int
	end   int
}

type partNumber struct {
	value int
	x     xLocation
	y     int
}

func ThreeOne(o config.Options) int {
	scanner, file := utils.OpenFile(2023, 3, o)
	defer file.Close()

	numberRegex := regexp.MustCompile("[0-9]+")
	symbolRegex := regexp.MustCompile(`[^\w.]`)

	symbolMap := make(map[[2]int]string)
	partNumbers := []partNumber{}

	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()

		numbers := numberRegex.FindAllIndex([]byte(line), -1)
		symbols := symbolRegex.FindAllIndex([]byte(line), -1)

		for _, sym := range symbols {
			s := sym[0]

			symbolMap[[2]int{lineNumber, s}] = string(line[s])
		}

		for _, number := range numbers {
			s, e := number[0], number[len(number)-1]

			ns := line[s:e]

			n, _ := strconv.Atoi(ns)

			pn := partNumber{
				value: n,
				y:     lineNumber,
				x: xLocation{
					start: s,
					end:   e,
				},
			}

			partNumbers = append(partNumbers, pn)
		}

		lineNumber++
	}

	count := 0

main:
	for _, num := range partNumbers {
		beforeKey := [2]int{num.y, num.x.start - 1}
		afterKey := [2]int{num.y, num.x.end}

		// Matches symbols immediately before or after
		if _, ok := symbolMap[beforeKey]; ok {
			count += num.value
			continue main
		}
		if _, ok := symbolMap[afterKey]; ok {
			count += num.value
			continue main
		}

		// Matches line above
		for i := num.x.start - 1; i <= num.x.end; i++ {
			key := [2]int{num.y - 1, i}
			if _, ok := symbolMap[key]; ok {
				count += num.value
				continue main
			}
		}

		// Matches line below
		for i := num.x.start - 1; i <= num.x.end; i++ {
			key := [2]int{num.y + 1, i}
			if _, ok := symbolMap[key]; ok {
				count += num.value
				continue main
			}
		}
	}

	return count
}

func ThreeTwo(o config.Options) int {
	scanner, file := utils.OpenFile(2023, 3, o)
	defer file.Close()

	numberRegex := regexp.MustCompile("[0-9]+")
	symbolRegex := regexp.MustCompile(`[^\w.]`)

	symbolMap := make(map[[2]int]string)
	partNumbers := []partNumber{}

	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()

		numbers := numberRegex.FindAllIndex([]byte(line), -1)
		symbols := symbolRegex.FindAllIndex([]byte(line), -1)

		for _, sym := range symbols {
			s := sym[0]

			if string(line[s]) == "*" {

				symbolMap[[2]int{lineNumber, s}] = string(line[s])
			}
		}

		for _, number := range numbers {
			s, e := number[0], number[len(number)-1]

			ns := line[s:e]

			n, _ := strconv.Atoi(ns)

			pn := partNumber{
				value: n,
				y:     lineNumber,
				x: xLocation{
					start: s,
					end:   e,
				},
			}

			partNumbers = append(partNumbers, pn)
		}

		lineNumber++
	}

	gears := make(map[[2]int][]int)
	count := 0
main:
	for _, num := range partNumbers {
		beforeKey := [2]int{num.y, num.x.start - 1}
		afterKey := [2]int{num.y, num.x.end}

		// Matches symbols immediately before or after
		if _, ok := symbolMap[beforeKey]; ok {
			if symbolMap[beforeKey] == "*" {
				gears[beforeKey] = append(gears[beforeKey], num.value)
			}
			continue main
		}
		if _, ok := symbolMap[afterKey]; ok {
			if symbolMap[afterKey] == "*" {
				gears[afterKey] = append(gears[afterKey], num.value)
			}
			continue main
		}
		// Matches line above
		for i := num.x.start - 1; i <= num.x.end; i++ {
			key := [2]int{num.y - 1, i}
			if _, ok := symbolMap[key]; ok {
				if symbolMap[key] == "*" {
					gears[key] = append(gears[key], num.value)
				}
				continue main
			}
		}
		// Matches line below
		for i := num.x.start - 1; i <= num.x.end; i++ {
			key := [2]int{num.y + 1, i}
			if _, ok := symbolMap[key]; ok {
				if symbolMap[key] == "*" {
					gears[key] = append(gears[key], num.value)
				} else {
					count += num.value
				}
				continue main
			}
		}
	}

	gearCount := 0
	for _, gear := range gears {
		if len(gear) == 2 {
			gearCount += gear[0] * gear[1]
		}
	}

	return gearCount
}
