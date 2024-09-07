package main

type Structure struct {
	x int
	y int
}

type Foo struct {
	x Structure
}

func main() {
	_ = Foo{Structure{1, 2}}.x
}
