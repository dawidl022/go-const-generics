package main

type Fooer[T any] interface {
	foo(x int) T
	b(x int, b Bar[T]) Bar[int]
}

type Bar[T any] struct {
}

func main() {
	_ = 0
}
