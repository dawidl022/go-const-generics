package main

type Foo struct {
}

func (f Foo) unbound() int {
	return x
}

func main() {
	_ = 0
}
