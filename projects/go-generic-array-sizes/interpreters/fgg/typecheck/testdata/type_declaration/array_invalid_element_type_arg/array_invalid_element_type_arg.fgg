package main

type any interface {
}

type fooer interface {
	foo() int
}

type Arr[N const, T fooer] [N]T

type Matrix[N const, T any] [N]Arr[N, T]

func main() {
	_ = 1
}
