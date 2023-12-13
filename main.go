package main

import (
	"flag"
	"fmt"

	_2015 "github.com/pauldolden/advent-go/2015"
	_2016 "github.com/pauldolden/advent-go/2016"
	_2023 "github.com/pauldolden/advent-go/2023"
	"github.com/pauldolden/advent-go/config"
)

func main() {
	year := flag.Int("year", 2015, "Year to run")
	day := flag.Int("day", 1, "Day to run")

	o := config.NewDefaultOptions()

	flag.Parse()

	fmt.Printf("Running %d/%d\n", *year, *day)

	switch *year {
	case 2015:
		switch *day {
		case 1:
			fmt.Println(_2015.OneOne(o))
			fmt.Println(_2015.OneTwo(o))
		case 2:
			fmt.Println(_2015.TwoOne(o))
			fmt.Println(_2015.TwoTwo(o))
		case 3:
			fmt.Println(_2015.ThreeOne(o))
			fmt.Println(_2015.ThreeTwo(o))
		case 4:
			fmt.Println(_2015.FourOne())
			fmt.Println(_2015.FourTwo())
		case 5:
			fmt.Println(_2015.FiveOne(o))
			fmt.Println(_2015.FiveTwo(o))
		case 6:
			fmt.Println(_2015.SixOne(o))
			fmt.Println(_2015.SixTwo(o))
		case 7:
			fmt.Println(_2015.SevenOne(o))
			fmt.Println(_2015.SevenTwo(o))
		case 8:
			fmt.Println(_2015.EightOne(o))
			fmt.Println(_2015.EightTwo())
		case 9:
			fmt.Println(_2015.NineOne())
			fmt.Println(_2015.NineTwo())
		case 10:
			fmt.Println(_2015.TenOne())
			fmt.Println(_2015.TenTwo())
		case 11:
			fmt.Println(_2015.ElevenOne())
			fmt.Println(_2015.ElevenTwo())
		case 12:
			fmt.Println(_2015.TwelveOne())
			fmt.Println(_2015.TwelveTwo())
		case 13:
			fmt.Println(_2015.ThirteenOne())
			fmt.Println(_2015.ThirteenTwo())
		case 14:
			fmt.Println(_2015.FourteenOne())
			fmt.Println(_2015.FourteenTwo())
		case 15:
			fmt.Println(_2015.FifteenOne())
			fmt.Println(_2015.FifteenTwo())
		case 16:
			fmt.Println(_2015.SixteenOne())
			fmt.Println(_2015.SixteenTwo())
		case 17:
			fmt.Println(_2015.SeventeenOne())
			fmt.Println(_2015.SeventeenTwo())
		case 19:
			fmt.Println(_2015.NineteenOne())
			fmt.Println(_2015.NineteenTwo())
		case 20:
			fmt.Println(_2015.TwentyOne())
			fmt.Println(_2015.TwentyTwo())
		case 21:
			fmt.Println(_2015.TwentyOneOne())
			fmt.Println(_2015.TwentyOneTwo())
		case 22:
			fmt.Println(_2015.TwentyTwoOne())
			fmt.Println(_2015.TwentyTwoTwo())
		case 23:
			fmt.Println(_2015.TwentyThreeOne())
			fmt.Println(_2015.TwentyThreeTwo())
		case 24:
			fmt.Println(_2015.TwentyFourOne())
			fmt.Println(_2015.TwentyFourTwo())
		case 25:
			fmt.Println(_2015.TwentyFiveOne())
			fmt.Println(_2015.TwentyFiveTwo())

		default:
			fmt.Println("Day not implemented")
		}
	case 2016:
		switch *day {
		case 1:
			fmt.Println(_2016.OneOne())
			fmt.Println(_2016.OneTwo())
		case 2:
			fmt.Println(_2016.TwoOne())
			fmt.Println(_2016.TwoTwo())
		case 3:
			fmt.Println(_2016.ThreeOne())
			fmt.Println(_2016.ThreeTwo())
		case 4:
			fmt.Println(_2016.FourOne())
			fmt.Println(_2016.FourTwo())
		case 5:
			fmt.Println(_2016.FiveOne())
			fmt.Println(_2016.FiveTwo())
		case 6:
			fmt.Println(_2016.SixOne())
			fmt.Println(_2016.SixTwo())
		case 7:
			fmt.Println(_2016.SevenOne())
			fmt.Println(_2016.SevenTwo())
		case 8:
			fmt.Println(_2016.EightOne())
			fmt.Println(_2016.EightTwo())
		case 9:
			fmt.Println(_2016.NineOne())
			fmt.Println(_2016.NineTwo())
		case 10:
			fmt.Println(_2016.TenOne())
			fmt.Println(_2016.TenTwo())
		case 11:
			fmt.Println(_2016.ElevenOne())
			fmt.Println(_2016.ElevenTwo())
		case 12:
			fmt.Println(_2016.TwelveOne())
			fmt.Println(_2016.TwelveTwo())
		case 13:
			fmt.Println(_2016.ThirteenOne())
			fmt.Println(_2016.ThirteenTwo())
		case 14:
			fmt.Println(_2016.FourteenOne())
			fmt.Println(_2016.FourteenTwo())
		case 15:
			fmt.Println(_2016.FifteenOne())
			fmt.Println(_2016.FifteenTwo())
		case 16:
			fmt.Println(_2016.SixteenOne())
			fmt.Println(_2016.SixteenTwo())
		case 17:
			fmt.Println(_2016.SeventeenOne())
			fmt.Println(_2016.SeventeenTwo())
		case 19:
			fmt.Println(_2016.NineteenOne())
			fmt.Println(_2016.NineteenTwo())
		case 20:
			fmt.Println(_2016.TwentyOne())
			fmt.Println(_2016.TwentyTwo())
		case 21:
			fmt.Println(_2016.TwentyOneOne())
			fmt.Println(_2016.TwentyOneTwo())
		case 22:
			fmt.Println(_2016.TwentyTwoOne())
			fmt.Println(_2016.TwentyTwoTwo())
		case 23:
			fmt.Println(_2016.TwentyThreeOne())
			fmt.Println(_2016.TwentyThreeTwo())
		case 24:
			fmt.Println(_2016.TwentyFourOne())
			fmt.Println(_2016.TwentyFourTwo())
		case 25:
			fmt.Println(_2016.TwentyFiveOne())
			fmt.Println(_2016.TwentyFiveTwo())

		default:
			fmt.Println("Day not implemented")
		}
	case 2023:
		switch *day {
		case 1:
			fmt.Println(_2023.OneOne(o))
			fmt.Println(_2023.OneTwo(o))
		case 2:
			fmt.Println(_2023.TwoOne(o))
			fmt.Println(_2023.TwoTwo(o))
		case 3:
			fmt.Println(_2023.ThreeOne(o))
			fmt.Println(_2023.ThreeTwo(o))
		case 4:
			fmt.Println(_2023.FourOne(o))
			fmt.Println(_2023.FourTwo(o))
		case 5:
			fmt.Println(_2023.FiveOne(o))
			fmt.Println(_2023.FiveTwo(o))
		case 6:
			fmt.Println(_2023.SixOne(o))
			fmt.Println(_2023.SixTwo(o))
		case 7:
			fmt.Println(_2023.SevenOne(o))
			fmt.Println(_2023.SevenTwo(o))
		case 8:
			fmt.Println(_2023.EightOne(o))
			fmt.Println(_2023.EightTwo(o))
		case 13:
			fmt.Println(_2023.ThirteenTwo(o))
		}

	default:
		fmt.Println("Year not implemented")
	}
}
