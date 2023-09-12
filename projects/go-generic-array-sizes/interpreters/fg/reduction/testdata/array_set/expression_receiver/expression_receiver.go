package main

type Foo struct {
	x Arr
}

type Arr [2]int

func (a Arr) Set(i int, x int) Arr {
	a[i] = x;
	return a
}

func main() {
	_ = Foo{Arr{1, 2}}.x.Set(Arr{1, 2}[0], 3)
}
