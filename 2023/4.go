package _2023

import (
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
)

type card struct {
	id             int
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

func cloneCards(c *card, m map[int]card) []card {
	var cards []card
	c.calculateNumberOfWins()
	cards = append(cards, *c)

	for i := 1; i <= c.numberOfWins; i++ {
		newCard := m[c.id+i]
		newCardCards := cloneCards(&newCard, m)

		cards = append(cards, newCardCards...)
	}

	return cards
}

func (c *card) playOne() {
	c.calculateNumberOfWins()
	c.calculatePoints()
}

func (c *card) playTwo(m map[int]card) []card {
	return cloneCards(c, m)
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
	mapOfCards := make(map[int]card)
	finalCards := []card{}

	for scanner.Scan() {
		line := scanner.Text()

		gameCard := parseLine(line)

		mapOfCards[gameCard.id] = gameCard
	}

	for _, c := range mapOfCards {
		cards := c.playTwo(mapOfCards)

		finalCards = append(finalCards, cards...)
	}

	return len(finalCards)
}

func parseLine(s string) card {
	ss := strings.Split(s, "|")

	idRegex, _ := regexp.Compile("[0-9]+")
	gameBlock := strings.Split(ss[0], ":")

	idString, winningNumbersString := idRegex.FindString(gameBlock[0]), gameBlock[1]
	gameNumbersString := ss[1]

	id, _ := strconv.Atoi(idString)

	gameNumbersSlice := strings.Fields(gameNumbersString)
	winningNumbersSlice := strings.Fields(winningNumbersString)

	return card{
		id:             id,
		winningNumbers: winningNumbersSlice,
		gameNumbers:    gameNumbersSlice,
	}
}
