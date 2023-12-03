package part2

import (
	"aoc2023/utils/math"
	"fmt"
	"strconv"
	"strings"
)

// the data does not contain any zeros
var numbers []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

func findFirstSpeltNumber(haystack string) (int, int) {
	for i, number := range numbers {
		foundIdx := strings.Index(haystack, number)
		if foundIdx != -1 {
			return foundIdx, i + 1
		}
	}

	return -1, 0
}

func firstDigit(line string) int {
	work := ""

	for _, rune := range line {
		if math.IsDigit(rune) {
			digit, err := strconv.ParseInt(string(rune), 10, 64)

			if err != nil {
				panic(err)
			}

			return int(digit)
		}

		work += string(rune)
		idx, digit := findFirstSpeltNumber(work)
		if idx == -1 {
			continue
		}

		if digit > 9 || digit < 1 {
			panic("invalid digit")
		}

		return digit
	}

	panic("No digit found")
}

func lastDigit(line string) int {
	work := ""

	for i := len(line) - 1; i >= 0; i-- {
		ch := line[i]
		rune := rune(ch)

		if math.IsDigit(rune) {
			digit, err := strconv.ParseInt(string(rune), 10, 64)

			if err != nil {
				panic(err)
			}

			return int(digit)
		}

		work = string(rune) + work
		idx, digit := findFirstSpeltNumber(work)
		if idx == -1 {
			continue
		}

		if digit > 9 || digit < 1 {
			panic("invalid digit")
		}

		return digit
	}

	panic("No digit found")
}

func processLine(line string) int {
	return 10*firstDigit(line) + lastDigit(line)
}

func Part2(lines []string) {
	sum := 0

	for _, line := range lines {
		num := processLine(line)
		sum += num
	}

	fmt.Println("Part 2:", sum)
}
