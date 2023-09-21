package main

type Foo struct {
}

func (f Foo) something(x int) int {
	return f.anything(x)
}

func (f Foo) anything(x int) any {
	return x
}

func main() {
	_ = 0
}
