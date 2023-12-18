package _2023

import (
	"math"
	"strconv"
	"strings"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
)

type dig struct {
	dir    int
	moves  int
	colour string
}

var dirMap18 = map[string]int{
	"U": 0,
	"R": 1,
	"D": 2,
	"L": 3,
}

type point struct {
	r, c int
}

// dirs18 is a slice of directions in the order of up, right, down, left
var dirs18 = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

// we want right, down left up
var dirs18Two = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

var dirMap18Two = map[int]int{
	0: 0,
	1: 1,
	2: 2,
	3: 3,
}

func EighteenOne(o config.Options) int {
	scanner, file := utils.OpenFile(2023, 18, o)
	defer file.Close()
	var digs []dig
	var points []point

	for scanner.Scan() {
		line := scanner.Text()

		ss := parse18Input(line)

		digs = append(digs, parseDig(ss))
	}

	r, c := 0, 0
	for _, d := range digs {
		for i := 0; i < d.moves; i++ {
			r += dirs18[d.dir][0]
			c += dirs18[d.dir][1]
			points = append(points, point{r, c})
		}
	}

	area := ShoelaceArea(points)
	b := len(points)

	i := area - (b / 2) + 1

	return i + b
}

func EighteenTwo(o config.Options) int {
	scanner, file := utils.OpenFile(2023, 18, o)
	defer file.Close()
	var digs []dig
	var points []point

	for scanner.Scan() {
		line := scanner.Text()

		ss := parse18Input(line)

		digs = append(digs, parseDigTwo(ss))
	}

	r, c := 0, 0
	for _, d := range digs {
		for i := 0; i < d.moves; i++ {
			r += dirs18Two[d.dir][0]
			c += dirs18Two[d.dir][1]
			points = append(points, point{r, c})
		}
	}

	area := ShoelaceArea(points)
	b := len(points)

	i := area - (b / 2) + 1

	return i + b
}

func parseDig(ss []string) dig {
	colour := ss[2][1 : len(ss[2])-1]
	moves, _ := strconv.Atoi(ss[1])
	return dig{
		dir:    dirMap18[ss[0]],
		moves:  moves,
		colour: colour,
	}
}

func parseDigTwo(ss []string) dig {
	colour := ss[2][2 : len(ss[2])-1]
	hex := colour[:len(colour)-1]
	dir := colour[len(colour)-1:]

	di, _ := strconv.Atoi(dir)
	moves, _ := strconv.ParseInt(hex, 16, 64)

	return dig{
		dir:    dirMap18Two[di],
		moves:  int(moves),
		colour: colour,
	}
}

func parse18Input(input string) []string {
	return strings.Fields(input)
}

// ShoelaceArea calculates the area of a polygon defined by a slice of Points
func ShoelaceArea(points []point) int {
	if len(points) < 3 {
		// A polygon must have at least 3 vertices
		return 0
	}

	var area float64
	j := len(points) - 1 // The last vertex is the 'previous' one to the first

	for i := 0; i < len(points); i++ {
		area += float64(points[j].c+points[i].c) * float64(points[j].r-points[i].r)
		j = i // j is previous vertex to i
	}

	return int(math.Abs(area) / 2)
}
