package main


type Foo struct {
	x int
	y int
}

func main() {
	_ = Foo{Foo{1, 2}.x, Foo{1, 2}.y}
}
