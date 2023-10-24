package main

type any interface {
}

type Arr[N const, T any] [N]T

type Matrix[N const, T any] [N]Arr[N, N]

func main() {
	_ = 1
}
