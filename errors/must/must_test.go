package must

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
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

func TestReturnError_noError(t *testing.T) {
	if repeatHello("0") != nil {
		t.Error("should return nil error")
	}
}

func ExampleReturnErr() {
	// TODO: Fix the fragile test dependency to the error message.
	err := func() (err error) {
		// All errors captured by must package is returned. Without this
		// defer line captured errors will panic.
		defer ReturnErr(&err)

		// Error handling is simplified here.
		// From other packages it looks like must.Int(...).
		i := Int(strconv.Atoi("a"))
		fmt.Println(i)

		// Or inlined.
		fmt.Println(Int(strconv.Atoi("b")))

		// Compare with the following.
		i, err = strconv.Atoi("c")
		if err != nil {
			return err
		}
		fmt.Println(i)

		return nil
	}()
	fmt.Println(err)
	// Output:
	// strconv.Atoi: parsing "a": invalid syntax
}

func ExampleReturnErr_multipleRecover() {
	_ = func() (err error) {
		// Deferred calls are executed in last-in-first-out order. If a
		// deferred function recovers from any panic, defer of the
		// function should come before defer ReturnErr.
		defer recoverAnyPanic()
		defer ReturnErr(&err)

		_ = Int(strconv.Atoi("1"))
		panic(errors.New("created panic"))
	}()
	// Output:
	// panic: created panic
}

func ExampleAny() {
	// TODO: Fix the fragile test dependency to the error message.
	err := func() (err error) {
		defer ReturnErr(&err)

		// Any can be used with any types, but it's less efficient. For
		// any types that must package does not support, consider
		// writing your own 2-line defer function, or use Any.
		var i = Any(strconv.Atoi("a")).(int)
		fmt.Println(i)

		return nil
	}()
	fmt.Println(err)
	// Output:
	// strconv.Atoi: parsing "a": invalid syntax
}

func ExampleLogErr() {
	// TODO: Fix the fragile test dependency to the error message.
	//
	// This is probably an anti-pattern because error was logged and
	// returned, handled twice.
	err := func() (err error) {
		defer ReturnErr(&err)
		// Log error to stderr with file name and line number.
		defer LogErr(log.Println)
		i := Int(strconv.Atoi("a"))
		fmt.Printf("i = %d\n", i)
		return nil
	}()
	fmt.Println(err)
	// Output:
	// strconv.Atoi: parsing "a": invalid syntax
}

func ExampleHandleErr_wrapError() {
	// TODO: Fix the fragile test dependency to the error message.
	//
	// This example shows how to wrap the error. You may use more
	// sophisticated wrappers like "github.com/pkg/errors".Wrap or multi
	// error appender.
	err := func() (err error) {
		defer HandleErr(func(newerr error) {
			err = fmt.Errorf("error occurred: %v", newerr)
		})
		fmt.Println(Int(strconv.Atoi("a")))
		return nil
	}()
	fmt.Println(err)
	// Output:
	// error occurred: strconv.Atoi: parsing "a": invalid syntax
}

func ExampleHandleErr_justReturnError() {
	// TODO: Fix the fragile test dependency to the error message.
	err := func() (err error) {
		// Equivalent with ReturnErr(&err). Param name is newerr for err
		// visibility.
		defer HandleErr(func(newerr error) {
			err = newerr
		})
		fmt.Println(Int(strconv.Atoi("a")))
		return nil
	}()
	fmt.Println(err)
	// Output:
	// strconv.Atoi: parsing "a": invalid syntax
}

func ExampleHandleErr_fromGoVersion2Draft() {
	// Translation of the code example of Go 2 error handling draft.
	// https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling-overview.md
	copyFile := func(src, dst string) (err error) {
		// Param name is newerr for err visibility.
		defer HandleErr(func(newerr error) {
			err = fmt.Errorf("copy %s %s: %v", src, dst, newerr)
		})

		r := Any(os.Open(src)).(*os.File)
		defer r.Close()

		w := Any(os.Create(dst)).(*os.File)
		// Note that this should be HandleErrNext(), not HandleErr().
		defer HandleErrNext(func(error) {
			w.Close()
			os.Remove(dst) // (only if a check fails)
		})

		Int64(io.Copy(w, r))
		Nil(w.Close())
		return nil
	}
	_ = copyFile
}

func ExampleHandleErr_fromGoVersion2DraftTestedSuccess() {
	// To demonstrate how it works when successful.
	copyFile := func(src, dst string) (err error) {
		// Param name is newerr for err visibility.
		defer HandleErr(func(newerr error) {
			err = fmt.Errorf("copy %s %s: %v", src, dst, newerr)
		})

		r := fmt.Sprintf("*File(%q)", src)
		fmt.Printf("os.Open(src) is called: %q\n", src)
		defer func() {
			fmt.Printf("r.Close() is called: %s\n", r)
		}()

		w := fmt.Sprintf("*File(%q)", dst)
		fmt.Printf("os.Create(dst) is called: %q\n", dst)
		// Note that this should be HandleErrNext(), not HandleErr().
		defer HandleErrNext(func(newerr error) {
			fmt.Printf("w.Close() is called in defer: %s\n", w)
			fmt.Printf("os.Remove(dst) is called: %q\n", dst)
		})

		fmt.Println("io.Copy(w, r) is called")
		fmt.Printf("w.Close() is called: %s\n", w)
		return nil
	}
	fmt.Println("returned error:", copyFile("input.txt", "output.txt"))
	// Output:
	// os.Open(src) is called: "input.txt"
	// os.Create(dst) is called: "output.txt"
	// io.Copy(w, r) is called
	// w.Close() is called: *File("output.txt")
	// r.Close() is called: *File("input.txt")
	// returned error: <nil>
}

func ExampleHandleErr_fromGoVersion2DraftTestedWriteFail() {
	// To demonstrate how it works when successful.
	copyFile := func(src, dst string) (err error) {
		// Param name is newerr for err visibility.
		defer HandleErr(func(newerr error) {
			err = fmt.Errorf("copy %s %s: %v", src, dst, newerr)
		})

		r := fmt.Sprintf("*File(%q)", src)
		fmt.Printf("os.Open(src) is called: %q\n", src)
		defer func() {
			fmt.Printf("r.Close() is called: %s\n", r)
		}()

		w := fmt.Sprintf("*File(%q)", dst)
		fmt.Printf("os.Create(dst) is called: %q\n", dst)
		// Note that this should be HandleErrNext(), not HandleErr().
		defer HandleErrNext(func(newerr error) {
			fmt.Printf("w.Close() is called in defer: %s\n", w)
			fmt.Printf("os.Remove(dst) is called: %q\n", dst)
		})

		fmt.Println("io.Copy(w, r) is called")
		Nil(errors.New("cannot write file"))
		fmt.Printf("w.Close() is called: %s\n", w)
		return nil
	}
	fmt.Println("returned error:", copyFile("input.txt", "output.txt"))
	// Output:
	// os.Open(src) is called: "input.txt"
	// os.Create(dst) is called: "output.txt"
	// io.Copy(w, r) is called
	// w.Close() is called in defer: *File("output.txt")
	// os.Remove(dst) is called: "output.txt"
	// r.Close() is called: *File("input.txt")
	// returned error: copy input.txt output.txt: cannot write file
}
