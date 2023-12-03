package part1

import (
	"aoc2023/q03/common"
	"fmt"
)

func Part1(schematic common.Schematic) {
	sum := 0

	for _, part := range schematic.Numbers {
		neighbours := schematic.NeighborSymbols(part)
		if len(neighbours) > 0 {
			sum += part.Value
		}
	}

	fmt.Println("Part 1:", sum)
}
