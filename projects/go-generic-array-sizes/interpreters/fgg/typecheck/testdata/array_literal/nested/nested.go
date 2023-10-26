package main

type any interface {
}

type Arr[N const, T any] [N]T

type Matrix[N const, T any] [N]Arr[N, T]

func main() {
	_ = Matrix[2, int]{Arr[2, int]{1, 2}, Arr[2, int]{3, 4}}
}
