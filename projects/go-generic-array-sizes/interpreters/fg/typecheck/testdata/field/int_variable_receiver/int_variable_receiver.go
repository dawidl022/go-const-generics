package main

type Foo struct {
}

func (f Foo) getField(x int) int {
	return x.field
}

func main() {
	_ = 0
}
