package main

type Foo struct {
	x int
}

func (f Foo) getX() int {
	return f.x
}

type Arr [2]Foo

func main() {
	_ = Arr{Foo{1}, Foo{2}}[1].getX()
}
