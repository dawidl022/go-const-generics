package main

type any interface {
}

type Foo[T any] struct {
	baz T
}

type Bar[T any] struct {
	foo Foo[T]
}

type Baz struct {
	bar Bar[Baz]
}

func main() {
	_ = 0
}
