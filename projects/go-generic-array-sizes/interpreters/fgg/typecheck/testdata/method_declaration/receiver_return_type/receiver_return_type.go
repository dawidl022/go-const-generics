package main

type any interface {
}

type Foo[T any, R any] struct {
	x R
}

func (f Foo[T, R]) foo() Foo[T, R] {
	return f
}

func main() {
	_ = 1
}
