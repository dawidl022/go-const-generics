package main

type Foo struct {
}

func (f Foo) recurse(x int, y int) int {
	return f.recurse(x, y)
}

func main() {
	_ = Foo{}.recurse(1, 2)
}
