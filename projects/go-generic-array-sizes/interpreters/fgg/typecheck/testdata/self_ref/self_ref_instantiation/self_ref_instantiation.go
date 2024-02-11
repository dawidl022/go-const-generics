package main

type any interface {
}

type Foo[T any] struct {
	x T
}

func main() {
	_ = Foo[Foo[int]]{Foo[int]{1}}
}
