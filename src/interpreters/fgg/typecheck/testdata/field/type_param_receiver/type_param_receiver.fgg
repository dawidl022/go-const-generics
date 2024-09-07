package main

type any interface {
}

type Foo[T any] struct {
	x T
}

func (f Foo[T]) invalidField() T {
	return f.x.y
}

func main() {
	_ = 1
}
