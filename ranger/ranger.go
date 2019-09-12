// Package ranger provides some functions to deal with integer indices.
package ranger

import (
	"fmt"
	"sort"
)

// Sgn returns the sign of the given integer x. It returns -1, 0, or 1.
func Sgn(x int) int {
	switch {
	case x > 0:
		return 1
	case x < 0:
		return -1
	default:
		return 0
	}
}

// Abs returns the absolute value of the integer x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// ArithProg implements infinite arithmetic progression.
type ArithProg struct {
	Begin int
	Step  int
}

// Ith returns the ith number.
func (ap ArithProg) Ith(i int) int {
	return ap.Begin + i*ap.Step
}

// Empty returns true if the sequence is empty. For infinite sequence, it's
// always false.
func (ap ArithProg) Empty() bool {
	return false
}

// Next returns the next integer and updates the progression.
func (ap *ArithProg) Next() int {
	n := ap.Begin
	ap.Begin += ap.Step
	return n
}

// FiniteIthFunc returns a function that checks if 0 <= i < size and returns ith
// number in sequencee.
func FiniteIthFunc(size int, ithFunc func(i int) int) func(i int) int {
	return func(i int) int {
		if i < 0 {
			panic(fmt.Sprintf("invlid index %d (i must be non-negative)", i))
		}
		if i >= size {
			panic(fmt.Sprintf("index %d out of bound", i))
		}
		return ithFunc(i)
	}
}

// Iterator is an interface to iterate integers.
type Iterator interface {
	Empty() bool
	Next() int
}

// IthIterator implements an infinite iterator with the given IthFunc.
type IthIterator struct {
	IthFunc func(i int) int
	nextIdx int
}

// Empty returns true if the sequence is empty. It alsways returns false as it's
// infinite.
func (i IthIterator) Empty() bool {
	return false
}

// Next returns the next integer and advance the iterator.
func (i *IthIterator) Next() int {
	n := i.IthFunc(i.nextIdx)
	i.nextIdx++
	return n
}

// FiniteIterator implements a finite iterator.
type FiniteIterator struct {
	Size int
	Itr  Iterator
}

// Empty returns true if the sequence is empty.
func (f FiniteIterator) Empty() bool {
	return f.Itr.Empty() || f.Size == 0
}

// Next returns the first number and then modify itself to the next numbers. The
// caller should check Empty() first before calling this function.
func (f *FiniteIterator) Next() int {
	f.Size--
	return f.Itr.Next()
}

// SizeOfRange returns the size of for-loop like begin, end, step range. Begin
// is inclusive and end is exclusive. If step is 0, it is treated as 1.
func SizeOfRange(begin, end, step int) int {
	if step == 0 {
		step = 1
	}
	if Sgn(end-begin) != Sgn(step) {
		return 0
	}
	return (Abs(end-begin) + Abs(step) - 1) / Abs(step)
}

// FiniteIth has the size and ith function together.
type FiniteIth struct {
	Size int
	Ith  func(i int) int
}

// Range returns a new finite range with the given begin, end, step.
// Begin is inclusive and end is exclusive. If step is 0, it is treated as 1.
func Range(begin, end, step int) FiniteIth {
	if step == 0 {
		step = 1
	}
	ap := ArithProg{Begin: begin, Step: step}
	size := SizeOfRange(begin, end, step)
	return FiniteIth{Size: size, Ith: FiniteIthFunc(size, ap.Ith)}
}

// FromIndices returns a new FiniteIth where ith is mapped to indices[i].
func FromIndices(indices []int) FiniteIth {
	return FiniteIth{Size: len(indices), Ith: func(i int) int {
		return indices[i]
	}}
}

// Partition returns ith part out of numParts parts of the range.
func (f FiniteIth) Partition(i, numParts int) FiniteIth {
	if numParts <= 0 || i >= numParts {
		panic("wrong arguments")
	}
	if numParts == 1 {
		return f
	}
	partSize := (f.Size + numParts - 1) / numParts
	begin := partSize * i
	if i == numParts-1 {
		partSize = f.Size - begin
	}
	return FiniteIth{Size: partSize, Ith: func(i int) int {
		return f.Ith(begin + i)
	}}
}

// AsSortInterface returns a sort.Interface based on the given data in this
// range.
func (f FiniteIth) AsSortInterface(data sort.Interface) SortFiniteIth {
	return SortFiniteIth{FiniteIth: f, data: data}
}

// SortFiniteIth implements sort interface.
type SortFiniteIth struct {
	FiniteIth
	data sort.Interface
}

// Len returns the size of FiniteIth. It implements sort.Interface.
func (sfi SortFiniteIth) Len() int {
	return sfi.Size
}

// Less returns true if the translated ith element is less than the translated
// jth element. It implements sort.Interface.
func (sfi SortFiniteIth) Less(i, j int) bool {
	return sfi.data.Less(sfi.Ith(i), sfi.Ith(j))
}

// Swap swaps the translated ith element and the translated jth element. It
// implements sort.Interface.
func (sfi SortFiniteIth) Swap(i, j int) {
	sfi.data.Swap(sfi.Ith(i), sfi.Ith(j))
}
