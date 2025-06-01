package timeout

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestDoWithTimeout_Success(t *testing.T) {
	ctx := context.Background()
	timeout := 100 * time.Millisecond
	f := func() error {
		time.Sleep(10 * time.Millisecond)
		return nil
	}

	err := DoWithTimeout(ctx, timeout, f)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestDoWithTimeout_Timeout(t *testing.T) {
	ctx := context.Background()
	timeout := 10 * time.Millisecond
	f := func() error {
		time.Sleep(100 * time.Millisecond)
		return nil
	}

	err := DoWithTimeout(ctx, timeout, f)
	if err != context.DeadlineExceeded {
		t.Errorf("Expected context.DeadlineExceeded, got %v", err)
	}
}

func TestDoWithTimeout_ContextCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	timeout := 100 * time.Millisecond
	f := func() error {
		time.Sleep(10 * time.Millisecond)
		return nil
	}

	// Cancel context before calling DoWithTimeout
	cancel()
	err := DoWithTimeout(ctx, timeout, f)
	if err != context.Canceled {
		t.Errorf("Expected context.Canceled, got %v", err)
	}
}

func ExampleDoWithTimeout() {
	// Scenario 1: Function completes successfully within timeout
	ctx1 := context.Background()
	timeout1 := 100 * time.Millisecond
	f1 := func() error {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("Function 1 completed")
		return nil
	}
	err1 := DoWithTimeout(ctx1, timeout1, f1)
	if err1 != nil {
		fmt.Printf("Function 1 error: %v\n", err1)
	}

	// Scenario 2: Function times out
	ctx2 := context.Background()
	timeout2 := 10 * time.Millisecond
	f2 := func() error {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Function 2 completed (this should not print if timeout works)")
		return nil
	}
	err2 := DoWithTimeout(ctx2, timeout2, f2)
	if err2 != nil {
		fmt.Printf("Function 2 error: %v\n", err2)
	}

	// Output:
	// Function 1 completed
	// Function 2 error: context deadline exceeded
}

func TestDoWithTimeout_FunctionError(t *testing.T) {
	ctx := context.Background()
	timeout := 100 * time.Millisecond
	expectedErr := errors.New("function error")
	f := func() error {
		return expectedErr
	}

	err := DoWithTimeout(ctx, timeout, f)
	if err != expectedErr {
		t.Errorf("Expected error %v, got %v", expectedErr, err)
	}
}

func TestDoWithTimeout_ContextCanceledDuringExecution(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	timeout := 100 * time.Millisecond
	f := func() error {
		// Sleep long enough for cancellation to occur
		time.Sleep(50 * time.Millisecond)
		return nil
	}

	go func() {
		// Wait a bit then cancel
		time.Sleep(10 * time.Millisecond)
		cancel()
	}()

	err := DoWithTimeout(ctx, timeout, f)
	if err != context.Canceled {
		t.Errorf("Expected context.Canceled, got %v", err)
	}
}
