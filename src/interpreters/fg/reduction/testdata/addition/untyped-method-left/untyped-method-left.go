package main

type Foo struct {
}

func (f Foo) foo() int {
	return x + 1
}

func main() {
	_ = Foo{}.foo() + 1
}
