package main

type Foo struct {
}

func (f Foo) firstArg(x int, y int) int {
	return x
}

func main() {
	_ = Foo{}.firstArg(1)
}
