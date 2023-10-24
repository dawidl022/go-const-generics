package main

type any interface {
}

type fooer interface {
	foo() int
}

type Arr[N const, T fooer] [N]T

type Foo[N const, T any] struct {
	x T
	arr Arr[N, T]
}

func main() {
	_ = 1
}
