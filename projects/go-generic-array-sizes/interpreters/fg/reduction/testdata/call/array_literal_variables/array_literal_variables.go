package main

type Arr [2] int

func (a Arr) update(x int, y int) Foo {
	return Arr{x, y}
}

func main() {
	_ = Arr{1, 2}.update(3, 4)
}
