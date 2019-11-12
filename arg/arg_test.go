package arg

import (
	"fmt"
	"sort"
)

func ExampleMin() {
	exs := []sort.Interface{
		sort.IntSlice([]int{5, 3, 9, 2, 3, 8}),
		sort.Float64Slice([]float64{5.0, 3.0, 9.2, 2.3, 3.4, 8.5}),
		sort.StringSlice([]string{"rabbit", "apple", "zebra", "dog", "cat"}),
	}
	fmt.Println("== Min ==")
	for _, ex := range exs {
		Min := Min(ex.Len(), ex.Less)
		fmt.Printf("Min(%#v) => data[%d]\n", ex, Min)
	}
	fmt.Println("== Max ==")
	for _, ex := range exs {
		Max := Max(ex.Len(), ex.Less)
		fmt.Printf("Max(%#v) => data[%d]\n", ex, Max)
	}
	// Output:
	// == Min ==
	// Min(sort.IntSlice{5, 3, 9, 2, 3, 8}) => data[3]
	// Min(sort.Float64Slice{5, 3, 9.2, 2.3, 3.4, 8.5}) => data[3]
	// Min(sort.StringSlice{"rabbit", "apple", "zebra", "dog", "cat"}) => data[1]
	// == Max ==
	// Max(sort.IntSlice{5, 3, 9, 2, 3, 8}) => data[2]
	// Max(sort.Float64Slice{5, 3, 9.2, 2.3, 3.4, 8.5}) => data[2]
	// Max(sort.StringSlice{"rabbit", "apple", "zebra", "dog", "cat"}) => data[2]
}
