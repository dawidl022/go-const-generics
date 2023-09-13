package main

type Arr [2]int

func (a Arr) first() int {
	return a[0]
}

func (a Arr) first() any {
	return a[0]
}

func main() {
	_ = Arr{1, 2}.first()
}
