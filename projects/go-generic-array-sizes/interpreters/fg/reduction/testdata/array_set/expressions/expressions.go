package main

type Arr [2]int

func (a Arr) Set(i int, x int) Arr {
	a[i] = x;
	return a
}

func main() {
	_ = Arr{1, 2}.Set(Arr{1, 2}[0], Arr{1, 2}[0])
}
