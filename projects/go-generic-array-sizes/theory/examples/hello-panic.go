package main

type AnyArray0 [0]any

func (this AnyArray0) First() any {
	// this[0] would cause a compile-time error, whereas the below code panics
	index := 0
	return this[index]
}

func main() {
	_ = AnyArray0{}.First()
}
