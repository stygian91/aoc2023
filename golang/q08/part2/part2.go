package part2

import (
	c "aoc2023/q08/common"
	"aoc2023/utils/slices"
	"fmt"
)

func findStartNodes(nodes c.Nodes) []string {
	result := []string{}

	for key := range nodes.Mapping {
		if key[2] == 'A' {
			result = append(result, key)
		}
	}

	return result
}

func smallestCommonMultiple(a, b uint) uint {
	runA, runB := a, b

	for {
		if runA == runB {
			break
		}

		if runA < runB {
			runA += a
		} else {
			runB += b
		}
	}

	return runA
}

func findFirstZCount(start string, nodes c.Nodes, instructions []c.LR) int {
	instructionIdx := 0
	instructionsTaken := 0
	current := start

	for {
		instruction := instructions[instructionIdx]

		var dirIdx int
		if instruction == c.Left {
			dirIdx = 0
		} else {
			dirIdx = 1
		}

		current = nodes.Mapping[current][dirIdx]
		instructionsTaken++
		instructionIdx++
		if instructionIdx == len(instructions) {
			instructionIdx = 0
		}

		if current[2] == 'Z' {
			break
		}
	}

	return instructionsTaken
}

func Part2(lines []string) {
	instructions := c.ParseLR(lines[0])
	nodes := c.ParseNodes(lines[2:])
	currNodes := findStartNodes(nodes)

	firstZs := slices.Map(func(_ int, node string) int {
		return findFirstZCount(node, nodes, instructions)
	}, currNodes)

	multi := uint(1)

	for _, firstZ := range firstZs {
		multi = smallestCommonMultiple(multi, uint(firstZ))
	}

	fmt.Println("Part 2:", multi)
}
