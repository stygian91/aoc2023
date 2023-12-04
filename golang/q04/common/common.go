package common

import (
	"strconv"
	"strings"
)

type Card struct {
	Winning, Guess []int
}

func (this Card) CountWinning() int {
	count := 0

	for _, win := range this.Winning {
		for _, guess := range this.Guess {
			if win == guess {
				count++
				break
			}
		}
	}

	return count
}

func ParseLine(line string) Card {
	colonIdx := strings.Index(line, ":")
	if colonIdx == -1 {
		panic("Colon not found while parsing")
	}

	remainder := line[colonIdx+2:]
	pipeIdx := strings.Index(remainder, "|")
	if pipeIdx == -1 {
		panic("Pipe not found while parsing")
	}

	winStr := remainder[:pipeIdx]
	guessStr := remainder[pipeIdx+2:]

	winNumbers := parseNumberTable(winStr)
	guessNumbers := parseNumberTable(guessStr)

	return Card{Winning: winNumbers, Guess: guessNumbers}
}

func parseNumberTable(table string) []int {
	result := []int{}

	for _, field := range strings.Fields(table) {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic("Error while parsing number: " + field)
		}

		result = append(result, num)
	}

	return result
}
