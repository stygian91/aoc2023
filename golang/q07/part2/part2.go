package part2

import (
	c "aoc2023/q07/common"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type HandType int
type CardType int

const (
	FiveOfAKind HandType = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

const (
	Ace CardType = iota
	King
	Queen
	Ten
	Nine
	Eight
	Seven
	Six
	Five
	Four
	Three
	Two
	Jack
)

type Hand struct {
	Cards []CardType
	Type  HandType
	Bid   int
}

func (this Hand) Compare(other Hand) int {
	if this.Type < other.Type {
		return 1
	}

	if this.Type > other.Type {
		return -1
	}

	for i := 0; i < 5; i++ {
		thisCard := this.Cards[i]
		otherCard := other.Cards[i]
		comp := thisCard.Compare(otherCard)

		if comp < 0 {
			return -1
		}

		if comp > 0 {
			return 1
		}
	}

	return 0
}

func ParseHand(line string) Hand {
	parts := strings.Split(line, " ")
	hand := parts[0]
	bid, err := strconv.Atoi(parts[1])
	if err != nil {
		panic("Invalid bid")
	}

	cards := []CardType{}
	for _, card := range hand {
		cards = append(cards, ParseCard(card))
	}

	return Hand{
		Cards: cards,
		Type:  getHandType(hand),
		Bid:   bid,
	}
}

func (this CardType) Compare(other CardType) int {
	return int(other) - int(this)
}

func ParseCard(card rune) CardType {
	switch card {
	case 'A':
		return Ace
	case 'K':
		return King
	case 'Q':
		return Queen
	case 'J':
		return Jack
	case 'T':
		return Ten
	case '9':
		return Nine
	case '8':
		return Eight
	case '7':
		return Seven
	case '6':
		return Six
	case '5':
		return Five
	case '4':
		return Four
	case '3':
		return Three
	case '2':
		return Two
	default:
		panic("Invalid card")
	}
}

func isFiveOfAKind(cardCounts map[rune]int) bool {
	cardLen := len(cardCounts)
	if cardLen == 1 {
		return true
	}

	jokerCount := cardCounts['J']
	if cardLen != 2 {
		return false
	}

	otherCount := 0
	for card, count := range cardCounts {
		if card == 'J' {
			continue
		}

		otherCount = count
	}

	return otherCount+jokerCount == 5
}

func isFourOfAKind(cardCounts map[rune]int) bool {
	for _, count := range cardCounts {
		if count == 4 {
			return true
		}
	}

	hasOne := false
	hasTwo := false
	hasThree := false

	for card, count := range cardCounts {
		if card == 'J' {
			continue
		}

		switch count {
		case 1:
			hasOne = true
		case 2:
			hasTwo = true
		case 3:
			hasThree = true
		}
	}

	return (hasOne && cardCounts['J'] == 3) || (hasTwo && cardCounts['J'] == 2) || (hasThree && cardCounts['J'] == 1)
}

func isFullHouse(counts map[rune]int) bool {
	twos := []rune{}
	hasThree := false

	for card, count := range counts {
		if card == 'J' {
			continue
		}

		switch count {
		case 2:
			twos = append(twos, card)
		case 3:
			hasThree = true
		}
	}

	if len(twos) > 0 && hasThree {
		return true
	}

	return len(twos) == 2 && counts['J'] == 1
}

func isThreeOfAKind(counts map[rune]int) bool {
	ones := []rune{}
	twos := []rune{}
	hasThree := false

	for card, count := range counts {
		if card == 'J' {
			continue
		}

		switch count {
		case 1:
			ones = append(ones, card)
		case 2:
			twos = append(twos, card)
		case 3:
			hasThree = true
		}

	}

	if hasThree && len(twos) == 0 && counts['J'] == 0 {
		return true
	}

	return (len(twos) == 1 && len(ones) == 2 && counts['J'] == 1) || (len(ones) == 3 && counts['J'] == 2)
}

func isTwoPair(counts map[rune]int) bool {
	ones := []rune{}
	twos := []rune{}

	for card, count := range counts {
		if card == 'J' {
			continue
		}

		switch count {
		case 1:
			ones = append(ones, card)
		case 2:
			twos = append(twos, card)
		}
	}

	return len(twos) == 2 && counts['J'] == 0
}

func isOnePair(counts map[rune]int) bool {
	ones := []rune{}
	twos := []rune{}

	for card, count := range counts {
		if card == 'J' {
			continue
		}

		switch count {
		case 1:
			ones = append(ones, card)
		case 2:
			twos = append(twos, card)
		}
	}

	return (len(twos) == 1 && counts['J'] == 0) || (len(ones) == 4 && counts['J'] == 1)
}

func getHandType(hand string) HandType {
	cardCounts := c.GetCardCounts(hand)

	if isFiveOfAKind(cardCounts) {
		return FiveOfAKind
	} else if isFourOfAKind(cardCounts) {
		return FourOfAKind
	} else if isFullHouse(cardCounts) {
		return FullHouse
	} else if isThreeOfAKind(cardCounts) {
		return ThreeOfAKind
	} else if isTwoPair(cardCounts) {
		return TwoPair
	} else if isOnePair(cardCounts) {
		return OnePair
	}

	return HighCard
}

func Part2(lines []string) {
	hands := []Hand{}

	for _, line := range lines {
		hands = append(hands, ParseHand(line))
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Compare(hands[j]) < 0
	})

	sum := 0

	for i, hand := range hands {
		rank := i + 1
		points := hand.Bid * rank
		sum += points
	}

	fmt.Println("Part 2:", sum)
}
