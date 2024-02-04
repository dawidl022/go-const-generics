package main

type any interface {
}

type Arr[T any, N const] [N]T

type Matrix[T any, N const, M const] [N]Arr[T, M]

func main() {
	_ = Matrix[int, 2, 3]{Arr[int, 3]{1, 2, 3}, Arr[int, 3]{4, 5, 6}}
}
