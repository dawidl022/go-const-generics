package main

type any interface{}

type Array[N const, T any] [N]T

func (this Array[N, T]) First() T {
	return this[0]
}

func main() {
	_ = Array[2, int]{1, 2}.First()
}
