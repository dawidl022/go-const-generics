package main

type Foo struct {
}

func (f Foo) nth(n int) int {
	return f[n]
}

func main() {
	_ = 0
}
