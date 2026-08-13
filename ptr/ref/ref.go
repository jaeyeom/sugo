// Package ref provides helper functions to create pointer to a value.
//
// Of is the generic helper and works for any type. Typed helpers such as Int
// and String are kept for compatibility.
//
// These functions are handy whenever you need a pointer to a primitive value.
// It's easy to create a pointer to a struct, but there's no simple way to do
// that for numbers, strings, etc. Protocol buffer Go API has these functions,
// too.
package ref

// Of returns a pointer to v.
func Of[T any](v T) *T {
	return &v
}

// Bool returns a pointer to v.
func Bool(v bool) *bool {
	return Of(v)
}

// String returns a pointer to v.
func String(v string) *string {
	return Of(v)
}

// Int returns a pointer to v.
func Int(v int) *int {
	return Of(v)
}

// Int8 returns a pointer to v.
func Int8(v int8) *int8 {
	return Of(v)
}

// Int16 returns a pointer to v.
func Int16(v int16) *int16 {
	return Of(v)
}

// Int32 returns a pointer to v.
func Int32(v int32) *int32 {
	return Of(v)
}

// Int64 returns a pointer to v.
func Int64(v int64) *int64 {
	return Of(v)
}

// Uint returns a pointer to v.
func Uint(v uint) *uint {
	return Of(v)
}

// Uint8 returns a pointer to v.
func Uint8(v uint8) *uint8 {
	return Of(v)
}

// Uint16 returns a pointer to v.
func Uint16(v uint16) *uint16 {
	return Of(v)
}

// Uint32 returns a pointer to v.
func Uint32(v uint32) *uint32 {
	return Of(v)
}

// Uint64 returns a pointer to v.
func Uint64(v uint64) *uint64 {
	return Of(v)
}

// Uintptr returns a pointer to v.
func Uintptr(v uintptr) *uintptr {
	return Of(v)
}

// Byte returns a pointer to v.
func Byte(v byte) *byte {
	return Of(v)
}

// Rune returns a pointer to v.
func Rune(v rune) *rune {
	return Of(v)
}

// Float32 returns a pointer to v.
func Float32(v float32) *float32 {
	return Of(v)
}

// Float64 returns a pointer to v.
func Float64(v float64) *float64 {
	return Of(v)
}

// Complex64 returns a pointer to v.
func Complex64(v complex64) *complex64 {
	return Of(v)
}

// Complex128 returns a pointer to v.
func Complex128(v complex128) *complex128 {
	return Of(v)
}
