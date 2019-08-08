package par

import (
	"fmt"
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
	// Output:
	// [124 123 346]
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
	sum := func(in <-chan int) int {
		var total int
		for n := range in {
			total += n
		}
		return total
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

func ExampleDo_sum() {
	// A different pattern to use Do function to run different pipeline
	// stages in parallel. Since it removes go func() from the code, it may
	// look simpler. But it's more difficult to reuse each stage worker
	// implementation.
	sum := func(in <-chan int) int {
		var total int
		for n := range in {
			total += n
		}
		return total
	}
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
