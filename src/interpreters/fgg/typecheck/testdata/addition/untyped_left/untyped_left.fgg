package main

type any interface {}

type Foo[T any] struct {
}

func (f Foo[T]) foo() int {
 	return Foo[T[T]]{}.foo() + Foo[T]{}.foo()
}

func main() {
	_ = 0
}
