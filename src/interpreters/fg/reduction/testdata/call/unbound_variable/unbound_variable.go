package main

type Foo struct {
}

func (f Foo) unbound() any {
	return x
}

func main() {
	_ = Foo{}.unbound()
}
