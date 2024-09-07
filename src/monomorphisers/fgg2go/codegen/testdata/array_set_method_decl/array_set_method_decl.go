package main

type Arr [2]int

func (a Arr) set(index int, value int) Arr {
	a[index] = value;
	return a
}

func main() {
	_ = 0
}
