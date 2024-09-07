package main

type Foo[T any] struct {
}

func (f Foo[T]) foo(x int, y T) Foo[T] {
	return Foo[T]{}
}

func main() {
	_ = 0
}
