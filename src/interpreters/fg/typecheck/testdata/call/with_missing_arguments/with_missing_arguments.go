package main

type Foo struct {
}

func (f Foo) answer(x int, y int) int {
	return y
}

func main() {
	_ = Foo{}.answer(1)
}
