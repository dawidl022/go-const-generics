package main

type any interface {
}

type Foo[N const, T any] struct {
	x T
}

func (f Foo[N, T]) foo() T {
	return f.x
}

func main() {
	_ = Foo[2, int]{1}.foo()
}
