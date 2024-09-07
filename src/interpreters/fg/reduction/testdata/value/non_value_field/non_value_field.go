package main

type Foo struct {
	x int
}

func main() {
	_ = Foo{Foo{1}.x}
}
