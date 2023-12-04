package aoc

import (
	"os"
	"strings"
)

func Read() []string {
	raw, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(raw)), "\n")
}
