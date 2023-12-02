package main

import (
	"aoc2023/q02/common"
	"aoc2023/q02/part1"
	"aoc2023/q02/part2"
	"aoc2023/utils/files"
)

func main() {
	lines, err := files.ReadLines("./data/input.txt")
	if err != nil {
		panic(err)
	}

	games := common.ParseGames(lines)
	part1.Part1(games)
	part2.Part2(games)
}
