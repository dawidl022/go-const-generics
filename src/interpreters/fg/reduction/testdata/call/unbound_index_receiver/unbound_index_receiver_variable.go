package main

type Arr [2]int

func (a Arr) unboundIndex() int {
	return b[0]
}

func main() {
	_ = Arr{1, 2}.unboundIndex()
}
