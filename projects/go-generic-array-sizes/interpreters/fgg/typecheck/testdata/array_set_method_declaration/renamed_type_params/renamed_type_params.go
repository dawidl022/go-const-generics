package main

type any interface {
}

type Arr[N const, T any] [N]any

func (a Arr[X, Y]) set(i int, v T) Arr[N, T] {
	a[i] = v;
	return a
}

func main() {
	_ = 0
}
