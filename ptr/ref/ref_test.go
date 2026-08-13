package ref

import (
	"fmt"
	"testing"
)

func ExampleOf() {
	fmt.Println(*Of(42))
	fmt.Println(*Of("hello"))
	// Output:
	// 42
	// hello
}

func TestOf(t *testing.T) {
	type point struct {
		X, Y int
	}
	p := Of(point{X: 1, Y: 2})
	if p == nil || *p != (point{X: 1, Y: 2}) {
		t.Errorf("Of(point) = %#v, want &{1 2}", p)
	}
}

func TestOf_copiesValue(t *testing.T) {
	v := 1
	p := Of(v)
	v = 2
	if *p != 1 || v != 2 {
		t.Errorf("Of() = %d after changing original to %d, want 1", *p, v)
	}
}
