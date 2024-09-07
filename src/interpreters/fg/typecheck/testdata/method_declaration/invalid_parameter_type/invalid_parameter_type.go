package main

type Foo struct {
}

func (f Foo) something(x int, y Bar) Foo {
	return f
}

func main() {
	_ = 0
}
