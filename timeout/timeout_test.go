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
	if !errors.Is(err, context.DeadlineExceeded) {
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
	if !errors.Is(err, context.Canceled) {
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
	done := make(chan struct{})
	f2 := func() error {
		defer close(done)
		time.Sleep(100 * time.Millisecond)
		return nil
	}
	err2 := DoWithTimeout(ctx2, timeout2, f2)
	if err2 != nil {
		fmt.Printf("Function 2 error: %v\n", err2)
	}
	// Wait for the goroutine to finish to avoid a data race on stdout.
	<-done

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
	if !errors.Is(err, expectedErr) {
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
	if !errors.Is(err, context.Canceled) {
		t.Errorf("Expected context.Canceled, got %v", err)
	}
}

func TestDo_Success(t *testing.T) {
	ctx := context.Background()
	timeout := 100 * time.Millisecond
	f := func() (string, error) {
		time.Sleep(10 * time.Millisecond)
		return "ok", nil
	}

	got, err := Do(ctx, timeout, f)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if got != "ok" {
		t.Errorf("Expected %q, got %q", "ok", got)
	}
}

func TestDo_Timeout(t *testing.T) {
	ctx := context.Background()
	timeout := 10 * time.Millisecond
	f := func() (string, error) {
		time.Sleep(100 * time.Millisecond)
		return "too late", nil
	}

	got, err := Do(ctx, timeout, f)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("Expected context.DeadlineExceeded, got %v", err)
	}
	if got != "" {
		t.Errorf("Expected zero value, got %q", got)
	}
}

func TestDo_ContextCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	timeout := 100 * time.Millisecond
	f := func() (string, error) {
		time.Sleep(10 * time.Millisecond)
		return "ok", nil
	}

	cancel()
	got, err := Do(ctx, timeout, f)
	if !errors.Is(err, context.Canceled) {
		t.Errorf("Expected context.Canceled, got %v", err)
	}
	if got != "" {
		t.Errorf("Expected zero value, got %q", got)
	}
}

func TestDo_FunctionError(t *testing.T) {
	ctx := context.Background()
	timeout := 100 * time.Millisecond
	expectedErr := errors.New("function error")
	f := func() (string, error) {
		return "partial", expectedErr
	}

	got, err := Do(ctx, timeout, f)
	if !errors.Is(err, expectedErr) {
		t.Errorf("Expected error %v, got %v", expectedErr, err)
	}
	if got != "partial" {
		t.Errorf("Expected %q, got %q", "partial", got)
	}
}

func TestDo_ContextCanceledDuringExecution(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	timeout := 100 * time.Millisecond
	f := func() (int, error) {
		time.Sleep(50 * time.Millisecond)
		return 42, nil
	}

	go func() {
		time.Sleep(10 * time.Millisecond)
		cancel()
	}()

	got, err := Do(ctx, timeout, f)
	if !errors.Is(err, context.Canceled) {
		t.Errorf("Expected context.Canceled, got %v", err)
	}
	if got != 0 {
		t.Errorf("Expected zero value, got %d", got)
	}
}

func ExampleDo() {
	ctx := context.Background()
	timeout := 100 * time.Millisecond
	f := func() (string, error) {
		return "hello", nil
	}
	got, err := Do(ctx, timeout, f)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Println(got)
	// Output:
	// hello
}
