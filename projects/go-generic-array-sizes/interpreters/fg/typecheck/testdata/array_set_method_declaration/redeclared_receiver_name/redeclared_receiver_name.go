package main

type any interface {
}

type Arr [2]any

func (a Arr) set(x int, a any) Arr {
	a[x] = a;
	return a
}

func main() {
	_ = 0
}
