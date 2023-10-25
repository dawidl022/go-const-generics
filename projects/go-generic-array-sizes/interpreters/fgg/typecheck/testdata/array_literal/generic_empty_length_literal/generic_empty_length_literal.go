package main

type any interface {
}

type Arr[N const, T any] [N]T

func (a Arr[N, T]) new() Arr[N, T] {
	return Arr[N, T]{}
}

func main() {
	_ = 1
}
