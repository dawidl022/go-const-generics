package main

type Foo[T any, R any] struct {
	x T
	y R
}

func main() {
	_ = Foo[int, int]{1, 2}
}
