package main

type any interface {
}

type Foo[T any, R any] struct {
	x R
}

func (f Foo[T, R]) foo(param T) R {
	return f.x
}

func main() {
	_ = 1
}
