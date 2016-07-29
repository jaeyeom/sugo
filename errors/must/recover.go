package must

type wrap struct {
	err error
}

// ReturnErr is a defer function to simplify returning errors. The pointer to
// the returning error variable should be passed. Errors captured by must
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
