package part1

import (
	"aoc2023/q09/common"
	"fmt"
)

func extra(derives [][]int) int {
	currentE := 0

	for i := len(derives) - 2; i >= 0; i-- {
		currentD := derives[i]
		last := currentD[len(currentD)-1]
		currentE += last
	}

	return currentE
}

func Part1(lines []string) {
	nums := common.Parse(lines)
	sum := 0

	for _, entry := range nums {
		sum += extra(common.DeriveAll(entry))
	}

	fmt.Println("Part 1:", sum)
}
