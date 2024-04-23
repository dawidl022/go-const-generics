package main

type any interface {}

type Foo[T any] struct {
}

func (f Foo[T]) foo() int {
 	return Foo[T]{}.foo() + Foo[T[T]]{}.foo()
}

func main() {
	_ = 0
}
