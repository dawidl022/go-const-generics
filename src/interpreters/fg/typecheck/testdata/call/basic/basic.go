package main

type Foo struct {
}

func (f Foo) answer() int {
	return 42
}

func main() {
	_ = Foo{}.answer()
}
