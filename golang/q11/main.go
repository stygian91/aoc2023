package main

import (
	"aoc2023/q11/common"
	"aoc2023/q11/part1"
)

func main() {
	// path := "./data/demo.txt"
	path := "./data/input.txt"
	universe, err := common.Parse(path)
	if err != nil {
		panic(err)
	}

	part1.Part1(&universe)
}
