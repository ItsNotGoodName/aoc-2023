package day3

import (
	"strconv"
	"unicode"
)

type Part struct {
	Symbol rune
	Y      int
	X      int
}

type Number struct {
	Length int
	Number int
	Y      int
	X      int
}

type Point struct {
	Y int
	X int
}

var aroundMatrix []Point = []Point{
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
	{-1, 0},
	{-1, 1},
}

func (n Number) Around(lines []string) []Part {
	var parts []Part
	for xOffset := 0; xOffset < n.Length; xOffset++ {
	outer:
		for _, mat := range aroundMatrix {
			y, x := n.Y+mat.Y, n.X+mat.X+xOffset
			if y < 0 || len(lines) <= y || x < 0 || len(lines[y]) <= x {
				continue
			}
			if lines[y][x] == '.' || unicode.IsNumber(rune(lines[y][x])) {
				continue
			}
			for _, p := range parts {
				if p.X == x && p.Y == y {
					continue outer
				}
			}

			parts = append(parts, Part{
				Symbol: rune(lines[y][x]),
				Y:      y,
				X:      x,
			})
		}
	}

	return parts
}

// func ParseParts(lines []string) []Part {
// 	var parts []Part
// 	for y, line := range lines {
// 		for x, s := range line {
// 			if s == '.' {
// 				break
// 			}
// 			_, err := strconv.Atoi(string(s))
// 			if err != nil {
// 				break
// 			}
//
// 			parts = append(parts, Part{
// 				Symbol: s,
// 				Y:      y,
// 				X:      x,
// 			})
// 		}
// 	}
//
// 	return parts
// }

func ParseNumbers(lines []string) []Number {
	var parts []Number
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			if !unicode.IsNumber(rune(line[x])) {
				continue
			}

			xStart := x
			for ; x < len(line); x++ {
				if !unicode.IsNumber(rune(line[x])) {
					break
				}
			}

			number, err := strconv.Atoi(line[xStart:x])
			if err != nil {
				panic(err)
			}

			parts = append(parts, Number{
				Number: number,
				Length: x - xStart,
				Y:      y,
				X:      xStart,
			})
		}
	}

	return parts
}
