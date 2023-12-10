#!/bin/sh

mkdir -p "day-$1/part-1/" "day-$1/part-2/"

echo 'package main

import (
	"fmt"

	"github.com/ItsNotGoodName/aoc-2023"
)

func main() {
	lines := aoc.Read()

	fmt.Println(lines)
}
' | tee "day-$1/part-1/main.go" "day-$1/part-2/main.go" > /dev/null

touch "day-$1/input.txt"
