package main

type any interface {
}

type Foo struct {
	x int
}

func (f Foo) getX() any {
	return f.x
}

func main() {
	_ = 0
}
