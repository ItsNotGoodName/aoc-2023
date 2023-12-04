package main

import (
	"fmt"

	"github.com/ItsNotGoodName/aoc-2023"
	day3 "github.com/ItsNotGoodName/aoc-2023/day-3"
)

func main() {
	lines := aoc.Read()

	var sum int

	for _, n := range day3.ParseNumbers(lines) {
		if len(n.Around(lines)) > 0 {
			sum += n.Number
		}
	}

	fmt.Println(sum)
}
