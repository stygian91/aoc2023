package part2

import (
	c "aoc2023/q02/common"
	"fmt"
)

func Part2(games []c.Game) {
	sum := 0

	for _, game := range games {
		power := 1
		maxColors := game.MaxColors()

		power = power * maxColors.Red
		power = power * maxColors.Green
		power = power * maxColors.Blue

		sum += power
	}

	fmt.Println("Part 2:", sum)
}
