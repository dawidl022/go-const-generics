package main

type Foo struct {
}

func (f Foo) firstArg(x int, y int, z int) int {
	return x
}

func (f Foo) unbound(x int) int {
	return f.firstArg(x, y, z)
}

func main() {
	_ = Foo{}.unbound(1)
}
