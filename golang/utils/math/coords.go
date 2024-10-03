package math

import (
	"golang.org/x/exp/constraints"
)

type Vec2 struct {
	X, Y int
}

func (this Vec2) Add(other Vec2) Vec2 {
	return Vec2{X: this.X + other.X, Y: this.Y + other.Y}
}

func (this Vec2) ManhattanDistance(other Vec2) int {
	res := 0

	res += absInt(this.X - other.X)
	res += absInt(this.Y - other.Y)

	return res
}

func IsInBound[N constraints.Ordered](x, lower, upper N) bool {
	return lower <= x && upper >= x
}

func absInt(a int) int {
	if a < 0 {
		return a * -1
	}

	return a
}
