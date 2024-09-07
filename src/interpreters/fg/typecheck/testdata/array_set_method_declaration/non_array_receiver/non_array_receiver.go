package main

type any interface {
}

type Foo struct {
}

func (a Foo) set(x int, y any) Foo {
	a[x] = y;
	return a
}

func main() {
	_ = 0
}
