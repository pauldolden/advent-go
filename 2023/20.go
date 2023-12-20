package _2023

import (
	"strings"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
)

type broadcast struct {
	sender     string
	recipients []string
	// 0 = low, 1 = high
	pulse int
}

type module struct {
	name   string
	pulse  int
	inputs []module
}

type queue []broadcast

func (q *queue) enqueue(b broadcast) {
	*q = append(*q, b)
}

func (q *queue) dequeue() broadcast {
	b := (*q)[0]
	*q = (*q)[1:]
	return b
}

func TwentyOne(o config.Options) int {
	scanner, file := utils.OpenFile(2023, 20, o)
	defer file.Close()
	// high level overview we can snapshot this, so we can see when a cycle is in the initial state
	// state := make(map[string]int)
	depMap := make(map[string]module)

	// we just take the first line to get the broadcaster
	scanner.Scan()
	// broadcaster
	b := scanner.Text()
	broadcaster := parseBroadcast(b)
	q := queue{}
	q.enqueue(broadcaster)

	for scanner.Scan() {
		line := scanner.Text()
		b := parseBroadcast(line)
	}

	return 0
}

func parseBroadcast(s string, pulseMap) broadcast {
	ss := strings.Split(s, " -> ")
	sender := ss[0]
	recipients := strings.Split(ss[1], ", ")



	return broadcast{
		sender:     sender,
		recipients: recipients,
		pulse:      0,
	}
}

func TwentyTwo(o config.Options) int {
	return 0
}
