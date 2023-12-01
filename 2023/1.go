package _2023

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/pauldolden/advent-go/utils"
)

func OneOne() int {
	scanner, file := utils.OpenFile(2023, 1)
	defer file.Close()

	count := 0
	for scanner.Scan() {
		line := scanner.Text()

		number := findNumbers(line)

		count += number
	}

	return count
}

func OneTwo() int {
	scanner, file := utils.OpenFile(2023, 1)
	defer file.Close()

	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		number := findNumbersTwo(line)

		fmt.Println(number, line)

		count += number
	}

	return count
}

func findNumbers(s string) int {
	regex, err := regexp.Compile("[0-9]")

	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return 0
	}

	matches := regex.FindAllString(s, -1)

	sn := matches[0] + matches[len(matches)-1]

	i, _ := strconv.Atoi(sn)

	return i
}

func findNumbersTwo(s string) int {
	var i int

	matches := findAllOverlapping("one|two|three|four|five|six|seven|eight|nine|[0-9]", s)
	if len(matches) == 0 {
		return 0
	}

	first := transformMatch(matches[0])
	var last string

	if len(matches) > 1 {
		last = transformMatch(matches[len(matches)-1])
	} else {
		last = first
	}

	sn := first + last

	i, _ = strconv.Atoi(sn)

	return i
}

func transformMatch(s string) string {
	switch s {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	}

	return s
}

func findAllOverlapping(pattern, text string) []string {
	var results []string
	re := regexp.MustCompile(pattern)
	for i := 0; i < len(text); {
		match := re.FindStringIndex(text[i:])
		if match == nil {
			break
		}
		results = append(results, text[i+match[0]:i+match[1]])
		i += match[0] + 1 // Move forward by one rune
	}
	return results
}
