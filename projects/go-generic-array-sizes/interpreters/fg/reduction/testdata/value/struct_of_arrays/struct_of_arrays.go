package main

type Foo struct {
	x Arr
	y Arr
}

type Arr [2]int

func main() {
	_ = Foo{Arr{1, 2}, Arr{3, 4}}
}
