package ranger

import (
	"fmt"
	"sort"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

var ins = []struct {
	begin int
	end   int
	step  int
}{
	{0, 10, 0},
	{0, 10, 1},
	{0, 10, -1},
	{10, 0, -1},
	{0, 10, 2},
	{0, 10, 3},
}

func ExampleSizeOfRange() {
	for _, in := range ins {
		fmt.Printf("%v => %d\n", in, SizeOfRange(in.begin, in.end, in.step))
	}
	// Output:
	// {0 10 0} => 10
	// {0 10 1} => 10
	// {0 10 -1} => 0
	// {10 0 -1} => 10
	// {0 10 2} => 5
	// {0 10 3} => 4
}

func TestSizeOfRange_empty(t *testing.T) {
	properties := gopter.NewProperties(nil)
	properties.Property("empty returns true iff the for-loop count equals zero", prop.ForAll(
		func(begin, end, step int) bool {
			var count int
			if step > 0 {
				for i := begin; i < end; i += step {
					count++
				}
			} else if step < 0 {
				for i := begin; i > end; i += step {
					count++
				}
			} else {
				for i := begin; i < end; i++ {
					count++
				}
			}
			size := SizeOfRange(begin, end, step)
			return count == 0 && size == 0 || count > 0 && size > 0
		},
		gen.IntRange(-10000, 10000),
		gen.IntRange(-10000, 10000),
		gen.IntRange(-30000, 30000),
	))
	properties.TestingRun(t)
}

func TestSizeOfRange(t *testing.T) {
	properties := gopter.NewProperties(nil)
	properties.Property("size should match with what the for-loop counts", prop.ForAll(
		func(begin, end, step int) bool {
			var count int
			if step > 0 {
				for i := begin; i < end; i += step {
					count++
				}
			} else if step < 0 {
				for i := begin; i > end; i += step {
					count++
				}
			} else {
				for i := begin; i < end; i++ {
					count++
				}
			}
			return count == SizeOfRange(begin, end, step)
		},
		gen.IntRange(-10000, 10000),
		gen.IntRange(-10000, 10000),
		gen.IntRange(-30000, 30000),
	))
	properties.TestingRun(t)
}

func ExampleRange() {
	for _, in := range ins {
		fi := Range(in.begin, in.end, in.step)
		var ints []int
		for i := 0; i < fi.Size; i++ {
			ints = append(ints, fi.Ith(i))
		}
		fmt.Printf("%v => %v\n", in, ints)
	}
	// Output:
	// {0 10 0} => [0 1 2 3 4 5 6 7 8 9]
	// {0 10 1} => [0 1 2 3 4 5 6 7 8 9]
	// {0 10 -1} => []
	// {10 0 -1} => [10 9 8 7 6 5 4 3 2 1]
	// {0 10 2} => [0 2 4 6 8]
	// {0 10 3} => [0 3 6 9]
}

func ExampleRange_iterator() {
	for _, in := range ins {
		fi := Range(in.begin, in.end, in.step)
		itr := FiniteIterator{Size: fi.Size, Itr: &IthIterator{IthFunc: fi.Ith}}
		var ints []int
		for !itr.Empty() {
			ints = append(ints, itr.Next())
		}
		fmt.Printf("%v => %v\n", in, ints)
	}
	// Output:
	// {0 10 0} => [0 1 2 3 4 5 6 7 8 9]
	// {0 10 1} => [0 1 2 3 4 5 6 7 8 9]
	// {0 10 -1} => []
	// {10 0 -1} => [10 9 8 7 6 5 4 3 2 1]
	// {0 10 2} => [0 2 4 6 8]
	// {0 10 3} => [0 3 6 9]
}

func TestRange_Ith(t *testing.T) {
	properties := gopter.NewProperties(nil)
	properties.Property("Ith should produce the same number that for-loop produces", prop.ForAll(
		func(begin, end, step int) bool {
			var count int
			fi := Range(begin, end, step)
			if step > 0 {
				for i := begin; i < end; i += step {
					if fi.Ith(count) != i {
						return false
					}
					count++
				}
			} else if step < 0 {
				for i := begin; i > end; i += step {
					if fi.Ith(count) != i {
						return false
					}
					count++
				}
			} else {
				for i := begin; i < end; i++ {
					if fi.Ith(count) != i {
						return false
					}
					count++
				}
			}
			return true
		},
		gen.IntRange(-10000, 10000),
		gen.IntRange(-10000, 10000),
		gen.IntRange(-30000, 30000),
	))
	properties.TestingRun(t)
}

func ExampleFromIndices() {
	fi := FromIndices([]int{1, 3, 5, 7})
	for i := 0; i < fi.Size; i++ {
		fmt.Println(fi.Ith(i))
	}
	// Output:
	// 1
	// 3
	// 5
	// 7
}

func ExampleFiniteIth_Partition() {
	ex := Range(0, 22, 1)
	chunks := 4
	for i := 0; i < chunks; i++ {
		p := ex.Partition(i, chunks)
		var elems []int
		for j := 0; j < p.Size; j++ {
			elems = append(elems, p.Ith(j))
		}
		fmt.Println(elems)
	}
	// Output:
	// [0 1 2 3 4 5]
	// [6 7 8 9 10 11]
	// [12 13 14 15 16 17]
	// [18 19 20 21]
}

func ExampleFiniteIth_AsSortInterface() {
	s := NDShape(ND{3, 3})
	alphas := sort.StringSlice{
		"a", "b", "c",
		"d", "e", "f",
		"g", "h", "i",
	}
	ordering := []int{
		s.ToIth(ND{0, 0}),
		s.ToIth(ND{0, 1}),
		s.ToIth(ND{1, 1}),
		s.ToIth(ND{1, 0}),
		s.ToIth(ND{2, 0}),
		s.ToIth(ND{2, 1}),
		s.ToIth(ND{2, 2}),
		s.ToIth(ND{1, 2}),
		s.ToIth(ND{0, 2}),
	}
	sort.Sort(FromIndices(ordering).AsSortInterface(alphas))
	for i := 0; i < 3; i++ {
		fmt.Println(alphas[s.ToIth(ND{i, 0}):s.ToIth(ND{i + 1, 0})])
	}
	// Output:
	// [a b i]
	// [d c h]
	// [e f g]
}
