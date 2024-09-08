package main

type any interface {
}

type Identity[T any] interface{
	f(t T) T
}

type Foo[T any] struct {
	b Identity[T]
}

type Bar struct {
	f Foo[Bar]
}

func main() {
	_ = 0
}
