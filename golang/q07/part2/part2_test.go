package part2

import "testing"

type isKindFn func(map[rune]int) bool

// Helpers --------
func assertKind(t *testing.T, fn isKindFn, cardCounts map[rune]int) {
	if !fn(cardCounts) {
		t.Errorf("assertKind: %+v", cardCounts)
	}
}

func assertNotKind(t *testing.T, fn isKindFn, cardCounts map[rune]int) {
	if fn(cardCounts) {
		t.Errorf("assertNotKind: %+v", cardCounts)
	}
}

// Tests --------
func TestIsFiveOfAKind(t *testing.T) {
	assertFive := func(counts map[rune]int) {
		assertKind(t, isFiveOfAKind, counts)
	}

	assertFive(map[rune]int{'J': 1, 'K': 4})
	assertFive(map[rune]int{'J': 2, 'Q': 3})
	assertFive(map[rune]int{'J': 3, 'T': 2})
	assertFive(map[rune]int{'J': 4, '9': 1})
	assertFive(map[rune]int{'J': 5})
	assertFive(map[rune]int{'K': 5})

	assertNotKind(t, isFiveOfAKind, map[rune]int{'J': 1, 'K': 3, 'Q': 1})
	assertNotKind(t, isFiveOfAKind, map[rune]int{'J': 0, 'K': 4, 'Q': 1})
	assertNotKind(t, isFiveOfAKind, map[rune]int{'J': 0, 'K': 3, 'Q': 2})
}

func TestIsFourOfAKind(t *testing.T) {
	assertFour := func(counts map[rune]int) {
		assertKind(t, isFourOfAKind, counts)
	}

	assertNotFour := func(counts map[rune]int) {
		assertNotKind(t, isFourOfAKind, counts)
	}

	assertFour(map[rune]int{'K': 4, 'Q': 1})
	assertFour(map[rune]int{'K': 3, 'Q': 1, 'J': 1})
	assertFour(map[rune]int{'K': 2, 'Q': 1, 'J': 2})
	assertFour(map[rune]int{'K': 1, 'Q': 1, 'J': 3})

	assertNotFour(map[rune]int{'K': 3, 'Q': 2})
	assertNotFour(map[rune]int{'K': 2, 'Q': 2, 'J': 1})
	assertNotFour(map[rune]int{'K': 1, 'Q': 1, '9': 1, 'J': 2})
}

func TestIsFullHouse(t *testing.T) {
	assertKind(t, isFullHouse, map[rune]int{'K': 3, 'Q': 2})
	assertKind(t, isFullHouse, map[rune]int{'K': 2, 'Q': 2, 'J': 1})
}

func TestIsThree(t *testing.T) {
	assertKind(t, isThreeOfAKind, map[rune]int{'K': 3, 'Q': 1, 'T': 1})
	assertKind(t, isThreeOfAKind, map[rune]int{'K': 2, 'Q': 1, 'T': 1, 'J': 1})
	assertKind(t, isThreeOfAKind, map[rune]int{'K': 1, 'Q': 1, 'T': 1, 'J': 2})
}

func TestIsTwo(t *testing.T) {
	assertKind(t, isTwoPair, map[rune]int{'K': 2, 'Q': 2, 'T': 1})
}

func TestIsOne(t *testing.T) {
	assertKind(t, isOnePair, map[rune]int{'K': 2, 'Q': 1, 'T': 1, '9': 1})
	assertKind(t, isOnePair, map[rune]int{'K': 1, 'Q': 1, 'T': 1, '9': 1, 'J': 1})
}
