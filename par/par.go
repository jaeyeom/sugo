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
