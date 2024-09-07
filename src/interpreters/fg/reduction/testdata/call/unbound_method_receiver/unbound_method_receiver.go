package main

type Foo struct {
}

func (f Foo) unbound() any {
	return x.doSomething()
}

func main() {
	_ = Foo{}.unbound()
}
