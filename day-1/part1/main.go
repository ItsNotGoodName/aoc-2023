package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	raw, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(raw)), "\n")

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

