package main

type any interface {
}

type Foo struct {
	x any
}

func (f Foo) getX() int {
	return Foo{1}.x
}

func main() {
	_ = 0
}
