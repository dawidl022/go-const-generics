package main

type any interface {
}

type Useless[N const] interface {
}

type Arr[N const, T any] [N]T

func (a Arr[N, T]) zero() int {
	return 0
}

type Foo[T any] struct {
}

func (f Foo[T]) foo(a Arr[2, T], b Arr[3, T]) Useless[1] {
	return Arr[1, int]{1}
}

func main() {
	_ = Foo[int]{}
}
