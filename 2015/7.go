package _2015

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
)

type operation struct {
	destination string
	operands    []string
	operator    string
}

func SevenOne(o config.Options) int {
	scanner, file := utils.OpenFile(2015, 7, o)
	defer file.Close()

	operations := make(map[string]operation)
	values := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()

		operation := parseOperation(line)
		mapDependantOperations(operation, operations, values)
	}

	for _, o := range operations {
		computeWireValue(o.destination, operations, values)
	}

	return values["a"]
}

func SevenTwo(o config.Options) int {
	scanner, file := utils.OpenFile(2015, 7, o)
	defer file.Close()

	operations := make(map[string]operation)
	values := make(map[string]int)
	valuesTwo := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()

		operation := parseOperation(line)
		mapDependantOperations(operation, operations, values)

	}

	for k, v := range values {
		valuesTwo[k] = v
	}

	for _, o := range operations {
		computeWireValue(o.destination, operations, values)
	}

	valuesTwo["b"] = values["a"]

	for _, o := range operations {
		computeWireValue(o.destination, operations, valuesTwo)
	}

	return valuesTwo["a"]
}

func parseOperation(s string) operation {
	var o operation
	ss := strings.Split(s, "->")

	expression := strings.TrimSpace(ss[0])
	o.operands = findOperands(expression)
	o.operator = findCommand(expression)
	o.destination = strings.TrimSpace(ss[1])

	return o
}

func mapDependantOperations(o operation, dm map[string]operation, vm map[string]int) {
	if len(o.operands) == 1 {
		if num, err := strconv.Atoi(o.operands[0]); err == nil {
			vm[o.destination] = num
			return
		}
	}

	dm[o.destination] = o
}

func findCommand(s string) string {
	var c string
	regex, err := regexp.Compile(`\b[A-Z]+\b`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
	}

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

func computeWireValue(w string, dm map[string]operation, vm map[string]int) int {
	if val, ok := vm[w]; ok {
		return val
	}

	op := dm[w]
	switch op.operator {
	case "NOT":
		operandVal := getOperandValue(op.operands[0], dm, vm)
		vm[w] = ^operandVal
	case "AND":
		vm[w] = getOperandValue(op.operands[0], dm, vm) & getOperandValue(op.operands[1], dm, vm)
	case "OR":
		vm[w] = getOperandValue(op.operands[0], dm, vm) | getOperandValue(op.operands[1], dm, vm)
	case "LSHIFT":
		vm[w] = getOperandValue(op.operands[0], dm, vm) << getOperandValue(op.operands[1], dm, vm)
	case "RSHIFT":
		vm[w] = getOperandValue(op.operands[0], dm, vm) >> getOperandValue(op.operands[1], dm, vm)
	default:
		if val, err := strconv.Atoi(op.operands[0]); err == nil {
			vm[w] = val
		} else {
			vm[w] = computeWireValue(op.operands[0], dm, vm)
		}
	}

	return vm[w]
}

func getOperandValue(operand string, dm map[string]operation, vm map[string]int) int {
	if val, err := strconv.Atoi(operand); err == nil {
		return val // It's a number
	}
	return computeWireValue(operand, dm, vm) // It's a wire
}
