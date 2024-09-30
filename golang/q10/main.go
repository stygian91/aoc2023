package main

import (
	"aoc2023/q10/common"
	// "aoc2023/q10/part1"
	"aoc2023/q10/part2"
	"aoc2023/utils/files"
)

func main() {
	// lines, err := files.ReadLines("./data/demo.txt")
	// lines, err := files.ReadLines("./data/demo2.txt")
	lines, err := files.ReadLines("./data/input.txt")
	if err != nil {
		panic(err)
	}

	board := common.Parse(lines)
	part2.Part2(board)
}
