// Package ref provides helper functions to create pointer to a value.
package ref

// Bool returns a pointer to v.
func Bool(v bool) *bool {
	return &v
}

// String returns a pointer to v.
func String(v string) *string {
	return &v
}

// Int returns a pointer to v.
func Int(v int) *int {
	return &v
}

// Int8 returns a pointer to v.
func Int8(v int8) *int8 {
	return &v
}

// Int16 returns a pointer to v.
func Int16(v int16) *int16 {
	return &v
}

// Int32 returns a pointer to v.
func Int32(v int32) *int32 {
	return &v
}

// Int64 returns a pointer to v.
func Int64(v int64) *int64 {
	return &v
}

// Uint returns a pointer to v.
func Uint(v uint) *uint {
	return &v
}

// Uint8 returns a pointer to v.
func Uint8(v uint8) *uint8 {
	return &v
}

// Uint16 returns a pointer to v.
func Uint16(v uint16) *uint16 {
	return &v
}

// Uint32 returns a pointer to v.
func Uint32(v uint32) *uint32 {
	return &v
}

// Uint64 returns a pointer to v.
func Uint64(v uint64) *uint64 {
	return &v
}

// Uintptr returns a pointer to v.
func Uintptr(v uintptr) *uintptr {
	return &v
}

// Byte returns a pointer to v.
func Byte(v byte) *byte {
	return &v
}

// Rune returns a pointer to v.
func Rune(v rune) *rune {
	return &v
}

// Float32 returns a pointer to v.
func Float32(v float32) *float32 {
	return &v
}

// Float64 returns a pointer to v.
func Float64(v float64) *float64 {
	return &v
}

// Complex64 returns a pointer to v.
func Complex64(v complex64) *complex64 {
	return &v
}

// Complex128 returns a pointer to v.
func Complex128(v complex128) *complex128 {
	return &v
}
