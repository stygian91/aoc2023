package part1

import (
	c "aoc2023/q10/common"
	m "aoc2023/utils/math"
	"fmt"
)

type state struct {
	visited [][]m.Vec2
	done    bool
}

func Part1(board c.Board) {
	starting := board.StartingPipes()
	if len(starting) != 2 {
		panic("Expected 2 starting pipes")
	}

	startingTile1, exists := board.GetTile(starting[0])
	if !exists {
		panic("starting1 doesn't exist")
	}

	st := state{
		visited: [][]m.Vec2{
			{board.Start, starting[0]},
			{board.Start, starting[1]},
		},
		done: false,
	}

	for !st.done {
		run(0, &st)
		run(1, &st)
	}
}

func run(idx int, st *state) {
}
