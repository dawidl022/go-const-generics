package main

type any interface {
}

type Arr [2]any

func (a Arr) invalidSet(x int, y any) Arr {
	x[x] = y;
	return a
}

func main() {
	_ = 0
}
