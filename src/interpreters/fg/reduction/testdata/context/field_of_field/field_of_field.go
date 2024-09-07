package main

type Bar struct {
	x Foo
}

type Foo struct {
	y int
}

func main() {
	_ = Bar{Foo{1}}.x.y
}
