package _2015

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pauldolden/advent-go/utils"
)

func SevenOne() int {
	scanner, file := utils.OpenFile(2015, 7)
	defer file.Close()

	values := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()

		operation := parseOperation(line)

		evaluateOperation(operation)
	}

	fmt.Println(values)

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

func evaluateOperation(o operation) {
	command := findCommand(o.expression)
	operands := findOperands(o.expression)

	fmt.Println(command)
	fmt.Println(operands)
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
