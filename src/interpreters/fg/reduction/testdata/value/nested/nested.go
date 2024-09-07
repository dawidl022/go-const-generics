package main

type Foo struct {
	x int
	y int
}

type Bar struct {
	x Foo
	y Foo
}

func main() {
	_ = Bar{Foo{1, 2}, Foo{3, 4}}
}
