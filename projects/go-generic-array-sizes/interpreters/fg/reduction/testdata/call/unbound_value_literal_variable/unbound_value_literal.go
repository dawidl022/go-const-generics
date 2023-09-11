package main

type Foo struct {
	x int
	y int
}

func (f Foo) unbound() Foo {
	return Foo{x, y}
}

func main() {
	_ = Foo{1, 2}.unbound()
}
