package main

type any interface {
}

type Foo[N const] struct {
}

func (f Foo[N]) something(x N) any {
	return x
}

func main() {
	_ = 1
}
