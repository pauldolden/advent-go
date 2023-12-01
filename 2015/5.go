package _2015

import (
	"github.com/pauldolden/advent-go/utils"
)

var vowels = []string{"a", "e", "i", "o", "u"}
var badStrings = []string{"ab", "cd", "pq", "xy"}

func FiveOne() int {
	scanner, file := utils.OpenFile(2015, 5)
	defer file.Close()

	type conditions struct {
		vowels   int
		repeated int
		bad      int
	}

	var nice []string

	for scanner.Scan() {
		var c conditions
		line := scanner.Text()

		for _, char := range line {
			s := string(char)

			if stringInSlice(s, vowels) {
				c.vowels++
			}
		}

		for i := 0; i < len(line)-1; i++ {
			cur := string(line[i])
			next := string(line[i+1])

			if cur == next {
				c.repeated++
			}
			if stringInSlice(cur+next, badStrings) {
				c.bad++
			}

		}

		if c.vowels >= 3 && c.repeated > 0 && c.bad == 0 {
			nice = append(nice, line)
		}
	}

	return len(nice)
}

func FiveTwo() int {
	scanner, file := utils.OpenFile(2015, 5)
	defer file.Close()

	niceCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		var repeatedSkipped bool
		repeatedPairs := make(map[string][]int)

		// Check for repeating pairs and skipped repetition
		for i := 0; i < len(line)-1; i++ {
			pair := line[i : i+2]
			repeatedPairs[pair] = append(repeatedPairs[pair], i)

			if i < len(line)-2 && line[i] == line[i+2] {
				repeatedSkipped = true
			}
		}

		var hasPair bool
		for _, positions := range repeatedPairs {
			for i, pos := range positions {
				if i > 0 && pos >= positions[i-1]+2 {
					hasPair = true
					break
				}
			}
			if hasPair {
				break
			}
		}

		if repeatedSkipped && hasPair {
			niceCount++
		}
	}
	return niceCount
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
