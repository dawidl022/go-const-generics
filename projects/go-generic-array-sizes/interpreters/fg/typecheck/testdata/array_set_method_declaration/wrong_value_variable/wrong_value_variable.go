package main

type any interface {
}

type Arr [2]any

func (a Arr) invalidSet(x int, y any) Arr {
	a[x] = x;
	return a
}

func main() {
	_ = 0
}