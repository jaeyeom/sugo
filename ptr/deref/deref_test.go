package deref

import (
	"fmt"
	"testing"
)

func ExampleOf() {
	n := 42
	fmt.Println(Of(&n))
	fmt.Println(Of((*int)(nil)))
	// Output:
	// 42
	// 0
}

func ExampleOr() {
	n := 42
	fmt.Println(Or(&n, 7))
	fmt.Println(Or((*int)(nil), 7))
	// Output:
	// 42
	// 7
}

func TestOf(t *testing.T) {
	type point struct {
		X, Y int
	}
	p := point{X: 1, Y: 2}

	if got, want := Of(&p), p; got != want {
		t.Errorf("Of(&p) = %#v, want %#v", got, want)
	}
	if got, want := Of((*point)(nil)), (point{}); got != want {
		t.Errorf("Of(nil) = %#v, want %#v", got, want)
	}
}

func TestOr(t *testing.T) {
	type point struct {
		X, Y int
	}
	p := point{X: 1, Y: 2}
	def := point{X: 9, Y: 9}

	if got := Or(&p, def); got != p {
		t.Errorf("Or(&p, def) = %#v, want %#v", got, p)
	}
	if got := Or((*point)(nil), def); got != def {
		t.Errorf("Or(nil, def) = %#v, want %#v", got, def)
	}
}
