package main

type any interface {
}

type Foo[T any] interface {
}

type Arr[T Foo] [2]T

func main() {
	_ = 1
}
