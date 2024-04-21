package main

type Foo struct {
}

func (f Foo) foo(x int) int {
	return 1 + x
}

func main() {
	_ = 1 + Foo{}.foo(1)
}
