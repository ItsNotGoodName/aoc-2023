package main

import (
	"fmt"
	"slices"

	"github.com/ItsNotGoodName/aoc-2023"
	day2 "github.com/ItsNotGoodName/aoc-2023/day-2"
)

func main() {
	lines := aoc.Read()
	games := day2.ParseGames(lines)

	var sum int

	for _, game := range games {
		red := slices.Max(game.Red)
		blue := slices.Max(game.Blue)
		green := slices.Max(game.Green)

		sum += red * blue * green
	}

	fmt.Println(sum)
}
