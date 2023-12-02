package part1

import (
	c "aoc2023/q02/common"
	"fmt"
)

var maxCubes c.Round = c.Round{Red: 12, Green: 13, Blue: 14}

func Part1(games []c.Game) {
	sum := 0

	for i, game := range games {
		gameMax := game.MaxColors()

		if gameMax.Red > maxCubes.Red || gameMax.Green > maxCubes.Green || gameMax.Blue > maxCubes.Blue {
			continue
		}

		sum += i + 1
	}

	fmt.Println("Part 1:", sum)
}
