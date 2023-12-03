package main

import (
	"aoc2023/q03/common"
	"aoc2023/q03/part1"
	"aoc2023/q03/part2"
	"aoc2023/utils/files"
)

func main() {
	// lines, err := files.ReadLines("./data/test1.txt")
	// lines, err := files.ReadLines("./data/demo.txt")
	lines, err := files.ReadLines("./data/input.txt")
	if err != nil {
		panic(err)
	}

	schematic := common.Parse(lines)
	part1.Part1(schematic)
	part2.Part2(schematic)
}
