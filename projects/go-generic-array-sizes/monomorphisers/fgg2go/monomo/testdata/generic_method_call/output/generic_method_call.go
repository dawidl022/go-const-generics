package main

type any interface {
}

type Foo__2[T any] struct {
	x T
}

func (f Foo__2[T]) foo() T {
	return f.x
}

func main() {
	_ = Foo__2[int]{1}.foo()
}
