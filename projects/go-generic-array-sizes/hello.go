package main

type any interface{}

type AnyArray2 [2]any

func (this AnyArray2) First() any {
	return this[0]
}

func main() {
	_ = AnyArray2{1, 2}.First()
}
