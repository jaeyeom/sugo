package defergroup

import (
	"errors"
	"fmt"
)

func ExampleGroup_noop() {
	f := func() {
		g := New()
		defer g.Done()
	}
	f()
	// Output:
}

func ExampleGroup_single() {
	f := func() {
		g := New()
		defer g.Done()
		g.Defer(func() {
			fmt.Println("deferred")
		})
	}
	f()
	// Output:
	// deferred
}

func ExampleGroup_multiple() {
	f := func() {
		g := New()
		defer g.Done()
		g.Defer(func() {
			fmt.Println("deferred 1")
		})
		g.Defer(func() {
			fmt.Println("deferred 2")
		})
	}
	f()
	// Output:
	// deferred 2
	// deferred 1
}

func ExampleGroup_clear() {
	f := func() {
		g := New()
		defer g.Done()
		g.Defer(func() {
			fmt.Println("deferred 1")
		})
		g.Defer(func() {
			fmt.Println("deferred 2")
		})
		g.Clear()
	}
	f()
	// Output:
}

func ExampleGroup_noError() {
	// When there's no error, the deferred functions are ignored.
	f := func() (err error) {
		g := New(WithError(&err))
		defer g.Done()
		g.Defer(func() {
			fmt.Println("deferred 1")
		})
		g.Defer(func() {
			fmt.Println("deferred 2")
		})
		return nil
	}
	if err := f(); err != nil {
		fmt.Printf("Function f failed: %v\n", err)
	}
	// Output:
}

func ExampleGroup_error() {
	// When there's an error, the deferred functions are executed.
	f := func() (err error) {
		g := New(WithError(&err))
		defer g.Done()
		g.Defer(func() {
			fmt.Println("deferred 1")
		})
		g.Defer(func() {
			fmt.Println("deferred 2")
		})
		return errors.New("some error")
	}
	if err := f(); err != nil {
		fmt.Printf("Function f failed: %v\n", err)
	}
	// Output:
	// deferred 2
	// deferred 1
	// Function f failed: some error
}

type resource struct {
	name string
}

func newResource(name string) (*resource, error) {
	return &resource{name: name}, nil
}

func newFailedResource(name string) (*resource, error) {
	return nil, errors.New("unable to create resource")
}

func (r *resource) Close() {
	fmt.Printf("closing %s\n", r.name)
}

func ExampleGroup_noCloseForNoError() {
	f := func() (err error) {
		g := New(WithError(&err))
		defer g.Done()
		r1, err := newResource("resource 1")
		if err != nil {
			return err
		}
		g.Defer(r1.Close)
		r2, err := newResource("resource 2")
		if err != nil {
			return err
		}
		g.Defer(r2.Close)
		return nil
	}
	if err := f(); err != nil {
		fmt.Printf("Function f failed: %v\n", err)
	}
	// Output:
}

func ExampleGroup_partialClose() {
	f := func() (err error) {
		g := New(WithError(&err))
		defer g.Done()
		r1, err := newResource("resource 1")
		if err != nil {
			return err
		}
		g.Defer(r1.Close)
		r2, err := newFailedResource("resource 2")
		if err != nil {
			return err
		}
		g.Defer(r2.Close)
		return nil
	}
	if err := f(); err != nil {
		fmt.Printf("Function f failed: %v\n", err)
	}
	// Output:
	// closing resource 1
	// Function f failed: unable to create resource
}
