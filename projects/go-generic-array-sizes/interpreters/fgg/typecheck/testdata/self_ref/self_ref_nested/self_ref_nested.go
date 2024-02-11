package main

type any interface {
}

type Foo[T any] struct {
	bar T
}

type Bar[T any] struct {
	baz T
}

type Baz struct {
	foo Foo[Bar[Baz]]
}

func main() {
	_ = 0
}
