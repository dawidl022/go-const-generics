package main

type any interface {
}

type Foo[T any] struct {
	b T
}

type Bar[T any] struct {
	f Foo[T]
}

type Baz struct {
	f Bar[Baz]
}

func main() {
	_ = 0
}
