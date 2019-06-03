package must

import (
	"fmt"
	"path/filepath"
)

type wrap struct {
	err  error
	file string
	line int
}

// Error returns an error string with file and line.
func (w wrap) Error() string {
	return fmt.Sprintf("%s:%d %v", filepath.Base(w.file), w.line, w.err)
}

// ReturnErr is a defer function to simplify returning errors. The pointer to
// the returning error variable perr should be passed. Errors captured by must
// package are handled. Other panic values won't be handled here.
func ReturnErr(perr *error) {
	if r := recover(); r != nil {
		if e, ok := r.(wrap); ok {
			*perr = e.err
		} else {
			panic(r)
		}
	}
}

// LogErr is a defer function to simplify logging.
func LogErr(logger func(...interface{})) {
	if r := recover(); r != nil {
		if e, ok := r.(wrap); ok {
			logger(e)
		}
		panic(r)
	}
}

// HandleErr is a defer function to customize error handling. This can be used
// in case the returning error is custom typed or logging is required.
func HandleErr(handler func(error)) {
	if r := recover(); r != nil {
		if e, ok := r.(wrap); ok {
			handler(e.err)
		} else {
			panic(r)
		}
	}
}

// HandleErrNext is a defer function to customize the following (2nd, 3rd, ...)
// error handling. This is similar to HandleErr except the panic is not
// consumed. Use this to let other deferred handlers above handle the error as
// well.
func HandleErrNext(handler func(error)) {
	if r := recover(); r != nil {
		if e, ok := r.(wrap); ok {
			handler(e.err)
		}
		panic(r)
	}
}
