package main

type any interface {
}

type Useless__1 interface {
}

type Arr__2[T any] [2]T

type Arr__3[T any] [3]T

type Arr__1[T any] [1]T

func (a Arr__2[T]) zero() int {
	return 0
}

func (a Arr__3[T]) zero() int {
	return 0
}

func (a Arr__1[T]) zero() int {
	return 0
}

type Foo[T any] struct {
}

func (f Foo[T]) foo(a Arr__2[T], b Arr__3[T]) Useless__1 {
	return Arr__1[int]{1}
}

func main() {
	_ = Foo[int]{}
}
