package main

type Arr [2]int

type Foo struct {
	x Arr
}

func main() {
	_ = Foo{Arr{1, 2}}.x
}
