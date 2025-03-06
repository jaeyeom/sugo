// Binary mustgen generates functions for must package.
package main

import (
	"fmt"
	"strings"
)

type typeNames []string

func (ts typeNames) funcName() (name string) {
	for _, t := range ts {
		if strings.HasPrefix(string(t), "[]") {
			t = t[2:] + "s"
		}
		name += strings.ToUpper(string(t)[:1]) + string(t)[1:]
	}
	return name
}

func (ts typeNames) argList() string {
	args := make([]string, len(ts))
	for i, t := range ts {
		args[i] = fmt.Sprintf("v%d %s", i, t)
	}
	return strings.Join(args, ", ")
}

func (ts typeNames) returnTypeList() string {
	if len(ts) == 1 {
		return ts[0]
	}
	return "(" + strings.Join([]string(ts), ", ") + ")"
}

func (ts typeNames) returnList() string {
	args := make([]string, len(ts))
	for i := range ts {
		args[i] = fmt.Sprintf("v%d", i)
	}
	return strings.Join(args, ", ")
}

var types = []typeNames{
	{"bool"},
	{"string"},
	{"int"},
	{"int8"},
	{"int16"},
	{"int32"},
	{"int64"},
	{"uint"},
	{"uint8"},
	{"uint16"},
	{"uint32"},
	{"uint64"},
	{"uintptr"},
	{"byte"},
	{"rune"},
	{"float32"},
	{"float64"},
	{"complex64"},
	{"complex128"},
	{"[]bool"},
	{"[]string"},
	{"[]int"},
	{"[]int8"},
	{"[]int16"},
	{"[]int32"},
	{"[]int64"},
	{"[]uint"},
	{"[]uint8"},
	{"[]uint16"},
	{"[]uint32"},
	{"[]uint64"},
	{"[]uintptr"},
	{"[]byte"},
	{"[]rune"},
	{"[]float32"},
	{"[]float64"},
	{"[]complex64"},
	{"[]complex128"},
	{"rune", "int"},
	{"rune", "bool", "string"},
	{"int", "[]byte"},
	{"[]byte", "bool"},
}

func main() {
	for _, ts := range types {
		fmt.Printf("// %s panics if err is non-nil and returns %s.\n", ts.funcName(), ts.returnTypeList())
		fmt.Printf("func %s(%s, err error) %s {\n", ts.funcName(), ts.argList(), ts.returnTypeList())
		fmt.Println("\tCheckErr(err, 2)")
		fmt.Printf("\treturn %s\n", ts.returnList())
		fmt.Println("}")
		fmt.Println()
	}
}
