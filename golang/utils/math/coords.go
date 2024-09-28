package math

import "golang.org/x/exp/constraints"

type Vec2 struct {
	X, Y int
}

func IsInBound[N constraints.Ordered](x, lower, upper N) bool {
	return lower <= x && upper >= x
}
