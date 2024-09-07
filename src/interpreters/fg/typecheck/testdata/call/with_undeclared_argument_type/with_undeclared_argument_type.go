package main

type any interface {
}

type Foo struct {
}

func (f Foo) answer(x int, y any) any {
	return y
}

func main() {
	_ = Foo{}.answer(1, Bar{})
}
