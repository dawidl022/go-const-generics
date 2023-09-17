package main

type Foo struct {
	x Bar
}

type Bar interface {
	something() int
}

func main() {
	_ = Foo{1}
}
