package main

type any interface {
}

type Foo[T any] struct {
}

type T[T any] struct {
}

func (f Foo[T]) foo(x T) T {
	return x
}

func main() {
	_ = 0
}
