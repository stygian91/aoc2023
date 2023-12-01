package part1

import (
	"aoc2023/q01/common"
	"fmt"
	"strconv"
)

func Part1(lines []string) {
	var sum int64 = 0
	for _, line := range lines {
		digits := ""

		for _, c := range line {
			if common.IsDigit(c) {
				digits = digits + string(c)
			}
		}

		first := digits[0]
		var combined string
		if len(digits) > 1 {
			combined = string(first) + string(digits[len(digits)-1])
		} else {
			combined = string(first) + string(first)
		}

		num, e := strconv.ParseInt(combined, 10, 64)
		if e != nil {
			panic(e)
		}

		sum += num
	}

	fmt.Println("Part 1:", sum)
}
