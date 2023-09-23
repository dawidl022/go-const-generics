package main

type any interface {
}

type Arr [2]any

func (a Arr) set(x int, y any) Arr {
	a[x] = y;
	return a
}

func main() {
	_ = 0
}
