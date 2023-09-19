package main

type Foo struct {
	x int
}

func main() {
	_ = Foo{1}.x
}
