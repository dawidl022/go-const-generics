package main

type any interface {
}

type Foo[T any] struct {
	x T
}

func (f Foo[T]) get() int {
	return Foo[int]{1}.x
}

func main() {
	_ = 1
}
