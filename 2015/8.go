package _2015

import (
	"regexp"
	"strings"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
)

func EightOne(o config.Options) int {
	scanner, file := utils.OpenFile(2015, 8, o)
	defer file.Close()

	var rawLen, transformedLen int
	hexRegex := regexp.MustCompile(`\\x[0-9a-fA-F]{2}`)

	for scanner.Scan() {
		line := scanner.Text()
		rawLen += len(line)

		// Handling escape sequences
		transformedLine := strings.ReplaceAll(
			line,
			`\\`,
			`S`,
		) // Single character placeholder for \\
		transformedLine = strings.ReplaceAll(
			transformedLine,
			`\"`,
			`Q`,
		) // Single character placeholder for \"

		// Handling hexadecimal escape sequences
		transformedLine = hexRegex.ReplaceAllStringFunc(transformedLine, func(s string) string {
			return "H" // Single character placeholder for \xNN
		})

		// Remove the surrounding double quotes
		transformedLine = transformedLine[1 : len(transformedLine)-1]

		transformedLen += len(transformedLine)
	}

	return rawLen - transformedLen
}

func EightTwo() int {
	return 0
}
