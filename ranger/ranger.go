// Package ranger provides some functions to deal with integer indices.
package ranger

import "fmt"

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

// FiniteIth returns a function that checks if 0 <= i < size and returns ith
// number in sequencee.
func FiniteIth(size int, ithFunc func(i int) int) func(i int) int {
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

// Range returns a new finite range with the given begin, end, step.
// Begin is inclusive and end is exclusive. If step is 0, it is treated as 1.
func Range(begin, end, step int) (size int, ithFunc func(i int) int) {
	if step == 0 {
		step = 1
	}
	ap := ArithProg{Begin: begin, Step: step}
	size = SizeOfRange(begin, end, step)
	return size, FiniteIth(size, ap.Ith)
}
