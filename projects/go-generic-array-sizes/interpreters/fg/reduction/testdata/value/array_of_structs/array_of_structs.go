package main

type Foo struct {
	x int
	y int
}

type Arr [2]Foo

func main() {
	_ = Arr{Foo{1, 2}, Foo{3, 4}}
}
