package main

type any interface {
}

type Arr[N const, T any] [N]T

type Matrix[N const, T any] [N]Arr[Foo, T]

func main() {
	_ = 1
}