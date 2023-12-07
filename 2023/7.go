package _2023

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
)

var CARDS = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

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
			// find index of card in CARDS
			idx := slices.Index(CARDS, string(card))
			// add score to cardsScores
			cardsScores = append(cardsScores, idx)
			instanceCount[string(card)]++
		}

		for _, count := range instanceCount {
			if count > hand.maxInstanceCount {
				hand.secondMaxInstanceCount = hand.maxInstanceCount // Update second max before updating max
				hand.maxInstanceCount = count
			} else if count > hand.secondMaxInstanceCount && count < hand.maxInstanceCount {
				hand.secondMaxInstanceCount = count
			}
		}

		hand.rawHand = cards
		hand.bid = bidInt
		hand.cardScores = cardsScores

		hands = append(hands, hand)
	}

	sort.Sort(ByPriority(hands))

	total := 0
	for idx, hand := range hands {
		fmt.Println("======")
		fmt.Println(hand)
		fmt.Println(idx+1, "x", hand.bid, "=", hand.bid*(idx+1))
		total += hand.bid * (idx + 1)
	}

	return total
}

func SevenTwo(o config.Options) int {
	return 0
}
