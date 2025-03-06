// Binary refgen generates functions for ref and deref package.
package main

import (
	"fmt"
	"strings"
)

type typeName string

func (t typeName) funcName() string {
	return strings.ToUpper(string(t)[:1]) + string(t)[1:]
}

func (t typeName) String() string {
	return string(t)
}

var types = []typeName{
	"bool",
	"string",
	"int",
	"int8",
	"int16",
	"int32",
	"int64",
	"uint",
	"uint8",
	"uint16",
	"uint32",
	"uint64",
	"uintptr",
	"byte",
	"rune",
	"float32",
	"float64",
	"complex64",
	"complex128",
}

func generateRefs(ts []typeName) {
	for _, t := range ts {
		fmt.Printf("// %s returns a pointer to v.\n", t.funcName())
		fmt.Printf("func %s(v %s) *%s {\n", t.funcName(), t.String(), t.String())
		fmt.Printf("\treturn &v\n")
		fmt.Println("}")
		fmt.Println()
	}
}

func generateDerefs(ts []typeName) {
	for _, t := range ts {
		fmt.Printf("// %sOr returns a dereferenced value or the given default value if p is nil.\n", t.funcName())
		fmt.Printf("func %sOr(p *%s, defVal %s) %s {\n", t.funcName(), t.String(), t.String(), t.String())
		fmt.Println("\tif p == nil {")
		fmt.Println("\t\treturn defVal")
		fmt.Println("\t}")
		fmt.Println("\treturn *p")
		fmt.Println("}")
		fmt.Println()
		fmt.Printf("// %s returns a dereferenced value or the zero value if p is nil.\n", t.funcName())
		fmt.Printf("func %s(p *%s) %s {\n", t.funcName(), t.String(), t.String())
		fmt.Printf("\tvar defVal %s\n", t.String())
		fmt.Printf("\treturn %sOr(p, defVal)\n", t.funcName())
		fmt.Println("}")
		fmt.Println()
	}
}

func main() {
	generateRefs(types)
	generateDerefs(types)
}
