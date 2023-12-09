package part2

import (
	"aoc2023/q09/common"
	"fmt"
)

func extra(derives [][]int) int {
	currentE := 0

	for i := len(derives) - 2; i >= 0; i-- {
		currentD := derives[i]
		first := currentD[0]

		currentE = first - currentE
	}

	return currentE
}

func Part2(lines []string) {
	nums := common.Parse(lines)
	sum := 0

	for _, entry := range nums {
		sum += extra(common.DeriveAll(entry))
	}

	fmt.Println("Part 2:", sum)
}
