package main

type any interface {
}

type Foo[T Bar] interface {
}

type Bar interface {
	foo() Baz[int]
}

type Baz[T Foo[Bar]] interface {
}

func main() {
	_ = 0
}
