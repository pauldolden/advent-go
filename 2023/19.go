package _2023

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pauldolden/advent-go/config"
)

type condition struct {
	operand     string
	condition   string
	target      int
	destination string
}

type part struct {
	x, m, a, s int
}

type Range struct {
	low, high int
}

func NineteenOne(o config.Options) int {
	bs, err := os.ReadFile("./2023/19.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(bs), "\n\n")
	rw := strings.Split(input[0], "\n")
	rp := strings.Split(input[1], "\n")

	return solveOne(processWorkflows(rw), processParts(rp))
}

func NineteenTwo(o config.Options) int {
	bs, err := os.ReadFile("./2023/19.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(bs), "\n\n")
	rw := strings.Split(input[0], "\n")

	workflows := processWorkflows(rw)

	ranges := make(map[string]Range)

	for _, c := range "xmas" {
		ranges[string(c)] = Range{1, 4000}
	}

	return solveTwo("in", workflows, ranges)
}

func solveTwo(destination string, workflows map[string][]condition, ranges map[string]Range) int {
	if destination == "R" {
		return 0
	}
	if destination == "A" {
		product := 1
		for _, r := range ranges {
			product *= r.high - r.low + 1
		}
		return product
	}

	total := 0

	for _, c := range workflows[destination] {
		var T, F Range
		hi, lo := ranges[c.operand].high, ranges[c.operand].low
		if c.condition == "<" {
			T = Range{lo, c.target - 1}
			F = Range{c.target, hi}
		} else {
			T = Range{c.target + 1, hi}
			F = Range{lo, c.target}
		}

		if T.low <= T.high {
			if c.operand == "GOTO" {
				continue
			}

			TRanges := make(map[string]Range)
			for _, r := range ranges {
				TRanges[string(c.operand)] = Range{r.low, r.high}
			}

			total += solveTwo(c.destination, workflows, TRanges)
		}

		if F.low <= F.high {
			if c.operand == "GOTO" {
				continue
			}

			FRanges := make(map[string]Range)

			for _, r := range ranges {
				FRanges[string(c.operand)] = Range{r.low, r.high}
			}

			ranges[string(c.operand)] = F
		}
	}

	return total
}

func solveOne(workflows map[string][]condition, parts []part) int {
	var accepted []part
	var rejected []part

	for _, p := range parts {
		workflow := workflows["in"]

		e := evaluate(workflow, p, workflows)

		if e {
			accepted = append(accepted, p)
		} else {
			rejected = append(rejected, p)
		}
	}

	total := 0
	for _, p := range accepted {
		total += p.x + p.m + p.a + p.s
	}

	fmt.Println("Rejected", len(rejected))
	fmt.Println("Accepted", len(accepted))

	return total
}

func evaluate(workflow []condition, p part, workflows map[string][]condition) bool {
	for _, w := range workflow {
		switch w.operand {
		case "GOTO":
			return evaluateGoto(w, workflows, p)

		case "x":
			if evaluateCondition(w.condition, p.x, w.target) {
				return evaluateDestination(w, workflows, p)
			}

		case "m":
			if evaluateCondition(w.condition, p.m, w.target) {
				return evaluateDestination(w, workflows, p)
			}

		case "a":
			if evaluateCondition(w.condition, p.a, w.target) {
				return evaluateDestination(w, workflows, p)
			}

		case "s":
			if evaluateCondition(w.condition, p.s, w.target) {
				return evaluateDestination(w, workflows, p)
			}
		}
	}
	return true
}

func evaluateGoto(w condition, workflows map[string][]condition, p part) bool {
	switch w.destination {
	case "A":
		return true
	case "R":
		return false
	default:
		return evaluate(workflows[w.destination], p, workflows)
	}
}

func evaluateDestination(w condition, workflows map[string][]condition, p part) bool {
	switch w.destination {
	case "A":
		return true
	case "R":
		return false
	default:
		return evaluate(workflows[w.destination], p, workflows)
	}
}

func evaluateCondition(condition string, operand, target int) bool {
	switch condition {
	case "<":
		return operand < target
	case ">":
		return operand > target
	case "=":
		return operand == target
	}
	return false
}

func processParts(ss []string) []part {
	var parts []part

	for _, s := range ss {
		if s == "" {
			continue
		}
		s = s[1 : len(s)-1]

		var p part
		ss := strings.Split(s, ",")

		for _, s := range ss {
			s = strings.TrimSpace(s)
			switch {
			case strings.HasPrefix(s, "x="):
				p.x = atoi(s[2:])
			case strings.HasPrefix(s, "m="):
				p.m = atoi(s[2:])
			case strings.HasPrefix(s, "a="):
				p.a = atoi(s[2:])
			case strings.HasPrefix(s, "s="):
				p.s = atoi(s[2:])
			}
		}

		parts = append(parts, p)
	}

	return parts
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return i
}

func processWorkflows(ss []string) map[string][]condition {
	workflows := make(map[string][]condition)

	for _, s := range ss {
		if s == "" {
			continue
		}

		ss := strings.Split(s, "{")
		key := ss[0]

		s = ss[1][:len(ss[1])-1]

		var conditions []condition

		for _, s := range strings.Split(s, ",") {
			var c condition

			// if it contains a : then it's a condition otherwise it's a destination
			if !strings.Contains(s, ":") {
				c.operand = "GOTO"
				c.condition = "->"
				c.destination = s
				conditions = append(conditions, c)
				continue
			}

			ss := strings.Split(s, ":")

			c.operand = string(ss[0][0])
			c.condition = string(ss[0][1])
			c.target = atoi(ss[0][2:])
			c.destination = ss[1]

			conditions = append(conditions, c)
		}

		workflows[key] = conditions
	}

	return workflows
}
