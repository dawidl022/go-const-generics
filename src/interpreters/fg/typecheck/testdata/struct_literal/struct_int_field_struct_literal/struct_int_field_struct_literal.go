package main

type Foo struct {
	x int
}

type Bar struct {
}

func main() {
	_ = Foo{Bar{}}
}
