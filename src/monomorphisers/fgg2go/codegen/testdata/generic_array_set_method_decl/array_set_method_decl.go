package main

type Arr[T any, N const] [N]T

func (a Arr[T, N]) set(i int, v T) Arr[T, N] {
	a[i] = v;
	return a
}

func main() {
	_ = 0
}
