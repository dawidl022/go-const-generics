package main

type Arr [2]int

func (a Arr) nth(n int) int {
	return a[n]
}

func main() {
	_ = Arr{1, 2}.nth(1)
}