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

	cardCount := make([]int, len(cards))

	for i, card := range cards {
		cardCount[i]++

		var count int
		for _, w := range card.Winning {
			if slices.Contains(card.Numbers, w) {
				count++
			}
		}

		for c := 1; c <= count; c++ {
			if i+c > len(cards) {
				break
			}

			cardCount[i+c] += cardCount[i]
		}
	}

	var sum int
	for _, v := range cardCount {
		sum += v
	}

	fmt.Println(sum)
}
