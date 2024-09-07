package main

type Arr [2]int

func (a Arr) outOfBounds() int {
	return a[2]
}

func main() {
	_ = 0
}
