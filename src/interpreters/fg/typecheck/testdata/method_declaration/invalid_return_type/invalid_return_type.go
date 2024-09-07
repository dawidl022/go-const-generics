package main

type Foo struct {
}

func (f Foo) something() Bar {
	return f
}

func main() {
	_ = 0
}
