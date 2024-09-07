package main

type any interface {
}

type Arr[N const, T any] [N]T

func (a Arr[N, T]) set(i int, val T) Arr[N, T] {
	a[i] = val;
	return a
}

func main() {
	_ = Arr[2, int]{1, 2}.set(1, 10)
}
