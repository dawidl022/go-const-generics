package main

type any interface {
}

type Foo[T any, R any] struct {
	x R
}

func (f Foo[T, R]) foo(param Foo) R {
	return param.x
}

func main() {
	_ = 1
}
