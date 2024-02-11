package main

type any interface {
}

type Foo[N const, T any] [N]T

type Bar[N const, T any] [N]Foo[N, T]

type Baz [2]Bar[2, Baz]

func main() {
	_ = 0
}
