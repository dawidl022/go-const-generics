package main

type any interface {
}

type Baz[T Foo[Bar]] interface {
}

type Foo[T Bar] interface {
}

type Bar interface {
	foo() Baz[int]
}

func main() {
	_ = 0
}
