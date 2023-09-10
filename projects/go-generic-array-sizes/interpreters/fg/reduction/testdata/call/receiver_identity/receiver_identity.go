package main

type Foo struct {
	x int
}

func (f Foo) Identity() int {
	return f
}

func main() {
	_ = Foo{1}.Identity()
}
