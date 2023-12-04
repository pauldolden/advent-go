package _2023

import (
	"math"
	"regexp"
	"slices"
	"strings"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
)

type card struct {
	id             string
	winningNumbers []string
	gameNumbers    []string
	numberOfWins   int
	score          int
}

func (c *card) calculatePoints() {
	c.score = int(math.Pow(float64(2), float64(c.numberOfWins-1)))
}

func (c *card) calculateNumberOfWins() {
	for _, winningNumber := range c.winningNumbers {
		if slices.Contains(c.gameNumbers, winningNumber) {
			c.numberOfWins++
		}
	}
}

func (c *card) playOne() {
	c.calculateNumberOfWins()
	c.calculatePoints()
}

func (c *card) playTwo() {
	c.calculateNumberOfWins()
}

func FourOne(o config.Options) int {
	scanner, file := utils.OpenFile(2023, 4, o)
	defer file.Close()
	var score int

	for scanner.Scan() {
		line := scanner.Text()

		gameCard := parseLine(line)
		gameCard.playOne()

		score += gameCard.score
	}

	return score
}

func FourTwo(o config.Options) int {
	scanner, file := utils.OpenFile(2023, 4, o)
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()

		gameCard := parseLine(line)

		gameCard.playTwo()
	}
	return 0
}

func parseLine(s string) card {
	ss := strings.Split(s, "|")

	idRegex, _ := regexp.Compile("[0-9]+")
	gameBlock := strings.Split(ss[0], ":")

	id, winningNumbersString := idRegex.FindString(gameBlock[0]), gameBlock[1]
	gameNumbersString := ss[1]

	gameNumbersSlice := strings.Fields(gameNumbersString)
	winningNumbersSlice := strings.Fields(winningNumbersString)

	return card{
		id:             id,
		winningNumbers: winningNumbersSlice,
		gameNumbers:    gameNumbersSlice,
	}
}
