package main

type any interface {
}

type Arr[N const, T any] [N]T

type Foo[N const] struct {
	x Arr[N, int]
}

func main() {
	_ = Foo[2]{Arr[2, int]{1, 2}}
}
