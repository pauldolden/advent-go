package _2023

import (
	"log"
	"strconv"
	"strings"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
)

func SixOne(o config.Options) int {
	scanner, file := utils.OpenFile(2023, 6, o)
	defer file.Close()
	var results []int
	game := make(map[string][]int)

	for scanner.Scan() {
		line := scanner.Text()

		makeGameOne(line, game)
	}

	for i := 0; i < len(game["Time"]); i++ {
		results = append(results, playGame(i, game))
	}

	var total int = 1
	for _, result := range results {
		if result > 0 {
			total *= result
		}
	}

	return total
}

func SixTwo(o config.Options) int {
	scanner, file := utils.OpenFile(2023, 6, o)
	defer file.Close()
	game := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()

		makeGameTwo(line, game)
	}

	res := playGameTwo(game)

	return res
}

func playGame(idx int, game map[string][]int) int {
	var wins int
	time := game["Time"][idx]
	targetDistance := game["Distance"][idx]

	for i := 0; i < time; i++ {
		diff := time - i

		if diff*i > targetDistance {
			wins++
		}
	}

	return wins
}

func makeGameOne(line string, m map[string][]int) {
	var values []int

	input := strings.Fields(line)

	key := input[0][:len(input[0])-1]

	for _, value := range input[1:] {
		i, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}

		values = append(values, i)
	}

	m[key] = values
}

func makeGameTwo(line string, m map[string]int) {
	input := strings.Split(line, ":")

	key := input[0]
	value := strings.ReplaceAll(input[1], " ", "")

	i, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}

	m[key] = i
}

func playGameTwo(game map[string]int) int {
	var wins int
	time := game["Time"]
	targetDistance := game["Distance"]

	for i := 0; i < time; i++ {
		diff := time - i

		if diff*i > targetDistance {
			wins++
		}
	}

	return wins
}
