package main

type any interface {
}

type Foo struct {
}

func (f Foo) nth(x any, n int) int {
	return x[n]
}

func main() {
	_ = 0
}
