package ranger

import "fmt"

func ExampleNDShape_ToIth() {
	data := []int{1, 2, 3, 4, 5, 6}
	// 2 rows and 3 cols
	rows, cols := 2, 3
	nds := NDShape(ND{rows, cols})
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("(%d)", data[nds.ToIth(ND{i, j})])
		}
		fmt.Println()
	}
	// Output:
	// (1)(2)(3)
	// (4)(5)(6)
}

func ExampleNDShape_FromIth() {
	rows, cols := 2, 3
	nds := NDShape(ND{rows, cols})
	for i := 0; i < rows*cols; i++ {
		fmt.Println(nds.FromIth(i))
	}
	// Output:
	// [0 0]
	// [0 1]
	// [0 2]
	// [1 0]
	// [1 1]
	// [1 2]
}

func ExampleNDRange_ToIth() {
	data := []int{
		0, 0, 0, 0, 0,
		0, 0, 0, 0, 0,
		0, 0, 0, 0, 0,
		0, 0, 0, 0, 0,
	}
	// The input data shape
	rows, cols := 4, 5
	nds := NDShape(ND{rows, cols})
	// Now I'd like to slice only inside 2x3 area.
	rowr, colr := Range(1, 3, 1), Range(1, 4, 1)
	ndr := NDRange{Shape: nds, Iths: []FiniteIth{rowr, colr}}
	for i := 0; i < rowr.Size; i++ {
		for j := 0; j < colr.Size; j++ {
			data[ndr.ToIth(ND{i, j})] = 1
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("(%d)", data[nds.ToIth(ND{i, j})])
		}
		fmt.Println()
	}
	// Output:
	// (0)(0)(0)(0)(0)
	// (0)(1)(1)(1)(0)
	// (0)(1)(1)(1)(0)
	// (0)(0)(0)(0)(0)
}

func ExampleNDRange_FromIth() {
	rows, cols := 4, 5
	nds := NDShape(ND{rows, cols})
	rowr, colr := Range(1, 3, 1), Range(1, 4, 1)
	ndr := NDRange{Shape: nds, Iths: []FiniteIth{rowr, colr}}
	for i := 0; i < ndr.ComputeSize(); i++ {
		fmt.Println(ndr.FromIth(i))
	}
	// Output:
	// [1 1]
	// [1 2]
	// [1 3]
	// [2 1]
	// [2 2]
	// [2 3]
}
