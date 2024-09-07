package main

type Foo struct {
}

func (f Foo) something(f Foo) Foo {
	return f
}

func main() {
	_ = 0
}
