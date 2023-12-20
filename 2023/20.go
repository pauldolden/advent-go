package _2023

import (
	"fmt"
	"strings"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
)

type broadcast struct {
	sender    string
	recipient string
}

type module struct {
	inputs []string
	mode   string
}

func TwentyOne(o config.Options) int {
	scanner, file := utils.OpenFile(2023, 20, o)
	defer file.Close()
	pulseMap := make(map[string]int)
	inputMap := make(map[string]module)
	var broadcasts [][]broadcast

	for scanner.Scan() {
		line := scanner.Text()
		parseInputs(line, inputMap)
		parseMode(line, inputMap)
		parsePulses(line, pulseMap)

		b := parseBroadcasts(line, pulseMap)

		broadcasts = append(broadcasts, b)
	}

	for _, b := range broadcasts {
		for _, r := range b {
			switch inputMap[r.recipient].mode {
			case "initial":
				fmt.Println(r.recipient)
			case "%":
				fmt.Println(r.recipient)
			case "&":
				fmt.Println(r.recipient)
			}
		}
	}

	return 0
}

func parseMode(s string, inputMap map[string]module) {
	ss := strings.Split(s, " -> ")
	sender := ss[0]
	var mode string

	if strings.Contains(sender, "%") || strings.Contains(sender, "&") {
		mode = sender[0:1]
		sender = sender[1:]
	} else {
		mode = "initial"
	}

	module := inputMap[sender]
	module.mode = mode
	inputMap[sender] = module
}

func parseBroadcasts(s string, pulseMap map[string]int) []broadcast {
	var broadcasts []broadcast
	ss := strings.Split(s, " -> ")
	sender := ss[0]

	if strings.Contains(sender, "%") || strings.Contains(sender, "&") {
		sender = sender[1:]
	}

	for _, r := range strings.Split(ss[1], ", ") {
		broadcasts = append(broadcasts, broadcast{
			sender:    sender,
			recipient: r,
		})
	}

	return broadcasts
}

func parsePulses(s string, pulseMap map[string]int) {
	ss := strings.Split(s, " -> ")
	sender := ss[0]

	if strings.Contains(sender, "%") || strings.Contains(sender, "&") {
		sender = sender[1:]
	}

	pulseMap[sender] = 0
}

func parseInputs(s string, inputMap map[string]module) {
	ss := strings.Split(s, " -> ")
	sender := ss[0]

	if strings.Contains(sender, "%") || strings.Contains(sender, "&") {
		sender = sender[1:]
	}
	for _, r := range strings.Split(ss[1], ", ") {
		if _, ok := inputMap[r]; !ok {
			inputMap[r] = module{
				inputs: []string{sender},
				mode:   "initial",
			}
		} else {
			inputMap[r] = module{
				inputs: append(inputMap[r].inputs, sender),
				mode:   "initial",
			}
		}
	}
}

func TwentyTwo(o config.Options) int {
	return 0
}
