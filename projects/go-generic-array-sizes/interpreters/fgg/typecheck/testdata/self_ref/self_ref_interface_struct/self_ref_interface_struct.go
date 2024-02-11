package main

type any interface {
}

type Fooer[T any] interface {
	foo() T
}

type Foo struct {
	x Fooer[Foo]
}

func main() {
	_ = 0
}
