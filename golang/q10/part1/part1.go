package part1

import (
	c "aoc2023/q10/common"
	"aoc2023/utils/math"
	"fmt"
)

type state struct {
	visited []math.Vec2
}

func Part1(board c.Board) {
	starting := board.StartingPipes()
	if len(starting) != 2 {
		panic("Expected 2 starting pipes")
	}

	fmt.Printf("starting: %#+v", starting)
}

