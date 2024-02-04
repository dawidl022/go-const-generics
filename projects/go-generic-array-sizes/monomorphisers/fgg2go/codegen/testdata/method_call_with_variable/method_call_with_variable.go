package main

type Foo struct {
}

func (f Foo) foo(x int) int {
	return x
}

func main() {
	_ = Foo{}.foo(1)
}
