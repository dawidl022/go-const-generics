package main

type any interface {
}

type Arr[N const, T any] [N]T

type Foo[T any, N const] struct {
	x T
	arr Arr[N, T]
}

func main() {
	_ = Foo[int, 2]{1, Arr[2, int]{2, 3}}
}
