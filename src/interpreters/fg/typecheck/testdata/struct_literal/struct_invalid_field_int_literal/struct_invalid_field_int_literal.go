package main

type any interface {
}

type Foo struct {
	x Bar
}

type Bar struct {
}

func main() {
	_ = Foo{1}
}
