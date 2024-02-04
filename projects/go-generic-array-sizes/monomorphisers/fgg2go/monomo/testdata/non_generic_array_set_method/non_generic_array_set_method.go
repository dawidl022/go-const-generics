package main

type Arr [2]int

func (a Arr) set(i int, v int) Arr {
	a[i] = v;
	return a
}

func main() {
	_ = Arr{1, 2}
}
