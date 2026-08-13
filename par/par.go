// Package par provides convenient functions for concurrency and parallelism.
package par

import "sync"

// For calls f(i) in n goroutines where i is from 0 to n-1. The function blocks
// until all goroutines finishes. It is recommended for dealing with multiple
// objects/elements with same code.
func For(n int, f func(i int)) {
	if n <= 0 {
		return
	}
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			f(i)
		}(i)
	}
	wg.Wait()
}

// Do calls the given functions in each goroutine and waits for all goroutines
// finishes. It is recommended for running multiple different codes.
func Do(fs ...func()) {
	var wg sync.WaitGroup
	wg.Add(len(fs))
	for _, f := range fs {
		go func(f func()) {
			defer wg.Done()
			f()
		}(f)
	}
	wg.Wait()
}

// ForEach calls f(v) concurrently for each element of s. The function blocks
// until all goroutines finish.
func ForEach[T any](s []T, f func(T)) {
	For(len(s), func(i int) {
		f(s[i])
	})
}

// Map applies f to each element of s concurrently and returns a slice of the
// results in the same order. The function blocks until all goroutines finish.
func Map[T, U any](s []T, f func(T) U) []U {
	out := make([]U, len(s))
	For(len(s), func(i int) {
		out[i] = f(s[i])
	})
	return out
}
