package main

type any interface {
}

type Foo[T any, N const] struct {
}

type Arr[T any, N const] [N]T

type Fooer[T any, N const] interface {
	foo(x T) Arr[T, N]
}

func main() {
	_ = 0
}
