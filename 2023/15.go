package _2023

import (
	"os"
	"strconv"
	"strings"

	"github.com/pauldolden/advent-go/config"
)

func FifteenOne(o config.Options) int {
	bs, err := os.ReadFile("./2023/15.txt")
	if err != nil {
		panic(err)
	}

	s := string(bs)
	ss := strings.Split(s, ",")

	total := 0
	for _, s := range ss {
		s = strings.TrimSpace(s)
		total += hash(s)
	}

	return total
}

func FifteenTwo(o config.Options) int {
	bs, err := os.ReadFile("./2023/15.txt")
	if err != nil {
		panic(err)
	}

	s := string(bs)
	ss := strings.Split(s, ",")

	hm := make(map[int][]lens)

	total := 0
	for _, s := range ss {
		s = strings.TrimSpace(s)

		if s[len(s)-1] == '-' {
			ss := strings.Split(s, "-")
			label, box := ss[0], hash(ss[0])
			hm[box] = remove(label, hm[box])
		} else {
			ss := strings.Split(s, "=")
			box, label, value := hash(ss[0]), ss[0], ss[1]

			if i, err := strconv.Atoi(value); err == nil {
				lens := lens{label, i}

				if lens.exists(hm[box]) {
					hm[box] = lens.update(hm[box])
				} else {
					hm[box] = append(hm[box], lens)
				}
			}
		}
	}

	for i, ls := range hm {
		box := i + 1
		for k, l := range ls {
			pos := k + 1
			total += box * pos * l.value
		}
	}

	return total
}

type lens struct {
	label string
	value int
}

func remove(label string, ls []lens) []lens {
	l := []lens{}
	for _, l2 := range ls {
		if l2.label != label {
			l = append(l, l2)
		}
	}

	return l
}

func (l *lens) update(ls []lens) []lens {
	for i, l2 := range ls {
		if l2.label == l.label {
			ls[i] = *l
			break
		}
	}

	return ls
}

func (l *lens) exists(ls []lens) bool {
	found := false

	for _, l2 := range ls {
		if l2.label == l.label {
			found = true
			break
		}
	}

	return found
}

func hash(s string) int {
	total := 0
	for _, c := range s {
		if c == 10 {
			continue
		}
		total += int(c)
		total *= 17
		total %= 256
	}
	return total
}
