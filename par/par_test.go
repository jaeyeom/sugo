package par

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/jaeyeom/sugo/ranger"
)

func ExampleFor_rowSum() {
	mat := [][]int{
		{49, 30, 20, 25},
		{22, 11, 85, 5},
		{342, 2, 0, 2},
	}
	sums := make([]int, len(mat))
	For(len(sums), func(i int) {
		row := mat[i]
		var sum int
		for _, n := range row {
			sum += n
		}
		sums[i] = sum
	})
	fmt.Println(sums)
	// Output: [124 123 346]
}

func ExampleFor_loopWithStep() {
	data := []int{1, 2, 3, 4, 5, 6}
	// Index 1, 3, 5
	fi := ranger.Range(1, len(data), 2)
	For(fi.Size, func(i int) {
		data[fi.Ith(i)] += 10
	})
	fmt.Println(data)
	// Output: [1 12 3 14 5 16]
}

// Sum adds all the numbers in the in channel and returns the sum.
func sum(in <-chan int) int {
	var total int
	for n := range in {
		total += n
	}
	return total
}

func ExampleFor_sum() {
	// An example with generator patterns. If the generator returns a
	// channel, the generator is responsible for closing the channel. Pros
	// of this pattern is that the generators can be used naturally like
	// normal function calls like sum(produce()).
	produce := func() <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for i := 0; i < 100; i++ {
				out <- i
			}
		}()
		return out
	}
	partial := func(n int, in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			For(n, func(i int) {
				out <- sum(in)
			})
		}()
		return out
	}
	// 1-to-1 producer-consumer.
	fmt.Println(sum(produce()))
	// Parallel workers.
	fmt.Println(sum(partial(10, produce())))
	// Output:
	// 4950
	// 4950
}

// A different pattern to use Do function to run different pipeline
// stages in parallel. Since it removes go func() from the code, it may
// look simpler. But it's more difficult to reuse each stage worker
// implementation.
func ExampleDo_sum() {
	nums, partial := make(chan int), make(chan int)
	Do(
		func() {
			defer close(nums)
			for i := 0; i < 100; i++ {
				nums <- i
			}
		},
		func() {
			defer close(partial)
			For(10, func(i int) {
				partial <- sum(nums)
			})
		},
		func() {
			fmt.Println(sum(partial))
		},
	)
	// Output:
	// 4950
}

// This example shows how to terminate all pipeline workers and returns the
// error, when an error occurred from one of the workers.
func ExampleDo_sumReturnErr() {
	intentionalError := false
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Defer not to leak.
	sumLines := func() (n int, err error) {
		lines, nums, partial := make(chan string), make(chan int), make(chan int)
		var once sync.Once
		returnErr := func(newErr error) {
			if newErr == nil {
				return
			}
			// Returns the first error only.
			once.Do(func() {
				cancel()
				err = newErr
			})
		}
		Do(
			func() {
				defer close(lines)
				for i := 0; i < 100; i++ {
					next := fmt.Sprint(i)
					if intentionalError && i == 55 {
						next = ""
					}
					select {
					case lines <- next:
					case <-ctx.Done():
						return
					}
				}
			},
			func() {
				defer close(nums)
				for line := range lines {
					n, err := strconv.Atoi(line)
					if err != nil {
						returnErr(errors.New("intentional error"))
						return
					}
					select {
					case nums <- n:
					case <-ctx.Done():
						return
					}
				}
			},
			func() {
				defer close(partial)
				For(10, func(i int) {
					select {
					case partial <- sum(nums):
					case <-ctx.Done():
						return
					}
				})
			},
			func() {
				n = sum(partial)
			},
		)
		if err != nil {
			return 0, err
		}
		return n, nil
	}
	fmt.Println(sumLines())
	intentionalError = true
	fmt.Println(sumLines())
	// Output:
	// 4950 <nil>
	// 0 intentional error
}
