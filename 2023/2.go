package _2023

import (
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/pauldolden/advent-go/config"
	"github.com/pauldolden/advent-go/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type cubes struct {
	Red   int
	Blue  int
	Green int
}

func TwoOne(o config.Options) int {
	// Max per game
	mc := cubes{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	results := make(map[int]bool)

	scanner, file := utils.OpenFile(2023, 2, o)
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()

		parseLineOne(line, results, mc)
	}

	count := 0

	for k, v := range results {
		if v {
			count += k
		}
	}

	return count
}

func TwoTwo(o config.Options) int {
	scanner, file := utils.OpenFile(2023, 2, o)
	defer file.Close()

	results := make(map[int]int)

	for scanner.Scan() {

		line := scanner.Text()

		parseLineTwo(line, results)
	}

	count := 0

	for _, v := range results {
		count += v
	}

	return count
}

func parseLineOne(s string, m map[int]bool, mc cubes) {
	id, gs := getId(s)

	games := strings.Split(gs, ";")
	result := true

	for _, game := range games {
		cubes := strings.Split(game, ",")
		for _, cube := range cubes {
			trimmedCube := strings.TrimSpace(cube)
			operands := strings.Split(trimmedCube, " ")

			num, err := strconv.Atoi(operands[0])
			if err != nil {
				log.Fatal(err.Error())
			}
			colour := cases.Title(language.English).String(operands[1])
			r := reflect.ValueOf(mc)
			f := r.FieldByName(colour)

			if f.IsValid() {
				if fieldValue, ok := f.Interface().(int); ok {
					if num > fieldValue {
						result = false
					}
				}
			}
		}
	}

	m[id] = result
}

func parseLineTwo(s string, m map[int]int) {
	id, gs := getId(s)

	games := strings.Split(gs, ";")

	c := cubes{
		Red:   0,
		Blue:  0,
		Green: 0,
	}

	for _, game := range games {
		cubes := strings.Split(game, ",")
		for _, cube := range cubes {
			trimmedCube := strings.TrimSpace(cube)
			operands := strings.Split(trimmedCube, " ")

			num, err := strconv.Atoi(operands[0])
			if err != nil {
				log.Fatal(err.Error())
			}
			colour := cases.Title(language.English).String(operands[1])
			r := reflect.ValueOf(&c).Elem()
			f := r.FieldByName(colour)

			if f.IsValid() {
				if fieldValue, ok := f.Interface().(int); ok {
					if num > fieldValue {
						f.SetInt(int64(num))
					}
				}
			}
		}
	}

	m[id] = c.Red * c.Green * c.Blue
}

func getId(s string) (int, string) {
	res := strings.Split(s, ":")
	gameNumber := res[0]
	gameString := res[1]

	regex, err := regexp.Compile("[0-9]+")

	if err != nil {
		log.Fatal(err)
	}

	idString := regex.FindAllString(gameNumber, -1)[0]

	id, err := strconv.Atoi(idString)

	if err != nil {
		log.Fatal(err)
	}

	return id, gameString
}
