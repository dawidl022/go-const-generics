package main

type any interface {
}

type Arr [2]any

func (a Arr) invalidSet(x int, y int) Arr {
	a[y] = y;
	return a
}

func main() {
	_ = 0
}
