package main

type Arr[N const, T any] [N]T

type Matrix[N const, T any] [N]Arr[T, T]

func main() {
	_ = 1
}
