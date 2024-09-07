package main

type any interface {
}

type Arr[N const, T any] [N]T

type Foo[N const, M const] interface {
	foo(a Arr[N, int]) Arr[M, int]
}

type FooImpl[N const, M const] struct {
	x Arr[M, int]
}

func (f FooImpl[N, M]) foo(a Arr[N, int]) Arr[M, int] {
	return f.x
}

func (f FooImpl[N, M]) getFoo() Foo[N, M] {
	return f
}

func (f FooImpl[N, M]) makeSwappedFoo(a Arr[N, int]) Foo[M, N] {
	return FooImpl[M, N]{a}
}

func main() {
	_ = FooImpl[2, 3]{Arr[3, int]{1, 2, 3}}.foo(Arr[2, int]{1, 2})
}
