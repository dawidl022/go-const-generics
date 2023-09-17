package main

type any interface {
}

type anyGetter interface {
	getAny(x int, y int) int
}

type Baz struct {
}

func (b Baz) getAny(x int, y int) int {
	return x
}

type Foo struct {
}

func (f Foo) getAny(x int, y Foo) int {
	return 42
}

type Bar struct {
	getter anyGetter
}

func main() {
	_ = Bar{Foo{}}
}
