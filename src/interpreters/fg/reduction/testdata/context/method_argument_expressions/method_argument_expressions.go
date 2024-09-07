package main

type Arr [1]int

type Foo struct {
	x int
	y int
	z int
}

func (f Foo) update(x int, y int, z int) Foo {
	return Foo{x, y, z}
}

func main() {
	_ = Foo{1, 2, 3}.update(4, Arr{5}[0], Arr{6}[0])
}
