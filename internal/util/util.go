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

func LastIndex[S ~[]E, E comparable](s S, v E) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == v {
			return i
		}
	}
	return -1
}

func MaxIndex[S ~[]E, E comparable](s S, v E) int {
	for i, e := range s {
		if e == v {
			return i
		}
	}
	return len(s)
}
