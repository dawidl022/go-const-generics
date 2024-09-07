package main

type any interface {}

type Foo struct {
	x int
	x any
}

func main() {
	_ = 0
}
