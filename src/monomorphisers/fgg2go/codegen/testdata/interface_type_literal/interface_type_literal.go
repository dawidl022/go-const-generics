package main

type Fooer[T any] interface {
	foo(x int) T
}

func main() {
	_ = 0
}
