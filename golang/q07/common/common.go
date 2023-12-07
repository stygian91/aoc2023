package common

func GetCardCounts(hand string) map[rune]int {
	cardCounts := map[rune]int{}

	for _, card := range hand {
		_, exists := cardCounts[card]
		if exists {
			cardCounts[card]++
		} else {
			cardCounts[card] = 1
		}
	}

	return cardCounts
}
