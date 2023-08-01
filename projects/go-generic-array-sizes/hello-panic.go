package main

type AnyArray0 [0]any

func (this AnyArray0) First() any {
	// when monomorphising, any constants indexing into a generic
	// should be first converted into non-constant ints
	//
	// this[0] would cause a compile-time error, whereas the below code panics
	index := 0
	return this[index]
}

func main() {
	_ = AnyArray0{}.First()
}
