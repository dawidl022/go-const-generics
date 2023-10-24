package main

type any interface {
}

type Foo[T any, T any] struct {
}

func main() {
	_ = 1
}
