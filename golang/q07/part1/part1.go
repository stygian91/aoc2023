package part1

import (
	c "aoc2023/q07/common"
	"fmt"
	"slices"
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
	Jack
	Ten
	Nine
	Eight
	Seven
	Six
	Five
	Four
	Three
	Two
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

func isFiveOfAKind(hand string) bool {
	first := hand[0]

	for i := 1; i < len(hand); i++ {
		ch := hand[i]
		if ch != first {
			return false
		}
	}

	return true
}

func isFourOfAKind(hand string) bool {
	cardCounts := c.GetCardCounts(hand)

	for _, count := range cardCounts {
		if count == 4 {
			return true
		}
	}

	return false
}

func isFullHouse(hand string) bool {
	hasTwo := false
	hasThree := false

	for _, count := range c.GetCardCounts(hand) {
		if count == 2 {
			hasTwo = true
		}

		if count == 3 {
			hasThree = true
		}
	}

	return hasTwo && hasThree
}

func isThreeOfAKind(hand string) bool {
	hasTwo := false
	hasThree := false

	for _, count := range c.GetCardCounts(hand) {
		if count == 2 {
			hasTwo = true
		}

		if count == 3 {
			hasThree = true
		}
	}

	return hasThree && !hasTwo
}

func isTwoPair(hand string) bool {
	twos := []rune{}

	for card, count := range c.GetCardCounts(hand) {
		if count == 2 && !slices.Contains(twos, card) {
			twos = append(twos, card)
		}
	}

	return len(twos) == 2
}

func isOnePair(hand string) bool {
	twos := []rune{}

	for card, count := range c.GetCardCounts(hand) {
		if count == 2 && !slices.Contains(twos, card) {
			twos = append(twos, card)
		}
	}

	return len(twos) == 1
}

func getHandType(hand string) HandType {
	if isFiveOfAKind(hand) {
		return FiveOfAKind
	} else if isFourOfAKind(hand) {
		return FourOfAKind
	} else if isFullHouse(hand) {
		return FullHouse
	} else if isThreeOfAKind(hand) {
		return ThreeOfAKind
	} else if isTwoPair(hand) {
		return TwoPair
	} else if isOnePair(hand) {
		return OnePair
	}

	return HighCard
}

func Part1(lines []string) {
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

	fmt.Println("Part 1:", sum)
}
