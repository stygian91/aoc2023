package part2

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

func Part2(board c.Board) {
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

	slices.Reverse(st.visited[1])
	vertices := slices.Concat(st.visited[0][1:len(st.visited[0])-1], st.visited[1])
	area := shoelace(vertices)
	answer := area - float64(len(vertices) / 2) + 1
	fmt.Printf("answer: %f\n", answer)
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
	otherLastPos := st.visited[otherIdx][len(st.visited[otherIdx])-1]
	if otherLastPos == unvisited[0].Pos {
		st.done = true
	}
}

func shoelace(xs []m.Vec2) float64 {
	sum1 := 0
	sum2 := 0
	for i := 0; i < len(xs); i++ {
		j := i + 1
		if j == len(xs) {
			j = 0
		}

		sum1 += xs[i].X * xs[j].Y
		sum2 += xs[i].Y * xs[j].X
	}

	return float64(sum1-sum2) / 2
}
