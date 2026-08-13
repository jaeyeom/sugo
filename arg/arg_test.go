package arg

import (
	"fmt"
	"sort"
	"testing"
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

func TestMinOf(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want int
	}{
		{name: "typical", s: []int{5, 3, 9, 2, 3, 8}, want: 3},
		{name: "first is min", s: []int{1, 2, 3}, want: 0},
		{name: "last is min", s: []int{3, 2, 1}, want: 2},
		{name: "single element", s: []int{42}, want: 0},
		{name: "first of ties", s: []int{2, 1, 1, 3}, want: 1},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := MinOf(tc.s); got != tc.want {
				t.Errorf("MinOf(%v) = %d, want %d", tc.s, got, tc.want)
			}
		})
	}
}

func TestMaxOf(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want int
	}{
		{name: "typical", s: []int{5, 3, 9, 2, 3, 8}, want: 2},
		{name: "first is max", s: []int{9, 2, 3}, want: 0},
		{name: "last is max", s: []int{3, 2, 9}, want: 2},
		{name: "single element", s: []int{42}, want: 0},
		{name: "first of ties", s: []int{2, 9, 9, 3}, want: 1},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := MaxOf(tc.s); got != tc.want {
				t.Errorf("MaxOf(%v) = %d, want %d", tc.s, got, tc.want)
			}
		})
	}
}

func TestMinOfStrings(t *testing.T) {
	s := []string{"rabbit", "apple", "zebra", "dog", "cat"}
	if got, want := MinOf(s), 1; got != want {
		t.Errorf("MinOf(%v) = %d, want %d", s, got, want)
	}
}

func TestMaxOfStrings(t *testing.T) {
	s := []string{"rabbit", "apple", "zebra", "dog", "cat"}
	if got, want := MaxOf(s), 2; got != want {
		t.Errorf("MaxOf(%v) = %d, want %d", s, got, want)
	}
}

func TestMinFunc(t *testing.T) {
	type person struct {
		name string
		age  int
	}
	s := []person{
		{name: "ann", age: 30},
		{name: "bob", age: 20},
		{name: "cara", age: 25},
	}
	lessAge := func(a, b person) bool { return a.age < b.age }
	if got, want := MinFunc(s, lessAge), 1; got != want {
		t.Errorf("MinFunc(by age) = %d, want %d", got, want)
	}
}

func TestMaxFunc(t *testing.T) {
	type person struct {
		name string
		age  int
	}
	s := []person{
		{name: "ann", age: 30},
		{name: "bob", age: 20},
		{name: "cara", age: 25},
	}
	lessAge := func(a, b person) bool { return a.age < b.age }
	if got, want := MaxFunc(s, lessAge), 0; got != want {
		t.Errorf("MaxFunc(by age) = %d, want %d", got, want)
	}
}

func TestMinOfEmptyPanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("MinOf([]int{}) did not panic")
		}
	}()
	MinOf([]int{})
}

func TestMaxOfEmptyPanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("MaxOf([]int{}) did not panic")
		}
	}()
	MaxOf([]int{})
}

func TestMinFuncEmptyPanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("MinFunc([]int{}, less) did not panic")
		}
	}()
	MinFunc([]int{}, func(a, b int) bool { return a < b })
}

func TestMaxFuncEmptyPanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("MaxFunc([]int{}, less) did not panic")
		}
	}()
	MaxFunc([]int{}, func(a, b int) bool { return a < b })
}

func ExampleMinOf() {
	fmt.Println(MinOf([]int{5, 3, 9, 2, 3, 8}))
	fmt.Println(MaxOf([]string{"rabbit", "apple", "zebra"}))
	// Output:
	// 3
	// 2
}

func ExampleMinFunc() {
	type person struct {
		name string
		age  int
	}
	people := []person{
		{name: "ann", age: 30},
		{name: "bob", age: 20},
		{name: "cara", age: 25},
	}
	lessAge := func(a, b person) bool { return a.age < b.age }
	fmt.Println(MinFunc(people, lessAge))
	fmt.Println(MaxFunc(people, lessAge))
	// Output:
	// 1
	// 0
}
