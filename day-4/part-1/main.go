package main

import (
	"fmt"
	"slices"

	"github.com/ItsNotGoodName/aoc-2023"
	day4 "github.com/ItsNotGoodName/aoc-2023/day-4"
)

func main() {
	lines := aoc.Read()

	cards := day4.ParseCards(lines)

	var sum int

	for _, card := range cards {
		var count int
		for _, w := range card.Winning {
			if slices.Contains(card.Numbers, w) {
				count += 1
			}
		}

		if count > 0 {
			sum += (1 << (count - 1))
		}
	}

	fmt.Println(sum)
}
