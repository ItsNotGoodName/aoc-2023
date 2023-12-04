package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func replacements() []string {
	var numbers []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty", "twenty-one", "twenty-two", "twenty-three", "twenty-four", "twenty-five", "twenty-six", "twenty-seven", "twenty-eight", "twenty-nine", "thirty", "thirty-one", "thirty-two", "thirty-three", "thirty-four", "thirty-five", "thirty-six", "thirty-seven", "thirty-eight", "thirty-nine", "forty", "forty-one", "forty-two", "forty-three", "forty-four", "forty-five", "forty-six", "forty-seven", "forty-eight", "forty-nine", "fifty", "fifty-one", "fifty-two", "fifty-three", "fifty-four", "fifty-five", "fifty-six", "fifty-seven", "fifty-eight", "fifty-nine", "sixty", "sixty-one", "sixty-two", "sixty-three", "sixty-four", "sixty-five", "sixty-six", "sixty-seven", "sixty-eight", "sixty-nine", "seventy", "seventy-one", "seventy-two", "seventy-three", "seventy-four", "seventy-five", "seventy-six", "seventy-seven", "seventy-eight", "seventy-nine", "eighty", "eighty-one", "eighty-two", "eighty-three", "eighty-four", "eighty-five", "eighty-six", "eighty-seven", "eighty-eight", "eighty-nine", "ninety", "ninety-one", "ninety-two", "ninety-three", "ninety-four", "ninety-five", "ninety-six", "ninety-seven", "ninety-eight", "ninety-nine", "hundred"}
	var replacements []string
	for i := range numbers {
		replacements = append(replacements, numbers[i], numbers[i]+strconv.Itoa(i+1)+numbers[i])
	}
	return replacements
}

func main() {
	raw, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(raw)), "\n")

	var sum int

	replacements := replacements()

	for _, prevLine := range lines {
		line := prevLine
		for i := 0; i < len(replacements); i += 2 {
			line = strings.ReplaceAll(line, replacements[i], replacements[i+1])
		}

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
