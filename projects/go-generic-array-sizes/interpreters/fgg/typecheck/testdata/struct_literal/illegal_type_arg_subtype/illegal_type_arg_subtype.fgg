package main

type fooer interface {
	foo() int
}

type Foo[T fooer] struct {
	x T
}

func main() {
	_ = Foo[int]{1}
}
