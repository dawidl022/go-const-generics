package main

type Foo[T any] struct {
	x T
}

type Bar struct {
	x Foo[Bar]
}

func main() {
}
