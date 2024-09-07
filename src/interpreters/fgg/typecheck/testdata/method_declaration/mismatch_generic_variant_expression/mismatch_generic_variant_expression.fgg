package main

type any interface {
}

type Foo[T any, R any] struct {
	x R
}

func (f Foo[T, R]) foo(param Foo[int, int]) Foo[T, R] {
	return param
}

func main() {
	_ = 1
}
