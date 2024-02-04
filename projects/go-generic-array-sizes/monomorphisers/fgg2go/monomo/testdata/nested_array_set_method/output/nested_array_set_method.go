package main

type any interface {
}

type Arr__3[T any] [3]T

type Matrix__2__3[T any] [2]Arr__3[T]

func (m Matrix__2__3[T]) set(i int, val Arr__3[T]) Matrix__2__3[T] {
	m[i] = val;
	return m
}

func main() {
	_ = Matrix__2__3[int]{Arr__3[int]{1, 2, 3}, Arr__3[int]{4, 5, 6}}
}
