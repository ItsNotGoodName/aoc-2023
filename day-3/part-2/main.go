package main

import (
	"fmt"

	"github.com/ItsNotGoodName/aoc-2023"
	day3 "github.com/ItsNotGoodName/aoc-2023/day-3"
)

func main() {
	lines := aoc.Read()
	allNumbers := day3.ParseNumbers(lines)

	var parts []day3.Part
	var numbers []day3.Number

	for _, n := range allNumbers {
		for _, p := range n.Around(lines) {
			if p.Symbol == '*' {
				parts = append(parts, p)
				numbers = append(numbers, n)
			}
		}
	}

	var sum int

	for _, p := range parts {
		var indexes []int
		for i, p2 := range parts {
			if p.Y == p2.Y && p.X == p2.X {
				indexes = append(indexes, i)
			}
		}
		if len(indexes) == 2 {
			sum += numbers[indexes[0]].Number * numbers[indexes[1]].Number
		}
	}

	fmt.Println(sum / 2)
}
