// Package defergroup provides a mechanism to manage groups of deferred functions,
// executing them in last-in-first-out order once triggered. This package is especially
// useful for managing cleanup tasks that need to be executed if earlier operations in
// a function fail or need rollback.
//
// The key component of this package is the Group struct, which maintains a list of
// deferred functions to be executed. Functions can be added using the Defer method
// and executed using the Done method. The execution of these functions can be conditional
// on the presence of an error through the use of options, specifically the WithError
// option, which skips execution if no error is set. This allows for flexible resource
// management, ensuring that cleanup code is only executed when necessary.
package defergroup

// Group is a deferred function group.
type Group struct {
	fns []func()
	err *error
}

// Option is a group option.
type Option func(*Group)

// WithError sets the error pointer for the group. If the error is nil, the
// group will be skipped. This is useful when the group is used to clean up
// partially initialized resources. Please ensure that the pointer to the named
// return error variable is passed.
func WithError(err *error) Option {
	return func(g *Group) {
		g.err = err
	}
}

// New creates a new group. If an error pointer is passed, the group will be
// skipped if the error is nil.
func New(opts ...Option) *Group {
	g := &Group{}
	for _, opt := range opts {
		opt(g)
	}
	return g
}

// Done runs the deferred functions in the group in last-in-first-out order. It
// will skip the deferred functions if the error is nil.
func (g *Group) Done() {
	if g.err != nil && *g.err == nil {
		return
	}
	for i := len(g.fns) - 1; i >= 0; i-- {
		g.fns[i]()
	}
}

// Defer adds a deferred function to the group.
func (g *Group) Defer(f func()) {
	g.fns = append(g.fns, f)
}

// Clear clears the deferred functions in the group.
func (g *Group) Clear() {
	g.fns = nil
}
