// Package util holds random things that are useful and probably involve generics.
package util

import "golang.org/x/exp/constraints"

// Min returns the smaller of two values.
func Min[T constraints.Ordered](v1, v2 T) T {
	if v1 < v2 {
		return v1
	}
	return v2
}

// Max returns the larger of two values.
func Max[T constraints.Ordered](v1, v2 T) T {
	if v1 > v2 {
		return v1
	}
	return v2
}

// Contain returns v unless it is outside of the range [min:max], in
// which case it returns the nearer of those two values.
func Contain[T constraints.Ordered](min, max, v T) T {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

// LastIndex returns the last index of v in s, or -1 if v is not in s.
func LastIndex[S ~[]E, E comparable](s S, v E) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == v {
			return i
		}
	}
	return -1
}

// MaxIndex returns the first index of v in s, or len(s) if v is not
// in s.
func MaxIndex[S ~[]E, E comparable](s S, v E) int {
	for i, e := range s {
		if e == v {
			return i
		}
	}
	return len(s)
}
