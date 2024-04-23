package main

type any interface {
}

type Foo[T any] [2]T

type Bar[T any] [2]Foo[T]

type Baz [2]Bar[Baz]

func main() {
	_ = 0
}
