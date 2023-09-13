package main

type any interface {}

type Foo interface {
	something(x int, x any) int
}

func main() {
	_ = 0
}
