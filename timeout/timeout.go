// Package timeout provides a utility function to execute a function with a
// timeout. It helps in scenarios where an operation needs to be bound by a time
// limit, preventing indefinite blocking.
package timeout

import (
	"context"
	"time"
)

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
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// Buffer of 1 to prevent sender from blocking if receiver is not ready
	done := make(chan error, 1)

	go func() {
		// Close the channel when the goroutine exits
		defer close(done)
		done <- f()
	}()

	select {
	case <-ctxWithTimeout.Done():
		return ctxWithTimeout.Err()
	case err := <-done:
		return err
	}
}
