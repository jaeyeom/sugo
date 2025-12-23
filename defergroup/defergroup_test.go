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

func ExampleGroup_CancelAll() {
	f := func() {
		g := New()
		defer g.Done()
		g.Defer(func() {
			fmt.Println("deferred 1")
		})
		g.Defer(func() {
			fmt.Println("deferred 2")
		})
		g.CancelAll()
	}
	f()
	// Output:
}

func ExampleGroup_noError() {
	// When there's no error, the deferred functions are ignored.
	f := func() (err error) {
		g := New(OnlyOnError(&err))
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
		g := New(OnlyOnError(&err))
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
		g := New(OnlyOnError(&err))
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
		g := New(OnlyOnError(&err))
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

// server holds multiple resources that need cleanup.
type server struct {
	r1 *resource
	r2 *resource
	dg *Group
}

// newServerWith creates a server using the provided resource constructors.
// If any acquisition fails, previously acquired resources are cleaned up automatically.
func newServerWith(newR1, newR2 func(string) (*resource, error)) (*server, error) {
	dg := New()
	defer dg.Done()

	r1, err := newR1("connection")
	if err != nil {
		return nil, err
	}
	dg.Defer(r1.Close)

	r2, err := newR2("file handle")
	if err != nil {
		return nil, err
	}
	dg.Defer(r2.Close)

	// All resources acquired successfully - transfer cleanup responsibility
	// to the server struct
	return &server{
		r1: r1,
		r2: r2,
		dg: dg.Transfer(),
	}, nil
}

// newServer creates a server, acquiring multiple resources. If any acquisition
// fails, previously acquired resources are cleaned up automatically.
func newServer() (*server, error) {
	return newServerWith(newResource, newResource)
}

func (s *server) Close() {
	s.dg.Done()
}

func ExampleGroup_Transfer() {
	srv, err := newServer()
	if err != nil {
		fmt.Println("failed to create server")
		return
	}
	defer srv.Close()
	fmt.Println("server created")
	// Output:
	// server created
	// closing file handle
	// closing connection
}

func ExampleGroup_Transfer_partialFailure() {
	srv, err := newServerWith(newResource, newFailedResource)
	if err != nil {
		fmt.Println("failed to create server")
		return
	}
	defer srv.Close()
	fmt.Println("server created")
	// Output:
	// closing connection
	// failed to create server
}
