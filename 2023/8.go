package _2023

import (
	"log"
	"os"
	"strings"

	"github.com/pauldolden/advent-go/config"
)

var dirMap = map[string]int{
	"L": 0,
	"R": 1,
}

func EightOne(o config.Options) int {
	input, err := os.ReadFile("./2023/8.txt")
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(string(input), "\n\n")

	rawSteps := parts[0]
	steps := strings.Split(rawSteps, "")
	mappings := parts[1]
	mappingSlice := strings.Split(mappings, "\n")
	mappingsMap := make(map[string][]string)

	for _, mapping := range mappingSlice {
		if mapping == "" {
			continue
		}
		source, opts := parseLine7(mapping)
		mappingsMap[source] = opts
	}

	found := false

	count := 0
	for !found {
		found = walk(0, "AAA", steps, mappingsMap, &count)
	}

	return count
}

func EightTwo(o config.Options) int {
	input, err := os.ReadFile("./2023/8.txt")
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(string(input), "\n\n")

	rawSteps := parts[0]
	steps := strings.Split(rawSteps, "")
	mappings := parts[1]
	mappingSlice := strings.Split(mappings, "\n")
	mappingsMap := make(map[string][]string)

	for _, mapping := range mappingSlice {
		if mapping == "" {
			continue
		}
		source, opts := parseLine7(mapping)
		mappingsMap[source] = opts
	}

	var sources []string
	for key := range mappingsMap {
		if string(key[2]) == "A" {
			sources = append(sources, key)
		}
	}

	count := walk2(sources, steps, mappingsMap)
	return count
}

func walk(
	idx int,
	source string,
	steps []string,
	mappingsMap map[string][]string,
	count *int,
) bool {
	if idx >= len(steps) {
		idx = 0
	}

	if source == "ZZZ" {
		return true
	}

	*count++

	dirIdx := dirMap[steps[idx]]
	newSource := mappingsMap[source][dirIdx]

	return walk(idx+1, newSource, steps, mappingsMap, count)
}

func walk2(sources []string, steps []string, mappingsMap map[string][]string) int {
	idx := 0
	count := 0

	for {
		if idx >= len(steps) {
			idx = 0
		}

		if allEndWithZ(sources) {
			break
		}

		var newSources []string
		count++

		dirIdx := dirMap[steps[idx]]
		for _, source := range sources {
			newSource := mappingsMap[source][dirIdx]
			newSources = append(newSources, newSource)
		}

		sources = newSources
		idx++
	}

	return count
}

func parseLine7(input string) (string, []string) {
	var source string
	var opts []string

	parts := strings.Split(input, " = ")
	source = parts[0]
	opts = strings.Split(strings.Trim(parts[1], "()"), ", ")

	return source, opts
}

func allEndWithZ(sources []string) bool {
	for _, source := range sources {
		if len(source) == 0 || source[len(source)-1] != 'Z' {
			return false
		}
	}
	return true
}
