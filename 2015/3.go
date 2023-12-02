package _2015

import (
	"fmt"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
)

type Point struct {
	x int
	y int
}

func ThreeOne(o config.Options) int {
	cur := Point{
		x: 0,
		y: 0,
	}

	visits := make(map[string]int)

	visit(&cur, &visits)

	scanner, file := utils.OpenFile(2015, 3, o)
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()

		for _, char := range line {
			dir := string(char)

			move(&cur, dir)
			visit(&cur, &visits)
		}
	}

	return len(visits)
}

func ThreeTwo(o config.Options) int {
	scur := Point{
		x: 0,
		y: 0,
	}

	rscur := Point{
		x: 0,
		y: 0,
	}

	visits := make(map[string]int)

	visit(&scur, &visits)
	visit(&rscur, &visits)

	scanner, file := utils.OpenFile(2015, 3, o)
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()

		for index, char := range line {
			rs := index%2 != 0

			var cur *Point

			dir := string(char)

			if rs {
				cur = &rscur
			} else {
				cur = &scur
			}

			move(cur, dir)
			visit(cur, &visits)
		}
	}

	return len(visits)
}

func move(current *Point, dir string) {
	switch dir {
	case "^":
		current.y += 1
	case ">":
		current.x += 1
	case "<":
		current.x -= 1
	case "v":
		current.y -= 1
	}
}

func visit(current *Point, visits *map[string]int) {
	key := fmt.Sprintf("%d,%d", current.x, current.y)

	(*visits)[key] += 1
}
