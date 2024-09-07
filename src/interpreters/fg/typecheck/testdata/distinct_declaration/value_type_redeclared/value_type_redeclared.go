package main

type Foo struct {
	x int
	y int
}

type Foo [2]int

func main() {
	_ = Foo{1, 2}
}
