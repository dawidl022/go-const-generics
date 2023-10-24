package main

type any interface {
}

type Foo[T any] interface {
}

type Arr[T any, F Foo[T]] [2]F

func main() {
	_ = 1
}
