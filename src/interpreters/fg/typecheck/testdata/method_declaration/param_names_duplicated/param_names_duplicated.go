package main

type Foo struct {
}

func (f Foo) something(x int, x Foo) Foo {
	return f
}

func main() {
	_ = 0
}
