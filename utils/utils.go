package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/pauldolden/advent-go/config"
)

func OpenFile(year int, day int, o config.Options) (*bufio.Scanner, *os.File) {
	var path string
	if o.Test {
		if o.SplitInputs {
			path = fmt.Sprintf("./%d_test_%d.txt", day, o.TestPart)
		} else {
			path = fmt.Sprintf("./%d_test.txt", day)
		}
	} else {
		path = fmt.Sprintf("./%d/%d.txt", year, day)
	}

	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	return scanner, file
}
