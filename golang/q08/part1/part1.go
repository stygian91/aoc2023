package part1

import (
	"aoc2023/q08/common"
	"fmt"
)

func Part1(lines []string) {
	instructions := common.ParseLR(lines[0])
	nodes := common.ParseNodes(lines[2:])

	instructionsTaken := 0
	instructionIdx := 0
	nodeKey := "AAA"

	for {
		instruction := instructions[instructionIdx]
		currentNode := nodes.Mapping[nodeKey]

		if instruction == common.Left {
			nodeKey = currentNode[0]
		} else {
			nodeKey = currentNode[1]
		}

		instructionsTaken++
		instructionIdx++
		if instructionIdx == len(instructions) {
			instructionIdx = 0
		}

		if nodeKey == "ZZZ" {
			break
		}
	}

	fmt.Println("Part 1:", instructionsTaken)
}
