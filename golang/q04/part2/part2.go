package part2

import (
	c "aoc2023/q04/common"
	"fmt"
)

func Part2(lines []string) {
	cards := []c.Card{}
	winCounts := []int{}
	cardCounts := []int{}

	for i, line := range lines {
		cards = append(cards, c.ParseLine(line))
		winCounts = append(winCounts, cards[i].CountWinning())
		cardCounts = append(cardCounts, 1)
	}

	totalCards := len(cards)
	for {
		newCards := 0

		for i := 0; i < len(cards); i++ {
			if cardCounts[i] == 0 {
				continue
			}

			for j := i + 1; j <= i+winCounts[i] && j < len(cards); j++ {
				cardCounts[j] += cardCounts[i]
			}

			newCards += cardCounts[i] * winCounts[i]

			cardCounts[i] = 0
		}

		totalCards += newCards

		if newCards == 0 {
			break
		}
	}

	fmt.Println("Part 2:", totalCards)
}
