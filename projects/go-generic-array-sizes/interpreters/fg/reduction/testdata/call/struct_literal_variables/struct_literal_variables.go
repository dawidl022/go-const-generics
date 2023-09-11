package main

type Foo struct {
	x int
	y int
}

func (f Foo) update(x int, y int) Foo {
	return Foo{x, y}
}

func main() {
	_ = Foo{1, 2}.update(3, 4)
}
