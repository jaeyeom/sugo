// Package arg implements generic functions for ArgMin/ArgMax.
package arg

import "cmp"

// Min returns an index of the minimum value in data.
func Min(size int, less func(i, j int) bool) int {
	return seqMin(size, less)
}

// Max returns an index of the maximum value in data.
func Max(size int, less func(i, j int) bool) int {
	return Min(size, func(i, j int) bool {
		return less(j, i)
	})
}

// MinOf returns an index of the minimum value in s.
func MinOf[T cmp.Ordered](s []T) int {
	return Min(len(s), func(i, j int) bool {
		return s[i] < s[j]
	})
}

// MaxOf returns an index of the maximum value in s.
func MaxOf[T cmp.Ordered](s []T) int {
	return Max(len(s), func(i, j int) bool {
		return s[i] < s[j]
	})
}

// MinFunc returns an index of the minimum value in s using less.
func MinFunc[T any](s []T, less func(a, b T) bool) int {
	return Min(len(s), func(i, j int) bool {
		return less(s[i], s[j])
	})
}

// MaxFunc returns an index of the maximum value in s using less.
func MaxFunc[T any](s []T, less func(a, b T) bool) int {
	return Max(len(s), func(i, j int) bool {
		return less(s[i], s[j])
	})
}

// seqMin finds Min value sequentially.
func seqMin(size int, less func(i, j int) bool) int {
	if size <= 0 {
		panic("cannot get minimum value in an empty data set")
	}
	Min := 0
	for i := 1; i < size; i++ {
		if less(i, Min) {
			Min = i
		}
	}
	return Min
}
