package main

type any interface {
}

type Arr[N const, T any] [N]T

func (a Arr[N, T]) getInt() int {
	return Arr[2, int]{1, 2}[1]
}

func main() {
	_ = 0
}
