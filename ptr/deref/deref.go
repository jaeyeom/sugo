// Package deref provides helper functions to dereference pointers.
//
// Of and Or are the generic helpers and work for any type. Typed helpers such
// as Int and IntOr are kept for compatibility.
//
// These functions are handy when you need to dereference a pointer to a value
// taking care of nil pointers.
package deref

// Or returns a dereferenced value or the given default value if p is nil.
func Or[T any](p *T, defVal T) T {
	if p == nil {
		return defVal
	}
	return *p
}

// Of returns a dereferenced value or the zero value if p is nil.
func Of[T any](p *T) T {
	var defVal T
	return Or(p, defVal)
}

// BoolOr returns a dereferenced value or the given default value if p is nil.
func BoolOr(p *bool, defVal bool) bool {
	return Or(p, defVal)
}

// Bool returns a dereferenced value or the zero value if p is nil.
func Bool(p *bool) bool {
	return Of(p)
}

// StringOr returns a dereferenced value or the given default value if p is nil.
func StringOr(p *string, defVal string) string {
	return Or(p, defVal)
}

// String returns a dereferenced value or the zero value if p is nil.
func String(p *string) string {
	return Of(p)
}

// IntOr returns a dereferenced value or the given default value if p is nil.
func IntOr(p *int, defVal int) int {
	return Or(p, defVal)
}

// Int returns a dereferenced value or the zero value if p is nil.
func Int(p *int) int {
	return Of(p)
}

// Int8Or returns a dereferenced value or the given default value if p is nil.
func Int8Or(p *int8, defVal int8) int8 {
	return Or(p, defVal)
}

// Int8 returns a dereferenced value or the zero value if p is nil.
func Int8(p *int8) int8 {
	return Of(p)
}

// Int16Or returns a dereferenced value or the given default value if p is nil.
func Int16Or(p *int16, defVal int16) int16 {
	return Or(p, defVal)
}

// Int16 returns a dereferenced value or the zero value if p is nil.
func Int16(p *int16) int16 {
	return Of(p)
}

// Int32Or returns a dereferenced value or the given default value if p is nil.
func Int32Or(p *int32, defVal int32) int32 {
	return Or(p, defVal)
}

// Int32 returns a dereferenced value or the zero value if p is nil.
func Int32(p *int32) int32 {
	return Of(p)
}

// Int64Or returns a dereferenced value or the given default value if p is nil.
func Int64Or(p *int64, defVal int64) int64 {
	return Or(p, defVal)
}

// Int64 returns a dereferenced value or the zero value if p is nil.
func Int64(p *int64) int64 {
	return Of(p)
}

// UintOr returns a dereferenced value or the given default value if p is nil.
func UintOr(p *uint, defVal uint) uint {
	return Or(p, defVal)
}

// Uint returns a dereferenced value or the zero value if p is nil.
func Uint(p *uint) uint {
	return Of(p)
}

// Uint8Or returns a dereferenced value or the given default value if p is nil.
func Uint8Or(p *uint8, defVal uint8) uint8 {
	return Or(p, defVal)
}

// Uint8 returns a dereferenced value or the zero value if p is nil.
func Uint8(p *uint8) uint8 {
	return Of(p)
}

// Uint16Or returns a dereferenced value or the given default value if p is nil.
func Uint16Or(p *uint16, defVal uint16) uint16 {
	return Or(p, defVal)
}

// Uint16 returns a dereferenced value or the zero value if p is nil.
func Uint16(p *uint16) uint16 {
	return Of(p)
}

// Uint32Or returns a dereferenced value or the given default value if p is nil.
func Uint32Or(p *uint32, defVal uint32) uint32 {
	return Or(p, defVal)
}

// Uint32 returns a dereferenced value or the zero value if p is nil.
func Uint32(p *uint32) uint32 {
	return Of(p)
}

// Uint64Or returns a dereferenced value or the given default value if p is nil.
func Uint64Or(p *uint64, defVal uint64) uint64 {
	return Or(p, defVal)
}

// Uint64 returns a dereferenced value or the zero value if p is nil.
func Uint64(p *uint64) uint64 {
	return Of(p)
}

// UintptrOr returns a dereferenced value or the given default value if p is nil.
func UintptrOr(p *uintptr, defVal uintptr) uintptr {
	return Or(p, defVal)
}

// Uintptr returns a dereferenced value or the zero value if p is nil.
func Uintptr(p *uintptr) uintptr {
	return Of(p)
}

// ByteOr returns a dereferenced value or the given default value if p is nil.
func ByteOr(p *byte, defVal byte) byte {
	return Or(p, defVal)
}

// Byte returns a dereferenced value or the zero value if p is nil.
func Byte(p *byte) byte {
	return Of(p)
}

// RuneOr returns a dereferenced value or the given default value if p is nil.
func RuneOr(p *rune, defVal rune) rune {
	return Or(p, defVal)
}

// Rune returns a dereferenced value or the zero value if p is nil.
func Rune(p *rune) rune {
	return Of(p)
}

// Float32Or returns a dereferenced value or the given default value if p is nil.
func Float32Or(p *float32, defVal float32) float32 {
	return Or(p, defVal)
}

// Float32 returns a dereferenced value or the zero value if p is nil.
func Float32(p *float32) float32 {
	return Of(p)
}

// Float64Or returns a dereferenced value or the given default value if p is nil.
func Float64Or(p *float64, defVal float64) float64 {
	return Or(p, defVal)
}

// Float64 returns a dereferenced value or the zero value if p is nil.
func Float64(p *float64) float64 {
	return Of(p)
}

// Complex64Or returns a dereferenced value or the given default value if p is nil.
func Complex64Or(p *complex64, defVal complex64) complex64 {
	return Or(p, defVal)
}

// Complex64 returns a dereferenced value or the zero value if p is nil.
func Complex64(p *complex64) complex64 {
	return Of(p)
}

// Complex128Or returns a dereferenced value or the given default value if p is nil.
func Complex128Or(p *complex128, defVal complex128) complex128 {
	return Or(p, defVal)
}

// Complex128 returns a dereferenced value or the zero value if p is nil.
func Complex128(p *complex128) complex128 {
	return Of(p)
}
