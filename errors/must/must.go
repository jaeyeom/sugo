// Package must implements helper functions to simplify repetitive error handling.
//
// Mainly there are 3 ways to handle errors.
//
// - Return errors.
// - Panic.
// - Log errors and continue.
//
// This package provides a way to hide error handling and focus on the main
// logic. The code may look a lot cleaner. The downside is that the explicit
// return statement is hidden to look unclear.
package must

import "runtime"

// CheckErr panics if err is non-nil and put the filename and line number and
// skip is the number of stack fragments to ascend. This function could be used
// to write custom must functions.
func CheckErr(err error, skip int) {
	if err != nil {
		_, file, line, _ := runtime.Caller(2)
		panic(wrap{err, file, line})
	}
}

// Nil panics if err is non-nil.
func Nil(err error) {
	CheckErr(err, 2)
}

// Any panics if err is non-nil and returns interface{}.
func Any(v0 interface{}, err error) interface{} {
	CheckErr(err, 2)
	return v0
}

// Functions below are auto-generated by github.com/jaeyeom/sugo/cmd/mustgen.

// Bool panics if err is non-nil and returns bool.
func Bool(v0 bool, err error) bool {
	CheckErr(err, 2)
	return v0
}

// String panics if err is non-nil and returns string.
func String(v0 string, err error) string {
	CheckErr(err, 2)
	return v0
}

// Int panics if err is non-nil and returns int.
func Int(v0 int, err error) int {
	CheckErr(err, 2)
	return v0
}

// Int8 panics if err is non-nil and returns int8.
func Int8(v0 int8, err error) int8 {
	CheckErr(err, 2)
	return v0
}

// Int16 panics if err is non-nil and returns int16.
func Int16(v0 int16, err error) int16 {
	CheckErr(err, 2)
	return v0
}

// Int32 panics if err is non-nil and returns int32.
func Int32(v0 int32, err error) int32 {
	CheckErr(err, 2)
	return v0
}

// Int64 panics if err is non-nil and returns int64.
func Int64(v0 int64, err error) int64 {
	CheckErr(err, 2)
	return v0
}

// Uint panics if err is non-nil and returns uint.
func Uint(v0 uint, err error) uint {
	CheckErr(err, 2)
	return v0
}

// Uint8 panics if err is non-nil and returns uint8.
func Uint8(v0 uint8, err error) uint8 {
	CheckErr(err, 2)
	return v0
}

// Uint16 panics if err is non-nil and returns uint16.
func Uint16(v0 uint16, err error) uint16 {
	CheckErr(err, 2)
	return v0
}

// Uint32 panics if err is non-nil and returns uint32.
func Uint32(v0 uint32, err error) uint32 {
	CheckErr(err, 2)
	return v0
}

// Uint64 panics if err is non-nil and returns uint64.
func Uint64(v0 uint64, err error) uint64 {
	CheckErr(err, 2)
	return v0
}

// Uintptr panics if err is non-nil and returns uintptr.
func Uintptr(v0 uintptr, err error) uintptr {
	CheckErr(err, 2)
	return v0
}

// Byte panics if err is non-nil and returns byte.
func Byte(v0 byte, err error) byte {
	CheckErr(err, 2)
	return v0
}

// Rune panics if err is non-nil and returns rune.
func Rune(v0 rune, err error) rune {
	CheckErr(err, 2)
	return v0
}

// Float32 panics if err is non-nil and returns float32.
func Float32(v0 float32, err error) float32 {
	CheckErr(err, 2)
	return v0
}

// Float64 panics if err is non-nil and returns float64.
func Float64(v0 float64, err error) float64 {
	CheckErr(err, 2)
	return v0
}

// Complex64 panics if err is non-nil and returns complex64.
func Complex64(v0 complex64, err error) complex64 {
	CheckErr(err, 2)
	return v0
}

// Complex128 panics if err is non-nil and returns complex128.
func Complex128(v0 complex128, err error) complex128 {
	CheckErr(err, 2)
	return v0
}

// Bools panics if err is non-nil and returns []bool.
func Bools(v0 []bool, err error) []bool {
	CheckErr(err, 2)
	return v0
}

// Strings panics if err is non-nil and returns []string.
func Strings(v0 []string, err error) []string {
	CheckErr(err, 2)
	return v0
}

// Ints panics if err is non-nil and returns []int.
func Ints(v0 []int, err error) []int {
	CheckErr(err, 2)
	return v0
}

// Int8s panics if err is non-nil and returns []int8.
func Int8s(v0 []int8, err error) []int8 {
	CheckErr(err, 2)
	return v0
}

// Int16s panics if err is non-nil and returns []int16.
func Int16s(v0 []int16, err error) []int16 {
	CheckErr(err, 2)
	return v0
}

// Int32s panics if err is non-nil and returns []int32.
func Int32s(v0 []int32, err error) []int32 {
	CheckErr(err, 2)
	return v0
}

// Int64s panics if err is non-nil and returns []int64.
func Int64s(v0 []int64, err error) []int64 {
	CheckErr(err, 2)
	return v0
}

// Uints panics if err is non-nil and returns []uint.
func Uints(v0 []uint, err error) []uint {
	CheckErr(err, 2)
	return v0
}

// Uint8s panics if err is non-nil and returns []uint8.
func Uint8s(v0 []uint8, err error) []uint8 {
	CheckErr(err, 2)
	return v0
}

// Uint16s panics if err is non-nil and returns []uint16.
func Uint16s(v0 []uint16, err error) []uint16 {
	CheckErr(err, 2)
	return v0
}

// Uint32s panics if err is non-nil and returns []uint32.
func Uint32s(v0 []uint32, err error) []uint32 {
	CheckErr(err, 2)
	return v0
}

// Uint64s panics if err is non-nil and returns []uint64.
func Uint64s(v0 []uint64, err error) []uint64 {
	CheckErr(err, 2)
	return v0
}

// Uintptrs panics if err is non-nil and returns []uintptr.
func Uintptrs(v0 []uintptr, err error) []uintptr {
	CheckErr(err, 2)
	return v0
}

// Bytes panics if err is non-nil and returns []byte.
func Bytes(v0 []byte, err error) []byte {
	CheckErr(err, 2)
	return v0
}

// Runes panics if err is non-nil and returns []rune.
func Runes(v0 []rune, err error) []rune {
	CheckErr(err, 2)
	return v0
}

// Float32s panics if err is non-nil and returns []float32.
func Float32s(v0 []float32, err error) []float32 {
	CheckErr(err, 2)
	return v0
}

// Float64s panics if err is non-nil and returns []float64.
func Float64s(v0 []float64, err error) []float64 {
	CheckErr(err, 2)
	return v0
}

// Complex64s panics if err is non-nil and returns []complex64.
func Complex64s(v0 []complex64, err error) []complex64 {
	CheckErr(err, 2)
	return v0
}

// Complex128s panics if err is non-nil and returns []complex128.
func Complex128s(v0 []complex128, err error) []complex128 {
	CheckErr(err, 2)
	return v0
}

// RuneInt panics if err is non-nil and returns (rune, int).
func RuneInt(v0 rune, v1 int, err error) (rune, int) {
	CheckErr(err, 2)
	return v0, v1
}

// RuneBoolString panics if err is non-nil and returns (rune, bool, string).
func RuneBoolString(v0 rune, v1 bool, v2 string, err error) (rune, bool, string) {
	CheckErr(err, 2)
	return v0, v1, v2
}

// IntBytes panics if err is non-nil and returns (int, []byte).
func IntBytes(v0 int, v1 []byte, err error) (int, []byte) {
	CheckErr(err, 2)
	return v0, v1
}

// BytesBool panics if err is non-nil and returns ([]byte, bool).
func BytesBool(v0 []byte, v1 bool, err error) ([]byte, bool) {
	CheckErr(err, 2)
	return v0, v1
}
