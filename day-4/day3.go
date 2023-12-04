package day4

import (
	"strconv"
	"strings"
)

type Card struct {
	ID      int
	Winning []int
	Numbers []int
}

func ParseCards(lines []string) []Card {
	var cards []Card
	for _, line := range lines {
		var card Card

		col := strings.Split(line, ":")

		id, err := strconv.Atoi(strings.TrimRight(strings.TrimSpace(strings.TrimLeft(col[0], "Card")), ":"))
		if err != nil {
			panic(err)
		}
		card.ID = id

		sets := strings.Split(col[1], "|")
		for i, set := range sets {
			var numbers []int
			for _, s := range strings.Split(strings.TrimSpace(set), " ") {
				if s == "" {
					continue
				}
				number, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				numbers = append(numbers, number)
			}
			if i == 0 {
				card.Winning = numbers
			} else if i == 1 {
				card.Numbers = numbers
			} else {
				panic("WTF")
			}
		}

		cards = append(cards, card)
	}
	return cards
}
