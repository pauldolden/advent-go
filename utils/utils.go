package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func OpenFile(year int, day int) (*bufio.Scanner, *os.File) {
	path := fmt.Sprintf("./%d/%d.txt", year, day)
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	return scanner, file
}
