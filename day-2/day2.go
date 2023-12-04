package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	ID    int
	Red   []int
	Blue  []int
	Green []int
}

// func (g Game) SumRed() int {
// 	var num int
// 	for _, v := range g.Red {
// 		num += v
// 	}
// 	return num
// }
//
// func (g Game) SumBlue() int {
// 	var num int
// 	for _, v := range g.Blue {
// 		num += v
// 	}
// 	return num
// }
//
// func (g Game) SumGreen() int {
// 	var num int
// 	for _, v := range g.Green {
// 		num += v
// 	}
// 	return num
// }

func ParseGames(lines []string) []Game {
	var games []Game
	for _, line := range lines {
		var game Game

		col := strings.Split(line, ":")
		id, err := strconv.Atoi(col[0][strings.Index(col[0], " ")+1 : len(col[0])])
		if err != nil {
			panic(err)
		}
		game.ID = id

		sets := strings.Split(col[1], ";")
		for _, set := range sets {

			var red, blue, green int
			for _, s := range strings.Split(set, ",") {
				s := strings.TrimSpace(s)

				num, err := strconv.Atoi(s[0:strings.Index(s, " ")])
				if err != nil {
					panic(err)
				}

				if strings.HasSuffix(s, "red") {
					red = num
				} else if strings.HasSuffix(s, "blue") {
					blue = num
				} else if strings.HasSuffix(s, "green") {
					green = num
				} else {
					panic(fmt.Sprintf("invalid input: %s", s))
				}
			}

			game.Red = append(game.Red, red)
			game.Blue = append(game.Blue, blue)
			game.Green = append(game.Green, green)
		}
		games = append(games, game)
	}
	return games
}
