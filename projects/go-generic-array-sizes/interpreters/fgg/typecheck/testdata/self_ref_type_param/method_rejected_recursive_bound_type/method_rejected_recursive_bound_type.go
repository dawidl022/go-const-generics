package main

type any interface {
}

type Bar interface {
	foo() Foo[Bar]
}

type Foo[T Bar] interface {
}

func main() {
	_ = 0
}
