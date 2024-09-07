package main

type any interface {
}

type Arr [N const, T any] [N]T

type Matrix[N const, M const, T any] [N]Arr[M, T]

func (m Matrix[N, M, T]) set(i int, val Arr[M, T]) Matrix[N, M, T] {
	m[i] = val;
	return m
}

func main() {
	_ = Matrix[2, 3, int]{Arr[3, int]{1, 2, 3}, Arr[3, int]{4, 5, 6}}
}
