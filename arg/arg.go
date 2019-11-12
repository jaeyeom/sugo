// Package arg implements generic functions for ArgMin/ArgMax.
package arg

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
