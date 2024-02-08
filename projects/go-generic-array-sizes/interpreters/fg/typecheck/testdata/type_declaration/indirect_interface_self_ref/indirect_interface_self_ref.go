package main

type Fooer interface {
	foo() Foo
}

type Foo struct {
	x Fooer
}

func main() {
	_ = 0
}
