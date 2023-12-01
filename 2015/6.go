package _2015

import (
	"log"
	"strconv"
	"strings"
	"unicode"

	"github.com/pauldolden/advent-go/utils"
)

type point struct {
	x int
	y int
}

type action struct {
	start   point
	end     point
	command string
}

const GRID_SIZE = 1000

func SixOne() int {
	scanner, file := utils.OpenFile(2015, 6)
	defer file.Close()

	grid := buildBoolGrid()

	for scanner.Scan() {
		line := scanner.Text()

		a := parseAction(line)
		executeBoolAction(a, &grid)
	}

	return countLights(&grid)
}

func SixTwo() int {
	scanner, file := utils.OpenFile(2015, 6)
	defer file.Close()

	grid := buildIntGrid()

	for scanner.Scan() {
		line := scanner.Text()

		a := parseAction(line)
		executeIntAction(a, &grid)
	}

	return countBrightness(&grid)
}

func buildBoolGrid() map[int][]bool {
	grid := make(map[int][]bool)

	for i := 0; i < GRID_SIZE; i++ {
		grid[i] = make([]bool, GRID_SIZE)
	}

	return grid
}

func buildIntGrid() map[int][]int {
	grid := make(map[int][]int)

	for i := 0; i < GRID_SIZE; i++ {
		grid[i] = make([]int, GRID_SIZE)
	}

	return grid
}

func parseAction(rawAction string) action {
	a := action{}

	actionString, coords := splitAtFirstNumber(rawAction)

	a.command = actionString

	start, end := splitAtFirstWord(coords, "through")

	setPoint(start, &a.start)
	setPoint(end, &a.end)

	return a
}

func splitAtFirstNumber(s string) (string, string) {
	for i, r := range s {
		if unicode.IsDigit(r) {
			return strings.Trim(s[:i], " "), strings.Trim(s[i:], " ")
		}
	}
	return s, ""
}

func splitAtFirstWord(s, word string) (string, string) {
	index := strings.Index(s, word)
	if index == -1 {
		return s, ""
	}
	return strings.Trim(s[:index], " "), strings.Trim(s[index+len(word):], " ")
}

func setPoint(s string, p *point) {
	ss := strings.Split(s, ",")

	x, err := strconv.Atoi(ss[0])

	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(ss[1])

	if err != nil {
		log.Fatal(err)
	}

	p.x = x
	p.y = y
}

func executeBoolAction(a action, g *map[int][]bool) {
	for i := a.start.y; i <= a.end.y; i++ {
		for k := a.start.x; k <= a.end.x; k++ {
			switch a.command {
			case "turn on":
				(*g)[i][k] = true
			case "turn off":
				(*g)[i][k] = false
			case "toggle":
				(*g)[i][k] = !(*g)[i][k]
			}
		}
	}
}

func executeIntAction(a action, g *map[int][]int) {
	for i := a.start.y; i <= a.end.y; i++ {
		for k := a.start.x; k <= a.end.x; k++ {
			switch a.command {
			case "turn on":
				(*g)[i][k] = (*g)[i][k] + 1
			case "turn off":
				if (*g)[i][k] == 0 {
					continue
				}
				(*g)[i][k] = (*g)[i][k] - 1
			case "toggle":
				(*g)[i][k] = (*g)[i][k] + 2
			}
		}
	}
}

func countLights(g *map[int][]bool) int {
	var count int

	for _, v := range *g {
		for _, b := range v {
			if b {
				count++
			}
		}
	}

	return count
}

func countBrightness(g *map[int][]int) int {
	var count int

	for _, v := range *g {
		for _, b := range v {
			count += b
		}
	}

	return count
}
