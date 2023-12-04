package main

import (
	"fmt"

	"github.com/ItsNotGoodName/aoc-2023"
	day2 "github.com/ItsNotGoodName/aoc-2023/day-2"
)

func main() {
	lines := aoc.Read()
	games := day2.ParseGames(lines)

	var sum int

outer:
	for _, game := range games {
		for _, v := range game.Red {
			if v > 12 {
				continue outer
			}
		}
		for _, v := range game.Green {
			if v > 13 {
				continue outer
			}
		}
		for _, v := range game.Blue {
			if v > 14 {
				continue outer
			}
		}

		sum += game.ID
	}

	fmt.Println(sum)
}
