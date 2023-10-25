package main

type Foo[N const] struct {
}

func (f Foo[N]) something(x int) 42 {
	return x
}

func main() {
	_ = 1
}
