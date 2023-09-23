package main

type any interface {
}

type Arr [2]any

func (a Arr) invalidSet(x int, y any) Arr {
	a[x] = y;
	return x
}

func main() {
	_ = 0
}
