package main

type any interface {
}

type ArrFactory[T any, N const] interface {
	new() Arr[T, N]
}

type Arr[T any, N const] [N]T

func (a Arr[T, N]) new() Arr[T, N] {
	return a
}

type Matrix[T ArrFactory[E, M], E any, N const, M const] [N]T

func main() {
	_ = Matrix[Arr[int, 3], int, 2, 3]{Arr[int, 3]{1, 2, 3}, Arr[int, 3]{4, 5, 6}}
}
