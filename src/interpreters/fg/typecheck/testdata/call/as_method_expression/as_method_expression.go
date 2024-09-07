package main

type Foo struct {
}

func (f Foo) recurse(x int) int {
	return f.recurse(x)
}

func main() {
	_ = 0
}
