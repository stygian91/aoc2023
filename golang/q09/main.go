package main

import (
	"aoc2023/q09/part1"
	"aoc2023/utils/files"
)

func main() {
	// lines, err := files.ReadLines("./data/demo.txt")
	lines, err := files.ReadLines("./data/input.txt")
	if err != nil {
		panic(err)
	}

	part1.Part1(lines)
}
