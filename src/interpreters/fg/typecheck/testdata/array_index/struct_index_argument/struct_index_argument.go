package main

type Foo struct {
}

type Arr [2]int

func (a Arr) nth() int {
	return a[Foo{}]
}

func main() {
	_ = 0
}
