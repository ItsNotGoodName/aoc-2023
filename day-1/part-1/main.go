package main

import (
	"fmt"
	"strconv"

	"github.com/ItsNotGoodName/aoc-2023"
)

func main() {
	lines := aoc.Read()

	var sum int

	for _, line := range lines {
		first, last := "", ""

		for _, r := range line {
			if r >= 48 && r <= 57 {
				if first == "" {
					first = string(r)
					if last == "" {
						last = first
					}
				} else {
					last = string(r)
				}
			}
		}

		i, err := strconv.Atoi(first + last)
		if err != nil {
			panic(err)
		}
		sum += i
	}

	fmt.Println(sum)
}
