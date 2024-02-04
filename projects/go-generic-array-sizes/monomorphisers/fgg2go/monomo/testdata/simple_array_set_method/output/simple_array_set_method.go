package main

type any interface {
}

type Arr__2[T any] [2]T

func (a Arr__2[T]) set(i int, val T) Arr__2[T] {
	a[i] = val;
	return a
}

func main() {
	_ = Arr__2[int]{1, 2}.set(1, 10)
}
