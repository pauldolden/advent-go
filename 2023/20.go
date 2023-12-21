package _2023

import (
	"strings"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
)

type broadcast struct {
	sender    string
	recipient string
}

type module struct {
	inputs     []input
	recipients []string
	mode       string
}

type input struct {
	name      string
	lastPulse int
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
	pulseMap := make(map[string]int)
	inputMap := make(map[string]module)
	var initialBroadcasts []broadcast
	q := queue{}

	for scanner.Scan() {
		line := scanner.Text()
		parseInputs(line, inputMap)
		parseMode(line, inputMap)
		parsePulses(line, pulseMap)
		parseRecipients(line, inputMap)
		b := parseBroadcast(line)
		initialBroadcasts = append(initialBroadcasts, b...)
	}

	lowPulseCount := 0
	highPulseCount := 0

	for i := 0; i < 1000; i++ {
		for _, b := range initialBroadcasts {
			q.enqueue(b)
		}
		lowPulseCount += 1 // the button always starts with a low pulse
		for len(q) > 0 {
			bc := q.dequeue()
			sPulse := pulseMap[bc.sender]
			recipient := inputMap[bc.recipient]

			if sPulse == 0 {
				lowPulseCount += 1
			} else {
				highPulseCount += 1
			}

			switch recipient.mode {
			case "%":
				if sPulse == 1 {
					continue
				} else {
					if pulseMap[bc.recipient] == 0 {
						pulseMap[bc.recipient] = 1
					} else {
						pulseMap[bc.recipient] = 0
					}
				}
			case "&":
				for i, input := range recipient.inputs {
					if input.name == bc.sender {
						recipient.inputs[i].lastPulse = sPulse
					}

					allHigh := true
					for _, input := range recipient.inputs {
						if input.lastPulse == 0 {
							allHigh = false
						}
					}
					if allHigh {
						pulseMap[bc.recipient] = 0
					} else {
						pulseMap[bc.recipient] = 1
					}
				}
			}

			for _, r := range recipient.recipients {
				q.enqueue(broadcast{sender: bc.recipient, recipient: r})
			}
		}
	}

	return lowPulseCount * highPulseCount
}

func parseBroadcast(s string) []broadcast {
	var broadcasts []broadcast
	ss := strings.Split(s, " -> ")
	sender := ss[0]
	recipient := ss[1]

	if sender == "broadcaster" {
		for _, r := range strings.Split(recipient, ", ") {
			broadcasts = append(broadcasts, broadcast{sender: sender, recipient: r})
		}
	}

	return broadcasts
}

func parseMode(s string, inputMap map[string]module) {
	ss := strings.Split(s, " -> ")
	sender := ss[0]
	var mode string

	if strings.Contains(sender, "%") || strings.Contains(sender, "&") {
		mode = sender[0:1]
		sender = sender[1:]
	} else {
		mode = "init"
	}

	module := inputMap[sender]
	module.mode = mode
	inputMap[sender] = module
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

	recipients := strings.Split(ss[1], ", ")

	for _, r := range recipients {
		if _, ok := inputMap[r]; !ok {
			inputMap[r] = module{
				inputs: []input{{name: sender, lastPulse: 0}},
			}
		} else {
			mod := inputMap[r]
			i := input{name: sender, lastPulse: 0}
			mod.inputs = append(mod.inputs, i)
			inputMap[r] = mod
		}
	}
}

func parseRecipients(s string, inputMap map[string]module) {
	ss := strings.Split(s, " -> ")
	sender := ss[0]

	if strings.Contains(sender, "%") || strings.Contains(sender, "&") {
		sender = sender[1:]
	}

	recipients := strings.Split(ss[1], ", ")

	mod := inputMap[sender]
	mod.recipients = recipients
	inputMap[sender] = mod
}

func TwentyTwo(o config.Options) int {
	scanner, file := utils.OpenFile(2023, 20, o)
	defer file.Close()
	pulseMap := make(map[string]int)
	inputMap := make(map[string]module)
	var initialBroadcasts []broadcast
	states := []map[string]int{}

	q := queue{}

	for scanner.Scan() {
		line := scanner.Text()
		parseInputs(line, inputMap)
		parseMode(line, inputMap)
		parsePulses(line, pulseMap)
		parseRecipients(line, inputMap)
		b := parseBroadcast(line)
		initialBroadcasts = append(initialBroadcasts, b...)
	}

	count := 0
	found := false
	for !found {
		for _, b := range initialBroadcasts {
			q.enqueue(b)
		}
		for len(q) > 0 {
			bc := q.dequeue()
			sPulse := pulseMap[bc.sender]
			recipient := inputMap[bc.recipient]

			if sPulse == 0 && bc.recipient == "rx" {
				found = true
				break
			}

			switch recipient.mode {
			case "%":
				if sPulse == 1 {
					continue
				} else {
					if pulseMap[bc.recipient] == 0 {
						pulseMap[bc.recipient] = 1
					} else {
						pulseMap[bc.recipient] = 0
					}
				}
			case "&":
				for i, input := range recipient.inputs {
					if input.name == bc.sender {
						recipient.inputs[i].lastPulse = sPulse
					}

					allHigh := true
					for _, input := range recipient.inputs {
						if input.lastPulse == 0 {
							allHigh = false
							break
						}
					}
					if allHigh {
						pulseMap[bc.recipient] = 0
					} else {
						pulseMap[bc.recipient] = 1
					}
				}
			}

			for _, r := range recipient.recipients {
				q.enqueue(broadcast{sender: bc.recipient, recipient: r})
			}
		}
		states = append(states, pulseMap)

		count += 1
	}

	return count
}
