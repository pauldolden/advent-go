package _2015

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
)

func SevenOne(o config.Options) int {
	scanner, file := utils.OpenFile(2015, 7, o)
	defer file.Close()

	operations := make(map[string][]string)
	values := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()

		operation := parseOperation(line)
		mapDependantOperations(operation, operations, values)
	}

	for k := range operations {
		computeWireValue(k, operations, values)
	}

	return 0
}

func SevenTwo() int {
	return 0
}

type operation struct {
	destination string
	expression  string
}

func parseOperation(s string) operation {
	var o operation
	ss := strings.Split(s, "->")

	o.expression = strings.TrimSpace(ss[0])
	o.destination = strings.TrimSpace(ss[1])

	return o
}

func mapDependantOperations(o operation, dm map[string][]string, vm map[string]int) {
	operands := findOperands(o.expression)

	if len(operands) == 1 {
		if num, err := strconv.Atoi(operands[0]); err == nil {
			vm[o.destination] = num
			return
		}
	}

	dm[o.destination] = operands
}

func findCommand(s string) string {
	var c string
	regex, err := regexp.Compile(`\b[A-Z]+\b`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
	}

	// Find all matches
	matches := regex.FindAllString(s, -1)

	if len(matches) > 0 {
		c = matches[0]
	}

	return c
}

func findOperands(s string) []string {
	var o []string
	regex, err := regexp.Compile(`\b[a-z]{1,2}|[0-9]+\b`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return o
	}

	// Find all matches
	matches := regex.FindAllString(s, -1)
	o = matches

	return o
}

func computeWireValue(w string, dm map[string][]string, vm map[string]int) int {
	val, ok := vm[w]

	if ok {
		return val
	} else {
		for _, v := range dm[w] {
			if v == "b" {
				fmt.Println(w)
			}
		}
	}
	return 0
}
