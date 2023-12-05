package _2023

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/pauldolden/advent-go/config"
	"github.com/samber/lo"
)

type resourceMap struct {
	sourceName string
	destName   string
	resources  []resource
}

type resource struct {
	sourceStart int
	destStart   int
	length      int
}

func FiveOne(o config.Options) int {
	input, err := os.ReadFile("./2023/5.txt")
	if err != nil {
		log.Fatal(err)
	}

	seeds, maps := parseFileOne(input)
	lowestLocation := math.MaxInt

	for _, seed := range seeds {
		currentLocation := findLocation(seed, maps, 0)
		if currentLocation < lowestLocation {
			lowestLocation = currentLocation
		}
	}

	return lowestLocation
}

func FiveTwo(o config.Options) int {
	input, err := os.ReadFile("./2023/5.txt")
	if err != nil {
		log.Fatal(err)
	}

	seeds, maps := parseFileTwo(input)
	lowestLocation := math.MaxInt

	for _, seed := range seeds {
		currentLocation := findLocation(seed, maps, 0)
		if currentLocation < lowestLocation {
			lowestLocation = currentLocation
		}
	}

	return lowestLocation
}

func findLocation(v int, rm []resourceMap, index int) int {
	if index == len(rm) {
		return v
	}

	m := rm[index]

	for _, r := range m.resources {
		if v >= r.sourceStart && v < r.sourceStart+r.length {
			v = r.destStart + (v - r.sourceStart)
			break
		}
	}

	return findLocation(v, rm, index+1)
}

func parseFileOne(input []byte) ([]int, []resourceMap) {
	var maps []resourceMap
	sections := strings.Split(string(input), "\n\n")

	seeds := parseSeedsOne(sections[0])

	resourceMaps := sections[1:]

	for _, m := range resourceMaps {
		rm := parseMap(m)

		maps = append(maps, rm)
	}

	return seeds, maps
}

func parseFileTwo(input []byte) ([]int, []resourceMap) {
	var maps []resourceMap
	sections := strings.Split(string(input), "\n\n")

	seeds := parseSeedsTwo(sections[0])

	resourceMaps := sections[1:]

	for _, m := range resourceMaps {
		rm := parseMap(m)

		maps = append(maps, rm)
	}

	return seeds, maps
}

func parseSeedsOne(s string) []int {
	var seeds []int

	seedsStringSlice := strings.Fields(strings.Split(s, ":")[1])

	for _, seed := range seedsStringSlice {
		seedInt, err := strconv.Atoi(seed)

		if err != nil {
			log.Fatal(err)
		}
		seeds = append(seeds, seedInt)
	}

	return seeds
}

func parseSeedsTwo(s string) []int {
	var rawSeeds []int
	var seeds []int

	seedsStringSlice := strings.Fields(strings.Split(s, ":")[1])

	for _, seed := range seedsStringSlice {
		seedInt, err := strconv.Atoi(seed)

		if err != nil {
			log.Fatal(err)
		}
		rawSeeds = append(rawSeeds, seedInt)
	}

	seedsChunks := lo.Chunk(rawSeeds, 2)

	for _, chunk := range seedsChunks {
		for i := 0; i < chunk[1]; i++ {
			seeds = append(seeds, chunk[0]+i)
		}
	}

	return seeds
}

func parseMap(m string) resourceMap {
	var k resourceMap

	ms := strings.Split(strings.TrimSpace(m), "\n")

	keyParts := strings.Split(ms[0], "-to-")

	k.sourceName = keyParts[0]
	k.destName = strings.Fields(keyParts[1])[0]

	for _, r := range ms[1:] {
		rp := parseResource(r)

		k.resources = append(k.resources, rp)
	}

	return k
}

func parseResource(r string) resource {
	var rp resource

	rs := strings.Fields(r)

	rp.sourceStart, _ = strconv.Atoi(rs[1])
	rp.destStart, _ = strconv.Atoi(rs[0])
	rp.length, _ = strconv.Atoi(rs[2])

	return rp
}
