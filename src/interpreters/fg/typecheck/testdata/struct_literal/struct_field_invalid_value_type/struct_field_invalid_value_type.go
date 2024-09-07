package main

type any interface {
}

type Foo struct {
	x int
	y any
}

func main() {
	_ = Foo{1, Bar{}}
}
