package ranger

// ND represents multi dimentional position or size.
type ND []int

// NDShape is multi dimentional array shape.
type NDShape ND

// ToIth converts the given multi dimentional index to a one dimentional index.
func (nds NDShape) ToIth(nd ND) int {
	var i int
	for j := range nd {
		i = i*nds[j] + nd[j]
	}
	return i
}

// FromIth converts the one dimensional index i into N-dimentional index.
func (nds NDShape) FromIth(i int) ND {
	dim := len(nds)
	p := make([]int, dim)
	for j := dim - 1; j > -1; j-- {
		p[j] = i % nds[j]
		i /= nds[j]
	}
	return ND(p)
}

// NDRange is an N-dimensional rectangular range.
type NDRange struct {
	Shape NDShape
	Iths  []FiniteIth
}

// ComputeSize returns the number of elements in this range.
func (ndr NDRange) ComputeSize() int {
	size := 1
	for _, fi := range ndr.Iths {
		size *= fi.Size
	}
	return size
}

// ToIth converts the given multi dimentional index to a one dimentional index.
func (ndr NDRange) ToIth(nd ND) int {
	var i int
	for j := range nd {
		i = i*ndr.Shape[j] + ndr.Iths[j].Ith(nd[j])
	}
	return i
}

// FromIth converts the one dimensional index i into N-dimentional index.
func (ndr NDRange) FromIth(i int) ND {
	dim := len(ndr.Iths)
	p := make([]int, dim)
	for j := dim - 1; j > -1; j-- {
		p[j] = ndr.Iths[j].Ith(i % ndr.Iths[j].Size)
		i /= ndr.Iths[j].Size
	}
	return ND(p)
}
