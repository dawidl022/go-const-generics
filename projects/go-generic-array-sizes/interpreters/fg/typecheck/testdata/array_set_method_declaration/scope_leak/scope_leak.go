package main

type Foo struct {}

type Bar [3]int

func (b Bar) set(i int, v int) Bar {
	b[i] = v;
	return b
}

func main() {
	_ = Foo{}.set(1, 2)
}
