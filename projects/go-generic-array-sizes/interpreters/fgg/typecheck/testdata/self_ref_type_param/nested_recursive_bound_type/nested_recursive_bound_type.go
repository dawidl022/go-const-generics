package main

type any interface {
}

type Foo[T any] interface {
}

type Eq[T Foo[Eq[T]]] interface {
	equal(other T) int
}

func main() {
	_ = 0
}
