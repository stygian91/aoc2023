package part2

import (
	c "aoc2023/q03/common"
	m "aoc2023/utils/math"
	"fmt"
)

func isNeighbour(part c.PartNumber, pos m.Vec2) bool {
	for y := pos.Y - 1; y <= pos.Y + 1; y++ {
		if y != part.NumberStart.Y {
			continue
		}

		for x := pos.X-1; x <= pos.X+1; x++ {
			if m.IsInBound(x, part.NumberStart.X, part.NumberEnd.X) {
				return true
			}
		}
	}

	return false
}

func Part2(schematic c.Schematic) {
	sum := 0

	for pos, symbol := range schematic.Symbols {
		if symbol != "*" {
			continue
		}

		neighbours := []c.PartNumber{}

		for _, part := range schematic.Numbers {
			if isNeighbour(part, pos) {
				neighbours = append(neighbours, part)
			}
		}

		if len(neighbours) != 2 {
			continue
		}

		power := neighbours[0].Value * neighbours[1].Value
		sum += power
	}

	fmt.Println("Part 2:", sum)
}
