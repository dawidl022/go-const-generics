package main

type any interface {
}

type Foo[T any, N const] [N]T

type Bar[T any, N const] [N]T

type Baz[N const] [N]Foo[Bar[Baz[N], N], N]

func main() {
	_ = 0
}
