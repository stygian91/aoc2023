package math

import "golang.org/x/exp/constraints"

type Vec2 struct {
	X, Y int
}

func (this Vec2) Add(other Vec2) Vec2 {
	return Vec2{X: this.X + other.X, Y: this.Y + other.Y}
}

func IsInBound[N constraints.Ordered](x, lower, upper N) bool {
	return lower <= x && upper >= x
}
