package main

type any interface {
}

type Foo[T any] [2]T

type Bar[T any] [2]T

type Baz [3]Foo[Bar[Baz]]

func main() {
	_ = 0
}
