package _2023

import (
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
)

var CARDS_ONE = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
var CARDS_TWO = []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}

type hand struct {
	rawHand                string
	cardScores             []int
	maxInstanceCount       int
	secondMaxInstanceCount int
	bid                    int
}

type (
	ByMaxInstanceCount       []hand
	BySecondMaxInstanceCount []hand
	ByCardScores             []hand
)

type ByPriority []hand

func (a ByPriority) Len() int      { return len(a) }
func (a ByPriority) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByPriority) Less(i, j int) bool {
	// First, sort by MaxInstanceCount
	if a[i].maxInstanceCount != a[j].maxInstanceCount {
		return a[i].maxInstanceCount < a[j].maxInstanceCount
	}
	// Next, sort by SecondMaxInstanceCount if MaxInstanceCount is equal
	if a[i].secondMaxInstanceCount != a[j].secondMaxInstanceCount {
		return a[i].secondMaxInstanceCount < a[j].secondMaxInstanceCount
	}
	// Finally, sort by CardScores if both the above are equal
	for idx := range a[i].cardScores {
		if a[i].cardScores[idx] != a[j].cardScores[idx] {
			return a[i].cardScores[idx] < a[j].cardScores[idx]
		}
	}
	return false // If all criteria are equal
}

func SevenOne(o config.Options) int {
	scanner, file := utils.OpenFile(2023, 7, o)
	defer file.Close()

	var hands []hand

	for scanner.Scan() {
		line := scanner.Text()
		var hand hand

		instanceCount := make(map[string]int)
		cardsScores := []int{}
		input := strings.Fields(line)

		cards, bid := input[0], input[1]

		bidInt, _ := strconv.Atoi(bid)

		for _, card := range cards {
			idx := slices.Index(CARDS_ONE, string(card))
			cardsScores = append(cardsScores, idx)
			instanceCount[string(card)]++
		}

		for _, count := range instanceCount {
			if count > hand.maxInstanceCount {
				hand.secondMaxInstanceCount = hand.maxInstanceCount
				hand.maxInstanceCount = count
			} else if count > hand.secondMaxInstanceCount && count <= hand.maxInstanceCount {
				hand.secondMaxInstanceCount = count
			}
		}

		hand.rawHand = cards
		hand.bid = bidInt
		hand.cardScores = cardsScores

		hands = append(hands, hand)
	}

	sort.Sort(ByPriority(hands))

	var total int
	for idx, hand := range hands {
		total += hand.bid * (idx + 1)
	}

	return total
}

func SevenTwo(o config.Options) int {
	scanner, file := utils.OpenFile(2023, 7, o)
	defer file.Close()
	var hands []hand

	for scanner.Scan() {
		line := scanner.Text()
		var hand hand
		instanceCount := make(map[string]int)
		cardsScores := []int{}
		input := strings.Fields(line)
		cards, bid := input[0], input[1]
		bidInt, _ := strconv.Atoi(bid)

		for _, card := range cards {
			idx := slices.Index(CARDS_TWO, string(card))
			cardsScores = append(cardsScores, idx)
			instanceCount[string(card)]++
		}

		numberOfJokers := instanceCount["J"]
		// remove jokers from instanceCount
		delete(instanceCount, "J")

		for _, count := range instanceCount {
			if count > hand.maxInstanceCount {
				hand.secondMaxInstanceCount = hand.maxInstanceCount
				hand.maxInstanceCount = count
			} else if count > hand.secondMaxInstanceCount && count <= hand.maxInstanceCount {
				hand.secondMaxInstanceCount = count
			}
		}

		// Add jokers to maxInstanceCount
		if numberOfJokers > 0 {
			hand.maxInstanceCount += numberOfJokers
		}

		hand.rawHand = cards
		hand.bid = bidInt
		hand.cardScores = cardsScores
		hands = append(hands, hand)
	}

	sort.Sort(ByPriority(hands))

	var total int
	for idx, hand := range hands {
		total += hand.bid * (idx + 1)
	}

	return total
}
