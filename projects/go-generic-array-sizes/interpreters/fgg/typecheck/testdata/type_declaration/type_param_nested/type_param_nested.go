package main

type any interface {
}

type barer interface {
	bar() int
}

type fooer[T any] interface {
	foo() T
}

type Foo[T fooer[barer]] struct {
}

type Bar struct {
	x Foo[fooer[barer]]
}

func main() {
	_ = 1
}
