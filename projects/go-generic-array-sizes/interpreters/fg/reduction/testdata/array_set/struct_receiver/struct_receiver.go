package main

type Foo struct {
	x int
	y int
}

func (f Foo) Set(i int, x int) Arr {
	f[i] = x;
	return f
}

func main() {
	_ = Foo{1, 2}.Set(1, 3)
}
