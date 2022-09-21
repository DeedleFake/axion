// Package util holds random things that are useful and probably involve generics.
package util

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](v1, v2 T) T {
	if v1 < v2 {
		return v1
	}
	return v2
}

func Max[T constraints.Ordered](v1, v2 T) T {
	if v1 > v2 {
		return v1
	}
	return v2
}

func Contain[T constraints.Ordered](min, max, v T) T {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
