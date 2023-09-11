package main

type Arr [2]int

func (a Arr) unboundIndex() int {
	a[i]
}

func main() {
	_ = Arr{1, 2}.unboundIndex()
}
