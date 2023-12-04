package part1

import (
	"fmt"
	"math"
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

func calcPoints(winCount int) int {
	if winCount < 1 {
		return 0
	}

	return int(math.Pow(float64(2), float64(winCount-1)))
}

func parseLine(line string) Card {
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

func Part1(lines []string) {
	total := 0

	for _, line := range lines {
		card := parseLine(line)
		winCount := card.CountWinning()
		total += calcPoints(winCount)
	}

	fmt.Println("Part 1:", total)
}
