package main

type any interface{}

type Foo struct {
	x any
}

func main() {
	_ = Foo{1}
}
