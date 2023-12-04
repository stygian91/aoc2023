package part1

import (
	"aoc2023/q04/common"
	"fmt"
	"math"
)

func calcPoints(winCount int) int {
	if winCount < 1 {
		return 0
	}

	return int(math.Pow(float64(2), float64(winCount-1)))
}

func Part1(lines []string) {
	total := 0

	for _, line := range lines {
		card := common.ParseLine(line)
		winCount := card.CountWinning()
		total += calcPoints(winCount)
	}

	fmt.Println("Part 1:", total)
}
