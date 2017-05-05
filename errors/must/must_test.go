package must

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"testing"
)

func recoverAnyPanic() {
	if r := recover(); r != nil {
		fmt.Println("panic:", r)
	}
}

func repeatHello(num string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
		}
	}()
	defer ReturnErr(&err)
	for i := Int(strconv.Atoi(num)); i > 0; i-- {
		fmt.Println("hello")
	}
	if num == "3" {
		panic(errors.New("panic for 3"))
	}
	return nil
}

func TestReturnError(t *testing.T) {
	if repeatHello("a") == nil {
		t.Error("should return non-nil error")
	}
}

func TestReturnError_NoError(t *testing.T) {
	if repeatHello("0") != nil {
		t.Error("should return nil error")
	}
}

func ExampleReturnErr() {
	// TODO: Fix the fragile test dependency to the error message.
	err := func() (err error) {
		// All errors captured by must package is returned. Without this defer line
		// captured errors will panic.
		defer ReturnErr(&err)

		// Error handling is simplified here.
		i := Int(strconv.Atoi("a"))
		fmt.Println(i)

		// Or inlined.
		fmt.Println(Int(strconv.Atoi("b")))

		// Compare with the following.
		i, err = strconv.Atoi("c")
		if err != nil {
			return err
		}

		return nil
	}()
	fmt.Println(err)
	// Output:
	// strconv.Atoi: parsing "a": invalid syntax
}

func ExampleReturnErr_multipleRecover() {
	func() (err error) {
		// Deferred calls are executed in last-in-first-out order. If a deferred
		// function recovers from any panic, defer of the function should come
		// before defer ReturnErr.
		defer recoverAnyPanic()
		defer ReturnErr(&err)

		_ = Int(strconv.Atoi("1"))
		panic(errors.New("created panic"))
		return nil
	}()
	// Output:
	// panic: created panic
}

func ExampleAny() {
	// TODO: Fix the fragile test dependency to the error message.
	err := func() (err error) {
		defer ReturnErr(&err)

		// Any can be used with any types, but it's less efficient. For any types
		// that must package does not support, consider writing your own 2-line
		// defer function, or use Any.
		var i int = Any(strconv.Atoi("a")).(int)
		fmt.Println(i)

		return nil
	}()
	fmt.Println(err)
	// Output:
	// strconv.Atoi: parsing "a": invalid syntax
}

func ExampleLogErr() {
	// TODO: Fix the fragile test dependency to the error message.
	err := func() (err error) {
		defer ReturnErr(&err)
		// Log error to stderr with file name and line number.
		defer LogErr(log.Println)
		_ = Int(strconv.Atoi("a"))
		return nil
	}()
	fmt.Println(err.Error())
	// Output:
	// strconv.Atoi: parsing "a": invalid syntax
}
