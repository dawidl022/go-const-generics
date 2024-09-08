package main

type any interface{}

type Array[N const, T any] [N]T

func (this Array[N, T]) Get(i int) T {
	return this[i]
}

func main() {
	_ = Array[2, int]{1, 2}.Get(0)   // 1
}
