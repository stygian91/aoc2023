package part1

import (
	c "aoc2023/q10/common"
	m "aoc2023/utils/math"
	"fmt"
	"slices"
)

type state struct {
	visited [][]m.Vec2
	done    bool
	board   c.Board
}

func Part1(board c.Board) {
	starting := board.StartingPipes()
	if len(starting) != 2 {
		panic("Expected 2 starting pipes")
	}

	st := state{
		visited: [][]m.Vec2{
			{board.Start, starting[0]},
			{board.Start, starting[1]},
		},
		done:  false,
		board: board,
	}

	for !st.done {
		run(0, &st)
		run(1, &st)
	}

	fmt.Printf("answer: %d\n", len(st.visited[0]) - 1)
}

func run(idx int, st *state) {
	lastPos := st.visited[idx][len(st.visited[idx])-1]

	neighbours := st.board.GetNeighbours(lastPos)
	unvisited := []c.TilePos{}

	for _, n := range neighbours {
		if slices.Contains(st.visited[idx], n.Pos) {
			continue
		}

		unvisited = append(unvisited, n)
	}

	if len(unvisited) != 1 {
		panic("no unvisited found")
	}

	var otherIdx int
	if idx == 0 {
		otherIdx = 1
	} else {
		otherIdx = 0
	}

	st.visited[idx] = append(st.visited[idx], unvisited[0].Pos)
	otherLastPos := st.visited[otherIdx][len(st.visited[otherIdx]) - 1]
	if otherLastPos == unvisited[0].Pos {
		st.done = true
	}
}
