// Package timeout provides a utility function to execute a function with a
// timeout. It helps in scenarios where an operation needs to be bound by a time
// limit, preventing indefinite blocking.
package timeout

import (
	"context"
	"time"
)

// Do executes f within the specified timeout and returns its result.
//
// If f completes within the timeout, Do returns the value and error from f. If
// the timeout duration is reached before f completes, Do returns the zero
// value of T and [context.DeadlineExceeded]. If the parent context is canceled
// before f completes, Do returns the zero value of T and [context.Canceled].
func Do[T any](ctx context.Context, timeout time.Duration, f func() (T, error)) (T, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	type result struct {
		val T
		err error
	}
	// Buffer of 1 to prevent sender from blocking if receiver is not ready.
	done := make(chan result, 1)

	go func() {
		// Close the channel when the goroutine exits.
		defer close(done)
		val, err := f()
		done <- result{val, err}
	}()

	select {
	case <-ctxWithTimeout.Done():
		var zero T
		return zero, ctxWithTimeout.Err()
	case r := <-done:
		return r.val, r.err
	}
}

// DoWithTimeout executes the given function f within the specified timeout
// duration. It takes a parent context, a timeout duration, and the function to
// execute. The function f is of type func() error.
//
// DoWithTimeout returns nil if f completes successfully within the timeout. If
// f returns an error, DoWithTimeout returns that error. If the timeout duration
// is reached before f completes, DoWithTimeout returns
// [context.DeadlineExceeded]. If the parent context is canceled before f
// completes, DoWithTimeout returns [context.Canceled].
func DoWithTimeout(ctx context.Context, timeout time.Duration, f func() error) error {
	_, err := Do(ctx, timeout, func() (struct{}, error) {
		return struct{}{}, f()
	})
	return err
}
