package main

type any interface {
}

type Arr[N const, T any] [N]T

func (a Arr[N, T]) get() T {
	return a[0]
}

func main() {
	_ = 0
}
