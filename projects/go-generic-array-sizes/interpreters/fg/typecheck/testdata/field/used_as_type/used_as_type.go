package main

type Foo struct {
	x int
}

func (f Foo) getX() int {
	return f.x
}

func main() {
	_ = 0
}
