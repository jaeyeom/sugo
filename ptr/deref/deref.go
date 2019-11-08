// Package deref provides helper functions to dereference pointers.
//
// These functions are handy when you need to dereference a pointer to a value
// taking care of nil pointers.
package deref

// BoolOr returns a dereferenced value or the given default value if p is nil.
func BoolOr(p *bool, defVal bool) bool {
	if p == nil {
		return defVal
	}
	return *p
}

// Bool returns a dereferenced value or the zero value if p is nil.
func Bool(p *bool) bool {
	var defVal bool
	return BoolOr(p, defVal)
}

// StringOr returns a dereferenced value or the given default value if p is nil.
func StringOr(p *string, defVal string) string {
	if p == nil {
		return defVal
	}
	return *p
}

// String returns a dereferenced value or the zero value if p is nil.
func String(p *string) string {
	var defVal string
	return StringOr(p, defVal)
}

// IntOr returns a dereferenced value or the given default value if p is nil.
func IntOr(p *int, defVal int) int {
	if p == nil {
		return defVal
	}
	return *p
}

// Int returns a dereferenced value or the zero value if p is nil.
func Int(p *int) int {
	var defVal int
	return IntOr(p, defVal)
}

// Int8Or returns a dereferenced value or the given default value if p is nil.
func Int8Or(p *int8, defVal int8) int8 {
	if p == nil {
		return defVal
	}
	return *p
}

// Int8 returns a dereferenced value or the zero value if p is nil.
func Int8(p *int8) int8 {
	var defVal int8
	return Int8Or(p, defVal)
}

// Int16Or returns a dereferenced value or the given default value if p is nil.
func Int16Or(p *int16, defVal int16) int16 {
	if p == nil {
		return defVal
	}
	return *p
}

// Int16 returns a dereferenced value or the zero value if p is nil.
func Int16(p *int16) int16 {
	var defVal int16
	return Int16Or(p, defVal)
}

// Int32Or returns a dereferenced value or the given default value if p is nil.
func Int32Or(p *int32, defVal int32) int32 {
	if p == nil {
		return defVal
	}
	return *p
}

// Int32 returns a dereferenced value or the zero value if p is nil.
func Int32(p *int32) int32 {
	var defVal int32
	return Int32Or(p, defVal)
}

// Int64Or returns a dereferenced value or the given default value if p is nil.
func Int64Or(p *int64, defVal int64) int64 {
	if p == nil {
		return defVal
	}
	return *p
}

// Int64 returns a dereferenced value or the zero value if p is nil.
func Int64(p *int64) int64 {
	var defVal int64
	return Int64Or(p, defVal)
}

// UintOr returns a dereferenced value or the given default value if p is nil.
func UintOr(p *uint, defVal uint) uint {
	if p == nil {
		return defVal
	}
	return *p
}

// Uint returns a dereferenced value or the zero value if p is nil.
func Uint(p *uint) uint {
	var defVal uint
	return UintOr(p, defVal)
}

// Uint8Or returns a dereferenced value or the given default value if p is nil.
func Uint8Or(p *uint8, defVal uint8) uint8 {
	if p == nil {
		return defVal
	}
	return *p
}

// Uint8 returns a dereferenced value or the zero value if p is nil.
func Uint8(p *uint8) uint8 {
	var defVal uint8
	return Uint8Or(p, defVal)
}

// Uint16Or returns a dereferenced value or the given default value if p is nil.
func Uint16Or(p *uint16, defVal uint16) uint16 {
	if p == nil {
		return defVal
	}
	return *p
}

// Uint16 returns a dereferenced value or the zero value if p is nil.
func Uint16(p *uint16) uint16 {
	var defVal uint16
	return Uint16Or(p, defVal)
}

// Uint32Or returns a dereferenced value or the given default value if p is nil.
func Uint32Or(p *uint32, defVal uint32) uint32 {
	if p == nil {
		return defVal
	}
	return *p
}

// Uint32 returns a dereferenced value or the zero value if p is nil.
func Uint32(p *uint32) uint32 {
	var defVal uint32
	return Uint32Or(p, defVal)
}

// Uint64Or returns a dereferenced value or the given default value if p is nil.
func Uint64Or(p *uint64, defVal uint64) uint64 {
	if p == nil {
		return defVal
	}
	return *p
}

// Uint64 returns a dereferenced value or the zero value if p is nil.
func Uint64(p *uint64) uint64 {
	var defVal uint64
	return Uint64Or(p, defVal)
}

// UintptrOr returns a dereferenced value or the given default value if p is nil.
func UintptrOr(p *uintptr, defVal uintptr) uintptr {
	if p == nil {
		return defVal
	}
	return *p
}

// Uintptr returns a dereferenced value or the zero value if p is nil.
func Uintptr(p *uintptr) uintptr {
	var defVal uintptr
	return UintptrOr(p, defVal)
}

// ByteOr returns a dereferenced value or the given default value if p is nil.
func ByteOr(p *byte, defVal byte) byte {
	if p == nil {
		return defVal
	}
	return *p
}

// Byte returns a dereferenced value or the zero value if p is nil.
func Byte(p *byte) byte {
	var defVal byte
	return ByteOr(p, defVal)
}

// RuneOr returns a dereferenced value or the given default value if p is nil.
func RuneOr(p *rune, defVal rune) rune {
	if p == nil {
		return defVal
	}
	return *p
}

// Rune returns a dereferenced value or the zero value if p is nil.
func Rune(p *rune) rune {
	var defVal rune
	return RuneOr(p, defVal)
}

// Float32Or returns a dereferenced value or the given default value if p is nil.
func Float32Or(p *float32, defVal float32) float32 {
	if p == nil {
		return defVal
	}
	return *p
}

// Float32 returns a dereferenced value or the zero value if p is nil.
func Float32(p *float32) float32 {
	var defVal float32
	return Float32Or(p, defVal)
}

// Float64Or returns a dereferenced value or the given default value if p is nil.
func Float64Or(p *float64, defVal float64) float64 {
	if p == nil {
		return defVal
	}
	return *p
}

// Float64 returns a dereferenced value or the zero value if p is nil.
func Float64(p *float64) float64 {
	var defVal float64
	return Float64Or(p, defVal)
}

// Complex64Or returns a dereferenced value or the given default value if p is nil.
func Complex64Or(p *complex64, defVal complex64) complex64 {
	if p == nil {
		return defVal
	}
	return *p
}

// Complex64 returns a dereferenced value or the zero value if p is nil.
func Complex64(p *complex64) complex64 {
	var defVal complex64
	return Complex64Or(p, defVal)
}

// Complex128Or returns a dereferenced value or the given default value if p is nil.
func Complex128Or(p *complex128, defVal complex128) complex128 {
	if p == nil {
		return defVal
	}
	return *p
}

// Complex128 returns a dereferenced value or the zero value if p is nil.
func Complex128(p *complex128) complex128 {
	var defVal complex128
	return Complex128Or(p, defVal)
}
