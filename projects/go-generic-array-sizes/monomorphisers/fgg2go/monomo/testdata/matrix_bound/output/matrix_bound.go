package main

type any interface {
}

type ArrFactory__3[T any] interface {
	new() Arr__3[T]
}

type Arr__3[T any] [3]T

func (a Arr__3[T]) new() Arr__3[T] {
	return a
}

type Matrix__2__3[T ArrFactory__3[E], E any] [2]T

func main() {
	_ = Matrix__2__3[Arr__3[int], int]{Arr__3[int]{1, 2, 3}, Arr__3[int]{4, 5, 6}}
}
