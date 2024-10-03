package common

import (
	"aoc2023/utils/files"
	m "aoc2023/utils/math"
	"fmt"
	"slices"
	"strings"
)

type Universe struct {
	Galaxies []m.Vec2
	Size     m.Vec2
}

type empty struct {
	Rows, Cols []int
}


func (this Universe) String() string {
	b := strings.Builder{}
	curr := 1

	for i := 0; i < this.Size.Y; i++ {
		for j := 0; j < this.Size.X; j++ {
			if this.IsGalaxy(m.Vec2{X: j, Y: i}) {
				b.WriteString(fmt.Sprintf("%d", curr))
				curr++
			} else {
				b.WriteRune('.')
			}
		}

		b.WriteRune('\n')
	}

	b.WriteRune('\n')
	return b.String()
}

func (this Universe) IsGalaxy(pos m.Vec2) bool {
	for _, g := range this.Galaxies {
		if pos == g {
			return true
		}
	}

	return false
}

func Parse(path string) (Universe, error) {
	universe := Universe{}

	lines, err := files.ReadLines(path)
	if err != nil {
		return universe, err
	}
	rowCount := len(lines)
	colCount := 0
	first := true

	galaxies := []m.Vec2{}

	for i, line := range lines {
		for j, r := range line {
			switch r {
			case '#':
				galaxies = append(galaxies, m.Vec2{X: j, Y: i})
			case '.':

			default:
				return universe, fmt.Errorf("Unexpected rune %c", r)
			}

			if first {
				colCount++
			}
		}

		first = false
	}

	universe.Galaxies = galaxies
	universe.Size.X = colCount
	universe.Size.Y = rowCount

	return universe, nil
}

func findEmpty(universe *Universe) empty {
	res := empty{}

	emptyCols := slices.Repeat([]bool{true}, universe.Size.X)
	emptyRows := slices.Repeat([]bool{true}, universe.Size.Y)

	for _, pos := range universe.Galaxies {
		emptyCols[pos.X] = false
		emptyRows[pos.Y] = false
	}

	var l int
	if universe.Size.X > universe.Size.Y {
		l = universe.Size.X
	} else {
		l = universe.Size.Y
	}

	for i := 0; i < l; i++ {
		if i < universe.Size.X && emptyCols[i] {
			res.Cols = append(res.Cols, i)
		}

		if i < universe.Size.Y && emptyRows[i] {
			res.Rows = append(res.Rows, i)
		}
	}

	return res
}

func Expand(universe *Universe, expandRate int) {
	diffs := make([]m.Vec2, len(universe.Galaxies))
	for i := 0; i < len(universe.Galaxies); i++ {
		diffs[i] = m.Vec2{X: 0, Y: 0}
	}

	e := findEmpty(universe)

	for _, r := range e.Rows {
		for i := range diffs {
			if universe.Galaxies[i].Y > r {
				diffs[i].Y += expandRate
			}
		}
	}

	for _, c := range e.Cols {
		for i := range diffs {
			if universe.Galaxies[i].X > c {
				diffs[i].X += expandRate
			}
		}
	}

	for i, d := range diffs {
		universe.Galaxies[i].X += d.X
		universe.Galaxies[i].Y += d.Y
	}

	universe.Size.X += len(e.Cols)
	universe.Size.Y += len(e.Rows)
}
