package main

type Foo struct {
	x int
}

func (f Foo) getX() int {
	return Foo{2}.x
}

func main() {
	_ = Foo{1}.getX()
}
