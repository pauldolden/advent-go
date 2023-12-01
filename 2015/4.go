package _2015

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

const PUZZLE_INPUT = "ckczppom"

func FourOne() int {
	count := 0

	for {
		s := strconv.Itoa(count)

		input := s

		hash := computeHexHash(input)

		ss := hash[:5]

		if ss == "00000" {
			break
		}

		count++
	}

	return count
}

func FourTwo() int {
	count := 0

	for {
		s := strconv.Itoa(count)

		input := s

		hash := computeHexHash(input)

		ss := hash[:6]

		if ss == "000000" {
			break
		}

		count++
	}

	return count
}

func computeHexHash(input string) string {
	i := []byte(PUZZLE_INPUT + input)

	h := md5.Sum(i)

	hh := hex.EncodeToString(h[:])

	return hh
}
