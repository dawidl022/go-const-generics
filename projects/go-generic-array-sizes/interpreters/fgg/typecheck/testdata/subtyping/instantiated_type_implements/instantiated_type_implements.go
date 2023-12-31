package main

type any interface{}

type Foo[T any] struct {
}

func (f Foo[T]) f(x T) T {
	return x
}

type myInterface interface {
	f(x int) int
}

type Bar struct {
}

func (b Bar) getInterface() myInterface {
	return Foo[int]{}
}

func main() {
	_ = Bar{}.getInterface()
}
