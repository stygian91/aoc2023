package part1

import (
	c "aoc2023/q11/common"
	"fmt"
)

func Part1(universe *c.Universe) {
	// part 1:
	// c.Expand(universe, 1)
	// part 2:
	c.Expand(universe, 999_999)
	// fmt.Println(universe)

	sum := 0

	for i := 0; i < len(universe.Galaxies)-1; i++ {
		for j := i + 1; j < len(universe.Galaxies); j++ {
			dist := universe.Galaxies[i].ManhattanDistance(universe.Galaxies[j])
			sum += dist
		}
	}

	fmt.Printf("Answer: %d\n", sum)
}
