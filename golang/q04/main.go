package main

import (
	"aoc2023/q04/part1"
	"aoc2023/q04/part2"
	"aoc2023/utils/files"
)

func main() {
	lines, err := files.ReadLines("./data/input.txt")
	if err != nil {
		panic(err)
	}

	part1.Part1(lines)
	part2.Part2(lines)
}
