package main

type Foo struct {
	x int
	y int
	z int
}

func main() {
	_ = Foo{1, 2, 3}.y
}
