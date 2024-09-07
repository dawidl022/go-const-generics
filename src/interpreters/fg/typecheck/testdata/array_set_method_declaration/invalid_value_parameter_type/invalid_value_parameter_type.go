package main

type Arr [2]int

func (a Arr) invalidSet(x int, y any) Arr {
	a[x] = y;
	return a
}

func main() {
	_ = 0
}
