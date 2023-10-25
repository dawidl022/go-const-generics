package main

type any interface {
}

type Arr[N const, T any] [N]T

func (a Arr[N, T]) set(i int, v T) Arr[N, int] {
	a[i] = v;
	return a
}

func main() {
	_ = 0
}
