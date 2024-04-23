package main

type Arr[N const, T any] [N]T

func (a Arr[N, T]) Set(i int, x T) Arr[N, T] {
	a[i] = x;
	return a
}

func main() {
	_ = Arr[3, int]{1, 2, 3}.Set(3, 4)
}
