package main

type Arr[N const, T any] [N]T

type Matrix[N const, T any] [N]Arr[N, Foo]

func main() {
	_ = 1
}
